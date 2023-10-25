package main

import (
	"fmt"
	"net/http"
)

// Responsible for selecting correct server node
// for partitioning purposes using consistent hashing method
type Client struct {
	config Config
}

func (c Client) Run() {
	for _, s := range c.config.Servers {
		fmt.Println(s)
	}
	http.ListenAndServe("127.0.0.1:9000", c)
}

func (c Client) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request")
}
