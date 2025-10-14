package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允许所有来源的跨域请求（生产环境需要更严格的配置）
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 客户端连接管理
type Client struct {
	conn *websocket.Conn
	send chan []byte
}

// 连接管理器
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			log.Printf("新客户端连接，当前在线: %d", len(h.clients))

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				log.Printf("客户端断开，当前在线: %d", len(h.clients))
			}

		case message := <-h.broadcast:
			// 广播消息给所有客户端
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// 读取客户端消息
func (c *Client) readPump(hub *Hub) {
	defer func() {
		hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("错误: %v", err)
			}
			break
		}

		log.Printf("收到消息: %s", message)

		// 广播消息给所有客户端
		response := fmt.Sprintf("服务器收到: %s", message)
		hub.broadcast <- []byte(response)
	}
}

// 发送消息给客户端
func (c *Client) writePump() {
	ticker := time.NewTicker(54 * time.Second)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// WebSocket 处理函数
func handleWebSocket(hub *Hub, w http.ResponseWriter, r *http.Request) {
	// 升级 HTTP 连接为 WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("升级失败:", err)
		return
	}

	client := &Client{
		conn: conn,
		send: make(chan []byte, 256),
	}
	hub.register <- client

	// 发送欢迎消息
	client.send <- []byte("欢迎连接到 WebSocket 服务器！")

	// 启动读写协程
	go client.writePump()
	go client.readPump(hub)
}

// 提供一个简单的 HTML 页面用于测试
func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html, err := os.ReadFile("websocket.html")
	if err != nil {
		w.Write([]byte("读取文件失败"))
		log.Printf("%v\n", err)
		return
	}
	w.Write(html)
}

func main() {
	hub := newHub()
	go hub.run()

	// 路由设置
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(hub, w, r)
	})

	addr := ":8080"
	log.Printf("WebSocket 服务器启动在 http://localhost%s", addr)
	log.Println("访问 http://localhost:8080 打开测试页面")

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}
