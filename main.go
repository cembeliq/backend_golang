package main

import (
	"cembeliq_app/config"
	"cembeliq_app/routes"
	"log"
)

func main() {
	env := config.EnvVariable("ENVIRONMENT")
	// result, err := models.Find("tst", "tes")
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Printf("%+v", result)

	// log.Println(result)

	log.Println("ENVIRONMENT TYPE: " + env)
	r := routes.Init()

	r.Logger.Fatal(r.Start(config.EnvVariable("APP_SERVER_PORT")))
}
