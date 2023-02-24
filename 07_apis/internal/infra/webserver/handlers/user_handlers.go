package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gabrielmq/apis/internal/entity"
	"github.com/gabrielmq/apis/internal/infra/database"
	"github.com/gabrielmq/apis/internal/infra/dto"
	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
	}
}

// GetJwt godoc
// @Summary      Get a user JWT
// @Description  Get a user JWT
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body     dto.GetJwtInput  true  "user credentials"
// @Success      200  {object}  dto.GetJwtOutput
// @Failure      401  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/generate_token [post]
func (uh *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	var userInput dto.GetJwtInput
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	user, err := uh.UserDB.FindByEmail(userInput.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	if !user.ValidatePassword(userInput.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("expires_in").(int)

	// criando o token
	_, token, _ := jwt.Encode(map[string]interface{}{
		"sub":        user.ID.String(),
		"expires_in": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	accessToken := dto.GetJwtOutput{AccessToken: token}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CreateUserInput  true  "user request"
// @Success      201  {object}  entity.User
// @Failure      500         {object}  Error
// @Router       /users [post]
func (uh *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var userInput dto.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	user, err := entity.NewUser(
		userInput.Name,
		userInput.Email,
		userInput.Password,
	)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	if err := uh.UserDB.Create(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
