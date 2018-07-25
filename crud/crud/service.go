package crud

import (
	"context"
	"fmt"

	"experiments/crud/crud/store"

	"github.com/jmoiron/sqlx"
)

type CrudService interface {
	Retrieve(context.Context, string) (string, error)
}

type Service struct {
	db *store.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		db: store.New(db),
	}
}

func (s Service) Retrieve(c context.Context, str string) (string, error) {
	fmt.Printf("Inside Retrieve implementation of interface. str = " + str)
	if str == "" {
		return "Hello", nil
	}
	return "Hi", nil
}
