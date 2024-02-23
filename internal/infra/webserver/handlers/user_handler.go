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

type UserHandler struct {
	DB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		DB: db,
	}
}

func(uh *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
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
		return
	}

	err = uh.DB.Create(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func(uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusBadRequest)
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
	accessToken := struct{
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
	w.WriteHeader(http.StatusOK)
}