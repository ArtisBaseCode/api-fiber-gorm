package routes

import (
	"errors"

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

func GetUsers(c *fiber.Ctx) error {
	users := []model.User{}

	database.Database.Db.Find(&users)
	usersReply := []CreateUserReq{}

	for _, user := range users {
		userReply := CreateUserRequest(user)
		usersReply = append(usersReply, userReply)
	}

	return c.Status(200).JSON(usersReply)
}

func findUser(id int, user *model.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("User does not exist")
	}

	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user model.User
	if err != nil {
		return c.Status(200).JSON("Please ensure that :id is an int")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	userReply := CreateUserRequest(user)

	return c.Status(200).JSON(userReply)
}
