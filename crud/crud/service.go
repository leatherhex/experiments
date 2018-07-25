package crud

import (
	"fmt"

	"experiments/crud/crud/store"

	"github.com/jmoiron/sqlx"
)

type CrudService interface {
	Create()
	Retrieve()
	Update()
	Delete()
}

type Service struct {
	db *store.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		db: store.New(db),
	}
}

func (s Service) Retrieve() {
	fmt.Printf("This implements the Retrieve function of the interface. Then what?")
}
