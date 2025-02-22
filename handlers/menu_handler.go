package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/seccret404/simple-crud-golang/config"
	"github.com/seccret404/simple-crud-golang/db/models"
)

type CreateMenuRequest struct {
	Name         string `json:"name_product"`
	Price        int64  `json:"price"`
	Description  string `json:"description_product"`
	ImageProduct string `json:"image_product"`
	StockProduct int64  `json:"stock_product"`
}

func CreateMenuHandler(c *fiber.Ctx) error{
	db := config.ConnectDB()
	queries := models.New(db)

	var req CreateMenuRequest
	if err := c.BodyParser(&req); err != nil{
		return c.Status(400).JSON(fiber.Map{"error" : "Invalid Request"})
	}

	result, err := queries.CreateMenu(c.Context(), models.CreateMenuParams{
		NameProduct: req.Name,
		Price: req.Price,
		DescriptionProduct: req.Description,
		ImageProduct: req.ImageProduct,
		StockProduct: int32(req.StockProduct),
	})

	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{"Error " : "Failerd to create menu"})
	}

	lastInsertID, _ := result.LastInsertId()
	return c.JSON(fiber.Map{"message": "Menu Create", "id" : lastInsertID})
}

