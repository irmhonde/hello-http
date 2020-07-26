package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := 8080

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	fmt.Printf("listening on port %v", port)

	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
