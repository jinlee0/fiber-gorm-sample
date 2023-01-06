package userRouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinlee0/fiber-gorm-sample/database"
	"github.com/jinlee0/fiber-gorm-sample/models"
)

func Route(parent fiber.Router, path string) {
	router := parent.Group(path)
	router.Post("", createUser)
	router.Get("", getUsers)
}

type user struct {
	// this is not then model user, see this as the serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func createResponseUser(userModel models.User) user {
	return user{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

func createUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := createResponseUser(user)
	return c.Status(201).JSON(responseUser)
}

func getUsers(c *fiber.Ctx) error {
	var users []models.User
	database.Database.Db.Find(&users)
	var responseUsers []user
	for _, user := range users {
		responseUser := createResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}
