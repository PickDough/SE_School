package controllers

import (
	"encoding/json"
	"github.com/Pick-Down/BTC_API/models"
	"github.com/Pick-Down/BTC_API/utils"
	"net/http"
)

type UserService interface {
	AddUser(user models.User) error
	LoginUser(user models.User) error
}

type AuthService interface {
	GenerateToken(userEmail string) (models.Token, error)
}

var UserServ UserService
var AuthServ AuthService

//Login Function that responds to user/login endpoint - logins user.
func Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, utils.Message("Invalid request parameters"))
		return
	}

	err := UserServ.LoginUser(*user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, utils.Message(err.Error()))
		return
	}

	//Generating token
	token, err := AuthServ.GenerateToken(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.Respond(w, utils.Message(err.Error()))
	}

	utils.Respond(w, map[string]interface{}{"accessToken": token})
}

//Create Function that responds to user/create endpoint - adds new user.
func Create(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.Respond(w, utils.Message("Invalid request parameters"))
		return
	}

	err := UserServ.AddUser(*user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.Respond(w, utils.Message(err.Error()))
		return
	}

	//Generating token
	token, err := AuthServ.GenerateToken(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.Respond(w, utils.Message(err.Error()))
	}

	utils.Respond(w, map[string]interface{}{"accessToken": token})
}
