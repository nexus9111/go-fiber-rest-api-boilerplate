package dbController

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func Connect() error {
	err := mgm.SetDefaultConfig(nil, os.Getenv("MONGO_DB"), options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		return err
	}
	return nil
}
