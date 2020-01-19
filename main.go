package main

import (
	"fmt"	
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
	Err string `json:"err,omitempty"`
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}



func main()  {
	fmt.Println("Hello, world")
}