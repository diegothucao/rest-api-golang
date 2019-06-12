# Building simple restful api with Go language 
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


Then you can request 
```
Post http://localhost:1111/api/login

{
"email": "cao.trung.thu@gmail.com",
"password": "tjhsdkafdksf"
}

Get http://localhost:1111/api/getUser

Token:
"Authorization": `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiJjYW8udHJ1bmcudGh1QGdtYWlsLmNvbSJ9.7N7vWh73ELZmqG0AxRtuzGVlB8JaAVSncmCQowP6cWQ`

```

If you see any issue, please do not hesitate to create an issue here or can contact me via email cao.trung.thu@gmail.com or [Linkedin](https://www.linkedin.com/in/diegothucao/)

Thanks

references
1. https://www.tutorialspoint.com/go/
2. https://golang.org
3. https://github.com/sohamkamani/jwt-go-example
4. https://github.com/adigunhammedolalekan/go-contacts
