package main

import (	
	"context"
	"github.com/go-kit/kit/endpoint"
	"encoding/json"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

// Interface for the van service
type VanService interface {
	StockList(string) (Van, error)
	Count(string) int
}

// Create a van object
type Van struct {
	make string
	model string
	length string
	engine float64
	bhp int
}

// Implement the interface
type vanService struct{}

func (vanService) StockList(s string) (Van, error) {
	return Van {
		make: "Volkswagen",
		model: "T5",
		length: "SWB",
		engine: 2.0,
		bhp: 140 }, nil
}

func (vanService) Count (s string) int {
	// Will eventually do a count to the db instance
	return 100
}


// Adding the Request and Response structs
// - Every method is modelled as an RPC call
type stockListRequest struct {
	S string `json:"s"`
}

type stockListResponse struct {
	v Van `json:"v"`
	Err string `json:"err,omitemptyŸ"`
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

//Adding the Endpoints
func makeStockListEndpoint(svc VanService) endpoint.Endpoint  {
	return func (_ context.Context, request interface{}) (interface{}, error)  {
		req := request.(stockListRequest)
		v, err := svc.StockList(req.S)

		if err != nil {
			return stockListResponse{v, err.Error()}, nil
		}
		return stockListResponse{v, ""}, nil
	}
}

func makeCountEndpoint(svc VanService) endpoint.Endpoint {
	return func (_ context.Context, request interface{}) (interface{}, error)  {
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{v}, nil
	}
}

//Adding transports
func main(){
	svc := vanService{}

	stockListHandler := httptransport.NewServer(
		makeStockListEndpoint(svc),
		decodeStockListRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/stocklist", stockListHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeStockListRequest(_ context.Context, r *http.Request) (interface{}, error)  {
	var request stockListRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil	
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error)  {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, Response interface{}) error{
	return json.NewEncoder(w).Encode(Response)
}