package main

import (
	"cembeliq_app/config"
	"cembeliq_app/routes"
	"log"
)

func main() {
	env := config.EnvVariable("ENVIRONMENT")

	log.Println("ENVIRONMENT TYPE: " + env)
	r := routes.Init()

	r.Logger.Fatal(r.Start(config.EnvVariable("APP_SERVER_PORT")))
}
