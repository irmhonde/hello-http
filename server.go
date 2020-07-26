package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := 8080

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello, world!"))
		if err != nil {
			panic(err)
		}
	})

	_, err := fmt.Printf("listening on port %v", port)
	if err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		panic(err)
	}
}
