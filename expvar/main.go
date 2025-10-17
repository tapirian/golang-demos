package main

import (
	_ "expvar"
	"net/http"
)

func main() {
	go func() {
		mux := http.DefaultServeMux
		http.ListenAndServe(":8080", mux)
	}()
	select {}
}
