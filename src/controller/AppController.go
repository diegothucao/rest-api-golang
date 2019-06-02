package controller

import (
	"encoding/json"
	"net/http"
	u "rest-api-golang/src/utils"
	"rest-api-golang/src/models"
	"github.com/dgrijalva/jwt-go"
	"os"
)

var Login = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) 
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	user.Password = ""
	tk := &models.Token{UserId: user.Email}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))

	resp := u.Message(true, "Successful")
	resp["data"] = user
	resp["token"] = tokenString
	u.Respond(w, resp)
}

var GetUser = func(w http.ResponseWriter, r *http.Request) {
	user1 := &models.User{}
	user2 := &models.User{}
	user1.Email = "cao.trung.thu@mail.com"
	user2.Email = "cao.trung.thu@hot.com"

	var users [2]*models.User
	users[0] = user1
	users[1] = user2
	resp := u.Message(true, "Successful")
	resp["data"] = users
	u.Respond(w, resp)
}
