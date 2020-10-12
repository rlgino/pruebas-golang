package main

import (
	"errors"
	"log"
)

// Addition stands as an exported method that execute a sum
// operation over the request args.
func (c *Calculator) Addition(rq *Request, rp *Response) error {
	log.Printf("Executing addition with args: %+v", rq)
	rp.Result = rq.A + rq.B
	return nil

}

// Subtraction stands as an exported method that execute a subtraction
// operation over the request args.
func (c *Calculator) Subtraction(rq *Request, rp *Response) error {
	log.Printf("Executing subtraction with args: %+v", rq)
	rp.Result = 0
	return errors.New("unsupported operation")

}
