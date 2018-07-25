package crud

import (
	"fmt"

	"experiments/crud/crud/store"

	"github.com/jmoiron/sqlx"
)

type CrudService interface {
	Retrieve(myObjectContext) (ObjectResponse, error)
}

type Service struct {
	db *store.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		db: store.New(db),
	}
}

func (s Service) Retrieve(c myObjectContext) (ObjectResponse, error) {
	fmt.Printf("Context's name is: " + c.name)
	var resp ObjectResponse
	fmt.Printf("This implements the Retrieve function of the interface. Then what?")
	return resp, nil
}
