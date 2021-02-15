package main

import "cembeliq_app/routes"

func main() {
	r := routes.Init()

	r.Logger.Fatal(r.Start(":9000"))
}
