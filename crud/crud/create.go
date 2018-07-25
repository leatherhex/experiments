package crud

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

type myObjectContext struct {
	name        string
	description string
}

type ObjectResponse struct {
	body string
	code int
}

func RetrieveObject(svc CrudService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("This returns a function that is actually an endpoint.")
		return nil, nil
	}
}
