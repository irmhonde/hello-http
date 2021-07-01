package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	port := 8080

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := strings.TrimPrefix(r.URL.Path, "/")
		if len(name) == 0 {
			name = "world"
		}

		_, err := w.Write([]byte(fmt.Sprintf("Hello world, %s!", strings.Title(name))))
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
