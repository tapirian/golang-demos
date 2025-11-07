package main

import "net/http"

func main() {
	svc := NewAddService()
	endpoint := MakeSumEndpoint(svc)
	handler := MakeHTTPHandler(endpoint)

	http.ListenAndServe(":8080", handler)
}
