package routes

import (
	"time"

	"github.com/artisbasecode/api-fiber-gorm/database"
	"github.com/artisbasecode/api-fiber-gorm/model"
	"github.com/gofiber/fiber/v2"
)

/*
How it will look

	id: 1
	user: {
		id:12
		first_name: abc
		lastName: abc
	}
	product: {
		id:24
		name: abc
		serial_name: abc
	}
*/
type OrderReq struct {
	ID        uint          `json:"id"`
	User      CreateUserReq `json:"user"`
	Product   ProductReq    `json:"product"`
	CreatedAt time.Time     `json:"order_date"`
}

func CreateOrderReply(order model.Order, user CreateUserReq, product ProductReq) OrderReq {
	return OrderReq{ID: order.ID, User: user, Product: product}
}

func CreateOrder(c *fiber.Ctx) error {
	var order model.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var user model.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product model.Product

	if err := findProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	userReply := CreateUserRequest(user)
	productReply := CreateProductReply(product)
	orderReply := CreateOrderReply(order, userReply, productReply)

	return c.Status(200).JSON(orderReply)
}
