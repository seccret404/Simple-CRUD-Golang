package main

import (
	"database/sql"
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
	db, err := sql.Open("mysql", config.GetEnv("DB_DSN", "root@tcp(127.0.0.1:3306)/menu_db"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	config.ConnectDB()
	routes.CreateRoutes(app, db)

	//set port
	port := ":3000"
	
	log.Println("Server running on port " + port)
	log.Fatal(app.Listen(port))

}
