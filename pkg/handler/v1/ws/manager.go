package ws

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

var GManager = &Manager{
	Clients:    make(map[string]*Client),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
}

// Client manager
type Manager struct {
	Clients    map[string]*Client
	Register   chan *Client
	Unregister chan *Client
}

type Client struct {
	Conn *websocket.Conn
	Send chan Message
}

func (c *Client) Reader() {
	defer c.Close()
	for {
		msg := new(Message)
		err := c.Conn.ReadJSON(msg)
		if err != nil {
			log.Println("read err:", err)
			return
		}
		log.Println("recv:", msg)
		// TODO: handle msg depend on type
		c.Send <- *msg

	}
}

func (c *Client) Writer() {
	defer c.Close()
	for msg := range c.Send {
		log.Println("accept:", msg)
		sendMsg, _ := json.Marshal(msg)
		c.Conn.WriteJSON(sendMsg)
	}
}

func (c *Client) Close() {
	defer c.Conn.Close()
	GManager.Unregister <- c
	close(c.Send)
}

func (m *Manager) Handle() {
	for {
		select {
		case c := <-m.Register:
			ip := c.Conn.RemoteAddr().String()
			m.Clients[ip] = c
		case c := <-m.Unregister:
			ip := c.Conn.RemoteAddr().String()
			if _, ok := m.Clients[ip]; !ok {
				continue
			} else {
				delete(m.Clients, ip)
			}
		}
	}
}
