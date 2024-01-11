package main

import (
	"my_spp/configs"
	"my_spp/database"
	m "my_spp/middlewares"
	"my_spp/routes"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc
	configs.Init()
	database.Init()
	e := routes.New()
	//implement middleware logger
	m.LogMiddlewares(e)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
