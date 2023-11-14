package routes

import (
	"github.com/artisbasecode/api-fiber-gorm/database"
	"github.com/artisbasecode/api-fiber-gorm/model"
	"github.com/gofiber/fiber/v2"
)

type CreateUserReq struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateUserRequest(user model.User) CreateUserReq {
	return CreateUserReq{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	userReply := CreateUserRequest(user)

	return c.Status(200).JSON(userReply)
}
