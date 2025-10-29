// endpoint.go
package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// 请求和响应结构
type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	Result int    `json:"result"`
	Err    string `json:"err,omitempty"`
}

func MakeSumEndpoint(svc AddService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(SumRequest)
		result, err := svc.Sum(ctx, req.A, req.B)
		if err != nil {
			return SumResponse{Result: 0, Err: err.Error()}, nil
		}
		return SumResponse{Result: result}, nil
	}
}
