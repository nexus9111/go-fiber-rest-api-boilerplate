package personController

import (
	"github.com/gofiber/fiber/v2"
	personModels "github.com/nexus9111/go-rest-api-boilerplate/app/dbController/users/models"
	"github.com/nexus9111/go-rest-api-boilerplate/app/models"
	"github.com/nexus9111/go-rest-api-boilerplate/app/personController/repository"
)

func NewPerson(c *fiber.Ctx) error {
	p := new(personModels.Person)

	err := c.BodyParser(p)
	if err != nil {
		return err
	}

	err = repository.NewPerson(p.Name, p.Age)
	if err != nil {
		return c.JSON(models.ResponseStructure{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(models.ResponseStructure{
		Success: true,
		Message: "Person created successfully",
		Data: map[string]interface{}{
			"person": p,
		},
	})
}
