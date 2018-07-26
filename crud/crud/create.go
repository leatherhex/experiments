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

type postObjectRequest struct {
	s string
}

type postObjectResponse struct {
	Err error `json:"err,omitempty"`
}

func (r postObjectResponse) error() error { return r.Err }

type getObjectRequest struct {
	ID string
}

type getObjectResponse struct {
	Object Object `json:"object,omitempty"`
	s      string `json:"Object,omitempty"`
	Err    error  `json:"err,omitempty"`
}

func (r getObjectResponse) error() error { return r.Err }

type putObjectRequest struct {
	ID string
	s  string
}

type putObjectResponse struct {
	Err error `json:"err,omitempty"`
}

func (r putObjectResponse) error() error { return nil }

type patchObjectRequest struct {
	ID string
	s  string
}

type patchObjectResponse struct {
	Err error `json:"err,omitempty"`
}

func (r patchObjectResponse) error() error { return r.Err }

type deleteObjectRequest struct {
	ID string
}

type deleteObjectResponse struct {
	Err error `json:"err,omitempty"`
}

func (r deleteObjectResponse) error() error { return r.Err }

func CreateObject(svc CrudService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("Inside CreateObject endpoint creator.")
		req := request.(ObjectRequest)
		v, err := svc.Create(ctx, req.S)
		if err != nil {
			return ObjectResponse{v, err.Error()}, nil
		}
		return ObjectResponse{v, ""}, nil
	}
}

func RetrieveObject(svc CrudService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("Inside RetrieveObject endpoint creator.")
		req := request.(ObjectRequest)
		v, err := svc.Retrieve(ctx, req.S)
		if err != nil {
			return ObjectResponse{v, err.Error()}, nil
		}
		return ObjectResponse{v, ""}, nil
	}
}

type Endpoints struct {
	PostObjectEndpoint   endpoint.Endpoint
	GetObjectEndpoint    endpoint.Endpoint
	PutObjectEndpoint    endpoint.Endpoint
	PatchObjectEndpoint  endpoint.Endpoint
	DeleteObjectEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s CrudService) Endpoints {
	fmt.Println("Inside MakeServerEndpoints.")
	return Endpoints{
		//PostObjectEndpoint:   MakePostObjectEndpoint(s),
		GetObjectEndpoint: MakeGetObjectEndpoint(s),
		//PutObjectEndpoint:    MakePutObjectEndpoint(s),
		//PatchObjectEndpoint:  MakePatchObjectEndpoint(s),
		//DeleteObjectEndpoint: MakeDeleteObjectEndpoint(s),
	}
}

func MakePostObjectEndpoint(s CrudService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// req := request.(ObjectRequest)
		// a, e := s.PostObject(ctx, req.S)
		// return postObjectResponse{Err: e}, nil
		return "hi", nil
	}
}

// MakeGetObjectEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakeGetObjectEndpoint(s CrudService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		fmt.Println("Inside MakeGetObjectEndpoint.")
		req := request.(getObjectRequest)
		p, e := s.GetObject(ctx, req.ID)
		return getObjectResponse{Object: p, Err: e}, nil
	}
}

// MakePutObjectEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakePutObjectEndpoint(s CrudService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// req := request.(putObjectRequest)
		// a, e := s.PutObject(ctx, req.ID)
		// return putObjectResponse{Err: e}, nil
		return "hi", nil
	}
}

// MakePatchObjectEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakePatchObjectEndpoint(s CrudService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// req := request.(patchObjectRequest)
		// a, e := s.PatchObject(ctx, req.ID)
		// return patchObjectResponse{Err: e}, nil
		return "hi", nil
	}
}

// MakeDeleteObjectEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakeDeleteObjectEndpoint(s CrudService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// req := request.(deleteObjectRequest)
		// a, e := s.DeleteObject(ctx, req.ID)
		// return deleteObjectResponse{Err: e}, nil
		return "hi", nil
	}
}
