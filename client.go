package main

import (
	"fmt"
	"hash/fnv"
	"net/http"
)

// Responsible for selecting correct server node
// for partitioning purposes using consistent hashing method
type Client struct {
	config Config
}

func (c Client) Run() {
	http.ListenAndServe("127.0.0.1:9000", c)
}

func (c Client) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request")
}

func (c Client) hash(val string) uint64 {
	h := fnv.New64()
	h.Write([]byte(val))

	return h.Sum64()
}
