package crud

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

func RetrieveObject(svc CrudService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("This returns a function that is actually an endpoint.")
		return nil, nil
	}
}
