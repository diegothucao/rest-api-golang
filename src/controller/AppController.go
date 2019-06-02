package controller

import (
	"encoding/json"
	"net/http"
	u "rest-api-golang/src/utils"
	"rest-api-golang/src/models"
)

var Login = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) 
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := u.Message(true, "Successful")
	resp["data"] = user
	u.Respond(w, resp)
}

var RefreshToken = func(w http.ResponseWriter, r *http.Request) {

}

var GetUser = func(w http.ResponseWriter, r *http.Request) {

}
