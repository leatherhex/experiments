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

type ObjectRequest struct {
	S string `json:"s"`
}

type ObjectResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

func RetrieveObject(svc CrudService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Printf("Inside RetrieveObject endpoint creator.")
		req := request.(ObjectRequest)
		v, err := svc.Retrieve(ctx, req.S)
		if err != nil {
			return ObjectResponse{v, err.Error()}, nil
		}
		return ObjectResponse{v, ""}, nil
	}
}
