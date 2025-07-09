package main

import (
	"inventory-app/config"
	"inventory-app/routes"
	"net/http"
)

func main() {
	config.InitDB()
	r := routes.SetupRoutes()
	http.ListenAndServe(":8080", r)
}
