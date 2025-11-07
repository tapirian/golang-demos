// transport.go
package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
)

// 请求编码
func decodeSumRequest(_ context.Context, r *http.Request) (any, error) {
	var req SumRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// 响应编码
func encodeSumResponse(_ context.Context, w http.ResponseWriter, response any) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// HTTP 处理器
func MakeHTTPHandler(endpoints endpoint.Endpoint) http.Handler {
	return kithttp.NewServer(
		endpoints,
		decodeSumRequest,
		encodeSumResponse,
	)
}
