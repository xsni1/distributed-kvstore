package main

import "fmt"

func main() {
	config, err := readConfig()
	if err != nil {
		fmt.Printf("main: %s", err)
		return
	}

	client := Client{config: *config}
	client.Run()
}
