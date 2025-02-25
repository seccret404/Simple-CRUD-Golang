package handlers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const URL = "http://localhost:3000"

func UploadImageHandler(c *fiber.Ctx) (string, error){
	//mengambil field img adri form

	file, err := c.FormFile("image_product")
	if err != nil{
		return "", err //can replace witj json response
	}

	//validasi ukuran file

	if file.Size > 2*1024*1024{
		return "", errors.New("ukuran file terllau besar, max 2mb doang")
	}

	//validasi untuk file 
	fileAllowed := map[string]bool{".jpg":true, ".png" :true , ".jpeg":true}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !fileAllowed[ext]{
		return "", errors.New("format file di tolak, coba jpe.png atau ga jpeg")
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
		return  "", err //can replace witj json response
	}

	//url untuk akses gambar
	imgUrl := fmt.Sprintf("%s/uploads/%s", URL, filename)

	log.Println("File berhasil di upload : ", imgUrl)

	return  imgUrl,nil //can replace witj json response

}	