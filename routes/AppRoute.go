package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var SetupServer = func(appPort string) {
	var router = mux.NewRouter()
	fmt.Println(appPort)
	err := http.ListenAndServe(":"+appPort, router)
	if err != nil {
		fmt.Print(err)
	}
}
