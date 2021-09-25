package main

import (
	"quiztopedia-backend/config"
	"quiztopedia-backend/routes"
)

func main() {
	config.Connect()
	routes.Run()
}
