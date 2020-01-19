package main

import (
	"context"
	"errors"
	"strings"
)

// Interface for the van service
type VanService interface {
	StockList() (Van, error)
	Count() int
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

func (vanService) StockList() (Van, error) {
	return Van {
		make: "Volkswagen",
		model: "T5",
		length: "SWB",
		engine: 2.0,
		bhp: 140
	}, nil
}

func (vanService) Count () int {
	// Will eventually do a count to the db instance

	return 1
}


func main()  {
	fmt.Println("Hello, world")
}