package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/TSRangel/Go-Test-Basic_API/internal/dto"
	"github.com/TSRangel/Go-Test-Basic_API/internal/entities/user"
	"github.com/TSRangel/Go-Test-Basic_API/internal/infra/database"

	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	DB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		DB: db,
	}
}

// Create 		godoc
// @Summary 	Create user
// @Description Create user
// @Tags 		users
// @Accept 		json
// @Produce 	json
// @Param 		request 	body 	dto.CreateUserInput 	true 	"user request"
// @Success 	201
// @Failure 	500		{object}	Error
// @Router		/users	[post]
func (uh *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newUserDTO dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&newUserDTO)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if newUserDTO.Name == "" || newUserDTO.Email == "" || newUserDTO.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUser, err := user.NewUser(newUserDTO.Name, newUserDTO.Email, newUserDTO.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = uh.DB.Create(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// Login 		godoc
// @Summary 	Get a user JWT
// @Description Get a user JWT
// @Tags 		users
// @Accpet 		json
// @Produce 	json
// @Param 		request 	body 	dto.UserLoginInput 	true 	"user credentials"
// @Success 	200 	{objetc} 	dto.JWTTokenOutput
// @Failure 	404		{object}	Error
// @Failure 	500 	{object} 	Error
// @Router 		/users/login 		[post]
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	jwtToken := r.Context().Value("token").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("expiresIn").(int)
	var userLoginDTO dto.UserLoginInput
	err := json.NewDecoder(r.Body).Decode(&userLoginDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	searchedUser, err := uh.DB.FindByEmail(userLoginDTO.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	if !searchedUser.ValidatePassword(userLoginDTO.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, err := jwtToken.Encode(map[string]interface{}{
		"sub": searchedUser.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	accessToken := dto.JWTTokenOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
	w.WriteHeader(http.StatusOK)
}
