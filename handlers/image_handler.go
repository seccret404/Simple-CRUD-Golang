package handlers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const URL = "http://localhost:3000"

func UploadImageHandler(c *fiber.Ctx) error{
	//mengambil field img adri form

	file, err := c.FormFile("image")
	if err != nil{
		return c.Status(400).JSON(fiber.Map{"error" : "Gagal dalam mengambil file"})
	}

	//check folder uploads
	uploaddDir := "uploads"
	if _, err := os.Stat(uploaddDir);os.IsNotExist(err){
		os.Mkdir(uploaddDir,os.ModePerm)
	}

	//set name file
	filename := fmt.Sprintf("%d-%s", os.Getpid(), file.Filename)
	filePath := filepath.Join(uploaddDir,filename)
	
	//sv file to folder 
	if err := c.SaveFile(file, filePath); err != nil{
		return c.Status(500).JSON(fiber.Map{"error" : "Gagal menyimpan file"})
	}

	//url untuk akses gambar
	imgUrl := fmt.Sprintf("%s/uploads/%s", URL, filename)

	log.Println("File berhasil di upload : ", imgUrl)

	return c.JSON(fiber.Map{"message" : "Upload Berhasil ", "Image url : " : imgUrl})

}	