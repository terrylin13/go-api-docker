package ws

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = &websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Handle(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Upgrade err:", err)
		return
	}

	client := &Client{
		Conn: conn,
		Send: make(chan Message),
	}

	GManager.Register <- client
	go client.Reader()
	go client.Writer()
}
