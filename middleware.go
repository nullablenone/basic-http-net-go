package main

import (
	"fmt"
	"net/http"
)

func middlewareCreate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request masuk : ", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

func create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "request masuk dan sudah melalui middleware")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", middlewareCreate(http.HandlerFunc(create)))
	http.ListenAndServe(":80", mux)
}
