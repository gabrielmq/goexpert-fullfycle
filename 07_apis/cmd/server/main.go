package main

import (
	"log"
	"net/http"

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

	log.Fatal(http.ListenAndServe(":8000", router))
}

// Criando um middleware que loga o request
// em Go middlawares devem retornar um http.Handler
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
