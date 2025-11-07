// service.go
package main

import (
	"context"
)

// 服务层， 定义接口
type AddService interface {
	Sum(ctx context.Context, a, b int) (int, error)
}

// 实现接口
type addService struct{}

func (addService) Sum(ctx context.Context, a, b int) (int, error) {
	return a + b, nil
}

// 创建服务实例
func NewAddService() AddService {
	return addService{}
}
