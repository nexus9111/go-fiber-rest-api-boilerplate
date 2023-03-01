package models

import "github.com/kamva/mgm/v3"

type Person struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model.
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Age              int    `json:"age" bson:"age"`
}

func (model *Person) CollectionName() string {
	return "person"
}
