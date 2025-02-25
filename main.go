package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/seccret404/simple-crud-golang/config"
	"github.com/seccret404/simple-crud-golang/routes"
)

func main() {
	//fiber inisialisasi
	app := fiber.New()
	config.LoadEnv()
	
	app.Static("/uploads", "./uploads")

	// secretkey := config.GetEnv("JWT_KEY_SECRET","default_value")
	// log.Println("The Key :", secretkey)

	config.ConnectDB()
	routes.CreateRoutes(app)

	//set port
	port := ":3000"
	
	log.Println("Server running on port " + port)
	log.Fatal(app.Listen(port))

}
