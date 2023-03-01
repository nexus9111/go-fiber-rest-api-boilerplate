package users

import (
	"fmt"
	"github.com/kamva/mgm/v3"
	"github.com/nexus9111/go-rest-api-boilerplate/app/dbController/users/models"
)

func Save(person *models.Person) error {
	errMsg := "error while saving person"

	err := mgm.Coll(person).Create(person)
	if err != nil {
		return fmt.Errorf("%s: %w", errMsg, err)
	}
	return nil
}

// TODO: update, delete, get
