package main

// Request represents the data args for the service.
type Request struct {
	A, B float64
}

// Response represents the data results from the service.
type Response struct {
	Result float64
}

// Operations struct to expose at RPC server
type Operations interface {
	Addition(*Request, *Response) error
	Substraction(*Request, *Response) error
}
