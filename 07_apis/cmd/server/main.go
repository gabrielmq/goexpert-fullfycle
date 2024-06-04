package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gabrielmq/apis/configs"
	_ "github.com/gabrielmq/apis/docs" // importa a doc para ser enncontrada
	"github.com/gabrielmq/apis/internal/entity"
	"github.com/gabrielmq/apis/internal/infra/database"
	"github.com/gabrielmq/apis/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// comentários que serão lidos pelo swag para gerar a doc
// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Gabriel
// @contact.url
// @contact.email

// @license.name
// @license.url

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// carregando as configs externas
	cfg := configs.LoadConfiguration(".")

	// abrindo uma conexao com a base
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Criando as tabelas
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productHandlers := handlers.NewProductHandler(database.NewProduct(db))
	usersHandlers := handlers.NewUserHandler(database.NewUser(db))

	// registrando as rotas
	router := chi.NewRouter()

	// adicionando os middlewares
	router.Use(middleware.Logger)

	// middlaware que recupera a app de panics, evitando que a morra
	router.Use(middleware.Recoverer)

	// middleware para injetar coisas no context da req
	router.Use(middleware.WithValue("jwt", cfg.TokenAuth))
	router.Use(middleware.WithValue("expires_in", cfg.JwtExperesIn))

	// /products
	router.Route("/products", func(r chi.Router) {
		// adiciona middleware de verificação do token jwt na req para /products
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		// valida o token
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandlers.CreateProduct)
		r.Put("/{id}", productHandlers.UpdateProduct)
		r.Get("/{id}", productHandlers.GetProduct)
		r.Get("/", productHandlers.FindAllProducts)
		r.Delete("/{id}", productHandlers.DeleteProduct)
	})

	// /users
	router.Post("/users", usersHandlers.Create)
	router.Post("/users/generate_token", usersHandlers.GetJwt)

	// criando a rota para a doc swagger
	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	slogJsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	slog.SetDefault(slog.New(slogJsonHandler))

	server := &http.Server{
		Addr:     ":8000",
		Handler:  router,
		ErrorLog: slog.NewLogLogger(slogJsonHandler, slog.LevelError),
	}

	serverError := make(chan error, 1)
	go func() {
		slog.Info("starting application on port :8000")
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				serverError <- err
			}
		}
	}()

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-serverError:
		slog.Error(fmt.Sprintf("failed to start application: %v", err))
	case <-shutdownSignal:
		slog.Info("shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			slog.Error(fmt.Sprintf("could not gracefully shutdown the server: %v", err))
			return
		}

		slog.Info("server shutdown gracefully")
	}
}

// Criando um middleware que loga o request
// em Go middlawares devem retornar um http.Handler
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
