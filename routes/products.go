package routes

import (
	"errors"

	"github.com/artisbasecode/api-fiber-gorm/database"
	"github.com/artisbasecode/api-fiber-gorm/model"
	"github.com/gofiber/fiber/v2"
)

var Db = database.Database.Db

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

func GetProducts(c *fiber.Ctx) error {
	products := []model.Product{}

	database.Database.Db.Find(&products)
	productsReply := []ProductReq{}

	for _, product := range products {
		productReply := CreateProductReply(product)
		productsReply = append(productsReply, productReply)
	}

	return c.Status(200).JSON(productsReply)
}

func findProduct(id int, product *model.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("Product does not exist")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(200).JSON("Please ensure that :id is an int")
	}
	var product model.Product

	if err := findProduct(id, &product); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	productReply := CreateProductReply(product)

	return c.Status(200).JSON(productReply)
}
