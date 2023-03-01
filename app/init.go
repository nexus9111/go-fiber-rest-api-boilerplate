package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"github.com/nexus9111/go-rest-api-boilerplate/app/models"
	"github.com/nexus9111/go-rest-api-boilerplate/app/personController"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const PORT = ":3000"

var router = []models.Router{
	{
		Method:      models.GET,
		Path:        "/",
		RelatedFunc: helloWorld,
	},
	{
		Method:      models.POST,
		Path:        "/",
		RelatedFunc: postHelloWorld,
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

func postHelloWorld(c *fiber.Ctx) error {
	p := new(models.Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	return c.JSON(p)
}

func Init() {
	err := mgm.SetDefaultConfig(nil, "go-rest-api-boilerplate", options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected")

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

	err = app.Listen(PORT)
	if err != nil {
		panic(err)
	}
}
