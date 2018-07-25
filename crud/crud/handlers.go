package crud

import (
	"context"
	"fmt"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

//MakeHandler The function that provides all the transport handling linking for this service
func MakeHandler(svc CrudService) {

	fmt.Println("In handler")
	e := RetrieveObject(svc)           //returns an endpoint (create.go)
	dec := decodeRetrieveObjectRequest //i don't know what this is for yet
	enc := encodeRetrieveObjectRequest //i don't know what this is for yet

	getObjectHandler := httptransport.NewServer(e, dec, enc)
	r := mux.NewRouter()

	fmt.Println("Calling handle")
	r.Handle("/my_endpoint.json", getObjectHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe("127.0.0.1:8080", r)
}

func decodeRetrieveObjectRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeRetrieveObjectRequest(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	w.WriteHeader(200)
	return nil
}
