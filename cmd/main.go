package main

import "calculator/application/http"

func main() {
	server, err := http.NewServer("localhost", 8080)
	if err != nil {
		panic(err)
	}

	server.ListenAndServe()
}
