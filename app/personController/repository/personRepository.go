package repository

import (
	"fmt"
	"github.com/nexus9111/go-rest-api-boilerplate/app/dbController/users"
	"github.com/nexus9111/go-rest-api-boilerplate/app/dbController/users/models"
)

func NewPerson(name string, age int) error {
	errMsg := "Error while creating person"

	newPerson := &models.Person{
		Name: name,
		Age:  age,
	}
	err := users.Save(newPerson)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}

	return nil
}
