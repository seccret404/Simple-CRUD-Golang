package handlers

import (
	"database/sql"
	"log"
	"strconv"

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

	//call function handler upload
		imageUrl, err := UploadImageHandler(c)
		if err != nil{
			return c.Status(400).JSON(fiber.Map{"error" : "Gagal menyimpan gambar", "detail": err.Error()})
		}
		//if we test frorm form-data
		req.ImageProduct = imageUrl
		req.Name = c.FormValue("name_product")
		req.Description = c.FormValue("description_product")

		//set type data string to int
		StockProduct := c.FormValue("stock_product")
		stock, err := strconv.Atoi(StockProduct)
		if err != nil{
			return c.Status(400).JSON(fiber.Map{"error" : "Gagal convert ke int"})
		}
		
		req.StockProduct = int64(stock) //final value

	result, err := queries.CreateMenu(c.Context(), models.CreateMenuParams{
		NameProduct: req.Name,
		Price: req.Price,
		DescriptionProduct: req.Description,
		ImageProduct: imageUrl,
		StockProduct: int32(req.StockProduct),
	})

	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{"Error " : "Failerd to create menu"})
	}

	lastInsertID, _ := result.LastInsertId()
	return c.JSON(fiber.Map{"message": "Menu Create", "id" : lastInsertID})
}

func GetListMenuHandler(c *fiber.Ctx) error{
	db := config.ConnectDB()
	queries := models.New(db)

	menu, err := queries.ListMenus(c.Context())
	if err != nil{
		return c.Status(500).JSON(fiber.Map{"error" : "Failed to get list menu"})
	}

	return c.JSON((menu))

}

func GetByIDHandler(c *fiber.Ctx) error{
	db := config.ConnectDB()
	queries := models.New(db)

	idStr := c.Params("id")
	
	//conversi ID
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	menu, err := queries.GetMenuByID(c.Context(), int32(id))
	if err == sql.ErrNoRows{
		return c.Status(400).JSON(fiber.Map{"error" : "Menu not found"})
	}else if err != nil{
		return c.Status(500).JSON(fiber.Map{"error" : "Internal servel error"})
	}

	return c.JSON(menu)

}

func UpdateByIDHandler(c *fiber.Ctx) error{
	db := config.ConnectDB()
	queries := models.New(db)

	idStr := c.Params("id")

	//convrsi id

	id, err := strconv.Atoi(idStr)
	if err != nil{
		return c.Status(400).JSON(fiber.Map{"error" : "Invalid ID format"})
	}

	var req CreateMenuRequest
	if err := c.BodyParser(&req); err != nil{
		return c.Status(400).JSON(fiber.Map{"error" : "Invalid Request"})
	}

	err = queries.UpdateMenu(c.Context(), models.UpdateMenuParams{
		NameProduct: req.Name,
		DescriptionProduct: req.Description,
		ImageProduct: req.ImageProduct,
		Price: req.Price,
		StockProduct: int32(req.StockProduct),
		ID: int32(id),
	})

	if err != nil{
		return c.Status(500).JSON(fiber.Map{"error" : "Failed to update mennu"})
	}

	return c.JSON(fiber.Map{"message" : "Menu updated"})

}

func DeleteMenuHandler(c *fiber.Ctx) error{
	db := config.ConnectDB()
	queries := models.New(db)

	idStr := c.Params("id")

	id, err := strconv.Atoi(idStr)
	if err != nil{
		return c.Status(400).JSON(fiber.Map{"error" : "Invalid ID Format"})
	}

	err = queries.DeleteMenu(c.Context(), int32(id))
	if err != nil{
		return c.Status(500).JSON(fiber.Map{"error" : "Failed to deleted menu"})
	}

	return c.JSON(fiber.Map{"message" : "Menu deleted"})
	
}

