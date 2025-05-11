package api

import (
	"fmt"
	"net/http"
)

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "guest"
	}

	fmt.Fprintf(w, "Hello, %s! Selamat datang di toko Skibidi!", name)
}
