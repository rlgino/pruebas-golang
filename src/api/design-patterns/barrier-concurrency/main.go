package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var endpoints = []string{
	"https://jsonplaceholder.typicode.com/posts/1",
	"https://jsonplaceholder.typicode.com/posts/2",
	"https://jsonplaceholder.typicode.com/posts/3",
	"https://jsonplaceholder.typicode.com/posts/4",
	"https://jsonplaceholder.typicode.com/posts/5",
	"https://jsonplaceholder.typicode.com/posts/7",
	"https://jsonplaceholder.typicode.com/posts/8",
	"https://jsonplaceholder.typicode.com/posts/9",
	"https://jsonplaceholder.typicode.com/posts/10",
}

var timeoutMilliseconds int = 5000

// BarrierResp response from jsonplaceholder
type BarrierResp struct {
	Err  error
	Resp string
}

func main() {
	defer func() { // Con esto capturaria la excepción
		if r := recover(); r != nil {
			fmt.Println("Catch exception: ", r)
		}
	}()

	responses := barrier(endpoints)
	fmt.Println(responses)
}

// Acá nos importa que el barrido se haga a cada URl que pasemos por párametro y nuestro objetivo será el de poder
// juntar esos datos
func barrier(endpoints []string) []BarrierResp {
	requestNumber := len(endpoints)
	in := make(chan BarrierResp, requestNumber) // make n channels of BarrierResponses
	defer close(in)

	responses := make([]BarrierResp, requestNumber) // make n channels of Barrier arrays responses
	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}

	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in // Acá no importa el orden, como es la misma cantadad pueden llegar desordenados
		if resp.Err != nil {
			fmt.Println("Error: ", resp.Err)
			hasError = true
		}
		responses[i] = resp
	}

	if hasError {
		panic("Invalid response")
	}

	return responses
}

func makeRequest(out chan<- BarrierResp, url string) {
	res := BarrierResp{}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) * time.Millisecond),
	}

	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}
	res.Resp = string(byt)
	out <- res
}
