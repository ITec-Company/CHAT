package wsHub

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	maxMessageSize = 512
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9 / 10)
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type client struct {
	hub  *Hub
	send chan []byte
	conn *websocket.Conn
}

func NewClient(hub *Hub, conn *websocket.Conn) *client {
	client := &client{
		hub:  hub,
		send: make(chan []byte, 256),
		conn: conn,
	}
	client.hub.register <- client

	return client
}

func (c *client) ReadPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(appData string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Print("error occurred while reading message.err: ", err)
			}
			log.Print("breack read pump ", err)

			break
		}
		log.Print("READIN MESSGAE ", message)

		//message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- message
	}
}

func (c *client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	select {
	case message, ok := <-c.send:
		log.Print("clients WritePump ", message)

		c.conn.SetWriteDeadline(time.Now().Add(writeWait))
		if !ok {
			log.Print("error occurred while getting message.")
			c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		w, err := c.conn.NextWriter(websocket.TextMessage)
		if err != nil {
			log.Print("error occurred while getting message. err: ", err)
			return
		}
		_, err = w.Write(message)
		if err != nil {
			log.Print("error occurred while writing message. err: ", err)
			return
		}

		//n := len(c.send)
		log.Print("clients WritePump 2", message)

		for i := 0; i < len(c.send); i++ {
			log.Print("RANGE SEND WRITEPUM ", message)
			w.Write(newline)
			w.Write(<-c.send)
		}

		log.Print("clients WritePump 3", message)

		err = w.Close()
		if err != nil {
			log.Print("error occurred while closing writer. err: ", err)
			return
		}

	case <-ticker.C:
		c.conn.SetWriteDeadline(time.Now().Add(writeWait))
		err := c.conn.WriteMessage(websocket.PingMessage, nil)
		if err != nil {
			log.Print("error occurred while ping. err: ", err)
			return
		}
	}

}
