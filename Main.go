package main

import (
	"os"
	"rest-api-golang/routes"
)

func main() {
	var appPort = os.Getenv("appPort")

	if appPort == "" {
		appPort = "1111"
	}

	routes.SetupServer(appPort)
}
