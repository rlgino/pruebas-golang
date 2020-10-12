package main

import (
	"log"
	"net/rpc"
)

// Calculator struct to wrapper RPC client
type Calculator struct {
	rpc *rpc.Client
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8767")
	if err != nil {
		log.Fatal("Dialing:", err)
	}

	// Build a new instance
	c := &Calculator{rpc: client}

	result, err := c.addition(4, 5)
	if err != nil {
		log.Println("Addition error: " + err.Error())
	} else {
		log.Printf("Addition result: %f", result)
	}

	result, err = c.substraction(4, 5)
	if err != nil {
		log.Println("Subtraction error: " + err.Error())
	} else {
		log.Printf("Subtraction result: %f", result)
	}
}

func (c *Calculator) addition(a, b float64) (float64, error) {
	req := Request{
		A: a,
		B: b,
	}
	var response Response
	err := c.rpc.Call("Calculator.Addition", req, &response)
	return response.Result, err
}

func (c *Calculator) substraction(a, b float64) (float64, error) {
	req := Request{
		A: a,
		B: b,
	}
	var response Response
	err := c.rpc.Call("Calculator.Subtraction", req, &response)
	return response.Result, err
}

// Request represents the data args for the service.
type Request struct {
	A, B float64
}

// Response represents the data results from the service.
type Response struct {
	Result float64
}
