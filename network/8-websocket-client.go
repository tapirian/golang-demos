package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	// 连接到 WebSocket 服务器
	serverURL := "ws://localhost:8080/ws"
	log.Printf("正在连接到 %s...", serverURL)

	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	defer conn.Close()

	log.Println("成功连接到 WebSocket 服务器")
	fmt.Println("输入消息并按回车发送，输入 'quit' 或 'exit' 退出")
	fmt.Println("----------------------------------------")

	// 处理中断信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	done := make(chan struct{})

	// 启动接收消息的 goroutine
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("读取消息失败:", err)
				return
			}
			fmt.Printf("\n收到: %s\n> ", string(message))
		}
	}()

	// 启动发送消息的 goroutine
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("> ")
			if !scanner.Scan() {
				break
			}

			message := strings.TrimSpace(scanner.Text())
			if message == "" {
				continue
			}

			// 检查退出命令
			if message == "quit" || message == "exit" {
				log.Println("正在关闭连接...")

				// 发送关闭消息
				err := conn.WriteMessage(websocket.CloseMessage,
					websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					log.Println("发送关闭消息失败:", err)
				}

				select {
				case <-done:
				case <-time.After(time.Second):
				}
				interrupt <- os.Interrupt
				return
			}

			// 发送消息
			err := conn.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("发送消息失败:", err)
				return
			}
		}
	}()

	// 等待中断信号或完成
	select {
	case <-done:
		log.Println("连接已关闭")
	case <-interrupt:
		log.Println("收到中断信号")

		// 优雅关闭
		err := conn.WriteMessage(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Println("发送关闭消息失败:", err)
			return
		}

		select {
		case <-done:
		case <-time.After(time.Second):
		}
	}
}
