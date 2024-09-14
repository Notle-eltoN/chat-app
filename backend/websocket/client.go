package websocket

import (
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

func (c *Client) ReadMessages() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println("Read error: ", err)
			break
		}
		hub.Broadcast <- message
	}
}

func (c *Client) WriteMessages() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Send
		if !ok {
			return
		}

		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Write error: ", err)
			return
		}
	}
}
