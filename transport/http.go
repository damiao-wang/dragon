package transport

import (
	"context"
	"encoding/json"
	"net/http"

	addendpoint "dragon/endpoint"
	"dragon/pb"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(endpoints addendpoint.Set, logger log.Logger) http.Handler {
	// Zipkin HTTP Server Trace can either be instantiated per endpoint with a
	// provided operation name or a global tracing service can be instantiated
	// without an operation name and fed to each Go kit endpoint as ServerOption.
	// In the latter case, the operation name will be the endpoint's http method.
	// We demonstrate a global tracing service here.

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorLogger(logger),
	}

	r := mux.NewRouter()
	r.Methods("POST").Path("/sum").Handler(httptransport.NewServer(
		endpoints.SumEndpoint,
		decodeSumRequest,
		encodeResponse,
		options...,
	))
	r.Methods("POST").Path("/concat").Handler(httptransport.NewServer(
		endpoints.ConcatEndpoint,
		decodeConcatRequest,
		encodeResponse,
		options...,
	))
	return r
}

func decodeSumRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.SumReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return &request, nil
}

func decodeConcatRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request pb.ConcatReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return &request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

type errorWrapper struct {
	Error string `json:"error"`
}
