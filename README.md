# Rest-api-Golang
This is an essential example to build server with restful api using Golang

Step to run
1. Clone the [repo](https://github.com/diegothucao/rest-api-golang)
2. `go build`
3. `go run rest-api-golang`

Define route 

```go 
var SetupServer = func(appPort string) {
	var router = mux.NewRouter()

	router.HandleFunc("/api/login", controller.Login).Methods("POST")
	router.HandleFunc("/api/getUser", controller.GetUser).Methods("GET")

	router.Use(JwtAuthentication) 

	err := http.ListenAndServe(":"+appPort, router)
	if err != nil {
		fmt.Print(err)
	}
}
```

Process a request 

```go 
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
```
