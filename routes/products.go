package routes

import (
	"github.com/artisbasecode/api-fiber-gorm/database"
	"github.com/artisbasecode/api-fiber-gorm/model"
	"github.com/gofiber/fiber/v2"
)

type ProductReq struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateProductReply(productModel model.Product) ProductReq {
	return ProductReq{ID: productModel.ID, Name: productModel.Name, SerialNumber: productModel.SerialNumber}
}

func CreateProduct(c *fiber.Ctx) error {
	var product model.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)
	productReply := CreateProductReply(product)

	return c.Status(200).JSON(productReply)
}
