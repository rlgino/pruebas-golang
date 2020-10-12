package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Calculator represents the service available over the network.
type Calculator int

func main() {
	s := new(Calculator)
	err := rpc.Register(s)
	if err != nil {
		log.Fatal(err)
	}

	rpc.HandleHTTP()

	l, err := net.Listen("tcp", "localhost:8767")
	if err != nil {
		log.Fatal("Error listen: ", err)
	}
	log.Println("Server started!")
	log.Println(l.Addr())
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatal("error serve: ", err)
	}
}
