package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/nexus9111/go-rest-api-boilerplate/app/dbController"
	"github.com/nexus9111/go-rest-api-boilerplate/app/models"
	"github.com/nexus9111/go-rest-api-boilerplate/app/personController"
	"os"
)

var router = []models.Router{
	{
		Method:      models.GET,
		Path:        "/",
		RelatedFunc: helloWorld,
	},
	{
		Method:      models.POST,
		Path:        "/person",
		RelatedFunc: personController.NewPerson,
	},
}

func helloWorld(c *fiber.Ctx) error {

	response := models.ResponseStructure{
		Success: true,
		Message: "Hello World",
		Data: map[string]string{
			"author": "nexus9111",
		},
	}

	return c.JSON(response)
}

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	err = dbController.Connect()
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	for _, r := range router {
		switch r.Method {
		case models.GET:
			app.Get(r.Path, r.RelatedFunc)
		case models.POST:
			app.Post(r.Path, r.RelatedFunc)
		default:
			panic("Method not already implemented")
		}
	}

	err = app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}
