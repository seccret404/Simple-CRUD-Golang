package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/seccret404/simple-crud-golang/handlers"
	"github.com/seccret404/simple-crud-golang/middleware"
)

func CreateRoutes(app *fiber.App, db *sql.DB){
	api := app.Group("/api")
	authHandler := handlers.AuthHandler{DB: db}
	//auth
	api.Post("/register", authHandler.RegisterUser)
	api.Post("/login", authHandler.LoginUser)
	
	//menu endpoint
	api.Post("/menu", middleware.JwtMiddleware, handlers.CreateMenuHandler)// middleware
	api.Get("menu/:id", handlers.GetByIDHandler)
	api.Get("/menu", handlers.GetListMenuHandler)
	api.Put("/menu/:id", handlers.UpdateByIDHandler)
	api.Delete("/menu/:id", handlers.DeleteMenuHandler)

	//file endpoint

	// api.Post("/uploads", handlers.UploadImageHandler) to test img


}