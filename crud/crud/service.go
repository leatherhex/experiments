package crud

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"experiments/crud/crud/store"

	"github.com/jmoiron/sqlx"
)

type Object struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

type CrudService interface {
	Create(context.Context, string) (string, error)
	Retrieve(context.Context, string) (string, error)

	PostObject(context.Context, string) (string, error)
	GetObject(context.Context, string) (Object, error)
	PutObject(context.Context, string) (string, error)
	PatchObject(context.Context, string) (string, error)
	DeleteObject(context.Context, string) (string, error)
}

type Service struct {
	db  *store.DB
	mtx sync.RWMutex
	m   map[string]Object
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		db: store.New(db),
	}
}

func (s Service) Create(c context.Context, str string) (string, error) {
	fmt.Println("Inside Create implementation of interface. str = " + str)
	if str == "" {
		return "Hello", nil
	}
	return "Hi", nil
}

func (s Service) Retrieve(c context.Context, str string) (string, error) {
	fmt.Println("Inside Retrieve implementation of interface. str = " + str)
	if str == "" {
		return "Hello", nil
	}
	return "Hi", nil
}

func (s Service) PostObject(c context.Context, str string) (string, error) {
	fmt.Println("Inside PostObject implementation of interface. str = " + str)
	if str == "" {
		return "Hello", nil
	}
	return "Hi", nil
}

func (s Service) GetObject(c context.Context, id string) (Object, error) {
	fmt.Println("Inside GetObject implementation of interface. str = " + id)
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	p, ok := s.m[id]
	fmt.Println("id = " + id)
	if !ok {
		return Object{}, ErrNotFound
	}
	return p, nil
}

func (s Service) PutObject(c context.Context, str string) (string, error) {
	fmt.Println("Inside PutObject implementation of interface. str = " + str)
	if str == "" {
		return "Hello", nil
	}
	return "Hi", nil
}

func (s Service) PatchObject(c context.Context, str string) (string, error) {
	fmt.Println("Inside PatchObject implementation of interface. str = " + str)
	if str == "" {
		return "Hello", nil
	}
	return "Hi", nil
}

func (s Service) DeleteObject(c context.Context, str string) (string, error) {
	fmt.Println("Inside DeleteObject implementation of interface. str = " + str)
	if str == "" {
		return "Hello", nil
	}
	return "Hi", nil
}
