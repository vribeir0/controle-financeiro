package main

import (
	"server/app/routes"
	"server/app/config"
)

func main() {
	config.Connect()
	routes.Run()

}
