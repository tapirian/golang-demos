package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 创建路由
	mux := http.NewServeMux()

	// 注册路由
	registerRoutes(mux)

	// 创建服务器
	httpServer := &http.Server{
		Addr:    "127.0.0.1:9000",
		Handler: mux,
	}

	// 启动HTTP服务器
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	select {}
}
func registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/user", GetUser)
	mux.HandleFunc("/api/role", GetRole)
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "user")
}

func GetRole(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "role")
}
