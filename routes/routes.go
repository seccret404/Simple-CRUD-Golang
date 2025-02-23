package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seccret404/simple-crud-golang/handlers"
)

func CreateRoutes(app *fiber.App){
	api := app.Group("/api")
	
	api.Post("/menu", handlers.CreateMenuHandler)
	api.Get("menu/:id", handlers.GetByIDHandler)
	api.Get("/menu", handlers.GetListMenuHandler)
	api.Put("/menu/:id", handlers.UpdateByIDHandler)
	api.Delete("/menu/:id", handlers.DeleteMenuHandler)

}