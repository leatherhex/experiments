package crud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

//MakeHandler The function that provides all the transport handling linking for this service
func MakeHandler(svc CrudService) {

	fmt.Println("Creating create handler")
	e := CreateObject(svc)           //returns an endpoint (create.go)
	dec := decodeCreateObjectRequest //i don't know what this is for yet
	enc := encodeCreateObjectRequest //i don't know what this is for yet
	createObjectHandler := httptransport.NewServer(e, dec, enc)
	r := mux.NewRouter()
	r.Handle("/my_endpoint.json", createObjectHandler).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe("127.0.0.1:8080", r)
}

func MakeHTTPHandler(s CrudService, logger log.Logger) http.Handler {
	fmt.Println("Inside MakeHTTPHandler")
	r := mux.NewRouter()
	e := MakeServerEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("GET").Path("/my_endpoint/retrieve.json").Handler(httptransport.NewServer(
		e.GetObjectEndpoint,
		decodeGetObjectRequest,
		encodeResponse,
		options...,
	))
	// r.Methods("POST").Path("/my_endpoint/create.json").Handler(httptransport.NewServer(

	// 	e.PostObjectEndpoint,
	// 	decodePostObjectRequest,
	// 	encodeResponse,
	// 	options...,
	// ))
	// r.Methods("PUT").Path("/my_endpoint/{id}").Handler(httptransport.NewServer(
	// 	e.PutObjectEndpoint,
	// 	decodePutObjectRequest,
	// 	encodeResponse,
	// 	options...,
	// ))
	// r.Methods("PATCH").Path("/my_endpoint/{id}").Handler(httptransport.NewServer(
	// 	e.PatchObjectEndpoint,
	// 	decodePatchObjectRequest,
	// 	encodeResponse,
	// 	options...,
	// ))
	// r.Methods("DELETE").Path("/my_endpoint/{id}").Handler(httptransport.NewServer(
	// 	e.DeleteObjectEndpoint,
	// 	decodeDeleteObjectRequest,
	// 	encodeResponse,
	// 	options...,
	// ))
	return r
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func decodeCreateObjectRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("Inside decodeCreateObjectRequest.")
	var request ObjectRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeCreateObjectRequest(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeGetObjectRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("Inside decodeRetrieveObjectRequest.")
	var request ObjectRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func codeFrom(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrAlreadyExists, ErrInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
