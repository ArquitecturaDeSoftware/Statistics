package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	http1 "net/http"

	http "github.com/go-kit/kit/transport/http"
	handlers "github.com/gorilla/handlers"
	mux "github.com/gorilla/mux"
	endpoint "github.com/jscastelblancoh/statistic_service/statistic/pkg/endpoint"
	"github.com/jscastelblancoh/statistic_service/statistic/pkg/io"
)

// makeGetHandler creates the handler logic
func makeGetHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/statistics/").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetEndpoint, decodeGetRequest, encodeGetResponse, options...)))
}

// decodeGetResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.GetRequest{}
	return req, nil
}

// encodeGetResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetbyIdHandler creates the handler logic
func makeGetbyIdHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("GET").Path("/statistics/{id}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"GET", "GetbyId"}), handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.GetbyIdEndpoint, decodeGetbyIdRequest, encodeGetbyIdResponse, options...)))
}

// decodeGetbyIdResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetbyIdRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("Not a valid ID")
	}
	req := endpoint.GetbyIdRequest{
		Id: id,
	}
	return req, nil
}

// encodeGetbyIdResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetbyIdResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makePutHandler creates the handler logic
func makePutHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("PUT").Path("/statistic/").Handler(handlers.CORS(handlers.AllowedMethods([]string{"POST"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.PutEndpoint, decodePutRequest, encodePutResponse, options...)))
}

// decodePutResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePutRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	req := endpoint.PutRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodePutResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePutResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDeleteHandler creates the handler logic
func makeDeleteHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("DELETE").Path("/statistics/{id}").Handler(handlers.CORS(handlers.AllowedMethods([]string{"DELETE"}), handlers.AllowedHeaders([]string{"Content-Type", "Content-Length"}), handlers.AllowedOrigins([]string{"*"}))(http.NewServer(endpoints.DeleteEndpoint, decodeDeleteRequest, encodeDeleteResponse, options...)))
}

// decodeDeleteResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDeleteRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("not a valid ID")
	}
	req := endpoint.DeleteRequest{
		Id: id,
	}
	return req, nil
}

// encodeDeleteResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDeleteResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makePostHandler creates the handler logic
func makePostHandler(m *mux.Router, endpoints endpoint.Endpoints, options []http.ServerOption) {
	m.Methods("POST").Path("/statistics/").Handler(
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"Content-Type", "application/json"}),
			handlers.AllowedMethods([]string{"POST"}),
		)(http.NewServer(endpoints.PostEndpoint, decodePostRequest, encodePostResponse, options...)))
}

// decodePostResponse  is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodePostRequest(_ context.Context, r *http1.Request) (interface{}, error) {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	var req2 io.Statistic
	req := endpoint.PostRequest{}
	err := json.Unmarshal(body, &req2)
	req.Statistic.Id = req2.Id
	req.Statistic.Id_restaurant = req2.Id_restaurant
	req.Statistic.Date = req2.Date
	req.Statistic.Sold_lunches = req2.Sold_lunches
	req.Statistic.Canceled_shifts = req2.Canceled_shifts
	req.Statistic.Av_time = req2.Av_time
	req.Statistic.Av_punctuation = req2.Av_punctuation
	req.Statistic.Bonus_sold = req2.Bonus_sold
	req.Statistic.Student_sold = req2.Student_sold
	req.Statistic.External_sold = req2.External_sold
	return req, err

	//req := endpoint.PostRequest{}
	//err := json.NewDecoder(r.Body).Decode(&req)
	//return req, err
}

// encodePostResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodePostResponse(ctx context.Context, w http1.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http1.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http1.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http1.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
