package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	io "github.com/jscastelblancoh/statistic_service/statistic/pkg/io"
	service "github.com/jscastelblancoh/statistic_service/statistic/pkg/service"
)

// GetRequest collects the request parameters for the Get method.
type GetRequest struct{}

// GetResponse collects the response parameters for the Get method.
type GetResponse struct {
	T   []io.Statistic `json:"t"`
	Err error          `json:"err"`
}

// MakeGetEndpoint returns an endpoint that invokes Get on the service.
func MakeGetEndpoint(s service.StatisticService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		t, err := s.Get(ctx)
		return GetResponse{
			Err: err,
			T:   t,
		}, nil
	}
}

// Failed implements Failer.
func (r GetResponse) Failed() error {
	return r.Err
}

// GetbyIdRequest collects the request parameters for the GetbyId method.
type GetbyIdRequest struct {
	Id string `json:"id"`
}

// GetbyIdResponse collects the response parameters for the GetbyId method.
type GetbyIdResponse struct {
	T   []io.Statistic `json:"t"`
	Err error          `json:"err"`
}

// MakeGetbyIdEndpoint returns an endpoint that invokes GetbyId on the service.
func MakeGetbyIdEndpoint(s service.StatisticService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetbyIdRequest)
		t, err := s.GetbyId(ctx, req.Id)
		return GetbyIdResponse{
			Err: err,
			T:   t,
		}, nil
	}
}

// Failed implements Failer.
func (r GetbyIdResponse) Failed() error {
	return r.Err
}

// PutRequest collects the request parameters for the Put method.
type PutRequest struct {
	Id string `json:"id"`
}

// PutResponse collects the response parameters for the Put method.
type PutResponse struct {
	T   io.Statistic `json:"t"`
	Err error        `json:"err"`
}

// MakePutEndpoint returns an endpoint that invokes Put on the service.
func MakePutEndpoint(s service.StatisticService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PutRequest)
		t, err := s.Put(ctx, req.Id)
		return PutResponse{
			Err: err,
			T:   t,
		}, nil
	}
}

// Failed implements Failer.
func (r PutResponse) Failed() error {
	return r.Err
}

// DeleteRequest collects the request parameters for the Delete method.
type DeleteRequest struct {
	Id string `json:"id"`
}

// DeleteResponse collects the response parameters for the Delete method.
type DeleteResponse struct {
	Err error `json:"err"`
}

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the service.
func MakeDeleteEndpoint(s service.StatisticService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		err := s.Delete(ctx, req.Id)
		return DeleteResponse{Err: err}, nil
	}
}

// Failed implements Failer.
func (r DeleteResponse) Failed() error {
	return r.Err
}

// PostRequest collects the request parameters for the Post method.
type PostRequest struct {
	Statistic io.Statistic `json:"statistic"`
}

// PostResponse collects the response parameters for the Post method.
type PostResponse struct {
	T   io.Statistic `json:"t"`
	Err error        `json:"err"`
}

// MakePostEndpoint returns an endpoint that invokes Post on the service.
func MakePostEndpoint(s service.StatisticService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostRequest)
		t, err := s.Post(ctx, req.Statistic)
		return PostResponse{
			Err: err,
			T:   t,
		}, nil
	}
}

// Failed implements Failer.
func (r PostResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Get implements Service. Primarily useful in a client.
func (e Endpoints) Get(ctx context.Context) (t []io.Statistic, err error) {
	request := GetRequest{}
	response, err := e.GetEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetResponse).T, response.(GetResponse).Err
}

// GetbyId implements Service. Primarily useful in a client.
func (e Endpoints) GetbyId(ctx context.Context, id string) (t []io.Statistic, err error) {
	request := GetbyIdRequest{Id: id}
	response, err := e.GetbyIdEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetbyIdResponse).T, response.(GetbyIdResponse).Err
}

// Put implements Service. Primarily useful in a client.
func (e Endpoints) Put(ctx context.Context, id string) (t io.Statistic, err error) {
	request := PutRequest{Id: id}
	response, err := e.PutEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PutResponse).T, response.(PutResponse).Err
}

// Delete implements Service. Primarily useful in a client.
func (e Endpoints) Delete(ctx context.Context, id string) (err error) {
	request := DeleteRequest{Id: id}
	response, err := e.DeleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteResponse).Err
}

// Post implements Service. Primarily useful in a client.
func (e Endpoints) Post(ctx context.Context, statistic io.Statistic) (t io.Statistic, err error) {
	request := PostRequest{Statistic: statistic}
	response, err := e.PostEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostResponse).T, response.(PostResponse).Err
}
