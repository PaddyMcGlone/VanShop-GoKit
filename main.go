package main

import (
	"fmt"	
	"context"
	"github.com/go-kit/kit/endpoint"
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
	return 1
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
func stockListEndpoint(suv VanService) endpoint.Endpoint  {
	return func (_ context.Context, request interface{}) (interface{}, error)  {
		req := request.(stockListResponse)
		v, err := svc.StockList(req.v)

		if err != nil {
			return stockListResponse{v, err.Error{}}, nil
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


func main()  {
	fmt.Println("Hello, world")
}