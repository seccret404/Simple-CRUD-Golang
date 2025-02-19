package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seccret404/simple-crud-golang/config"
)

func main() {
	db	:= config.ConnectDB()
	defer config.CLoseDB()

	//fiber inisialisasi
	app := fiber.New();

	app.Get("/", func(c *fiber.Ctx) error{
		err := db.Ping()
		if err != nil{
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Database tidak terhubung"})
		}
		return c.JSON(fiber.Map{"status": "success", "message": "Database terhubung"})
	})

	app.Listen(":3000");
	
}
