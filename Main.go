package main

import (
	"rest-api-golang/routes"
	"os"
)


func main() {
	var appPort = os.Getenv("appPort")

	if appPort == "" {
		appPort = "1111"
	}

	routes.setupServer(appPort)
}
