package routes

import (
	"fmt"
	"net/http"
	"rest-api-golang/src/controller"

	"github.com/gorilla/mux"
)

var SetupServer = func(appPort string) {
	var router = mux.NewRouter()

	router.HandleFunc("/api/login", controller.Login).Methods("POST")
	router.HandleFunc("/api/getUser", controller.GetUser).Methods("GET")

	fmt.Println(appPort)
	err := http.ListenAndServe(":"+appPort, router)
	if err != nil {
		fmt.Print(err)
	}
}
