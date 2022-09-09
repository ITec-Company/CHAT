package wsHub

import (
	"time"

	"github.com/gorilla/websocket"
	"itec.chat/pkg/logging"
)

const (
	maxMessageSize = 512
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9 / 10)
)

type client struct {
	hub    *Hub
	send   chan []byte
	conn   *websocket.Conn
	logger logging.Logger
}

func NewClient(hub *Hub, conn *websocket.Conn, logger logging.Logger) *client {
	client := &client{
		hub:    hub,
		send:   make(chan []byte, 256),
		conn:   conn,
		logger: logger,
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
				c.logger.Errorf("error occurred while reading message.err: ", err)
			}
			c.logger.Errorf("error occurred while reading message.err: ", err)
			break
		}
		c.hub.broadcast <- message
	}
}

func (c *client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.logger.Errorf("error occurred while setting write deadline.")
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				c.logger.Errorf("error occurred while getting writer. err: ", err)
				return
			}
			_, err = w.Write(message)
			if err != nil {
				c.logger.Errorf("error occurred while writing message. err: ", err)
				return
			}

			for i := 0; i < len(c.send); i++ {
				w.Write(<-c.send)
			}

			err = w.Close()
			if err != nil {
				c.logger.Errorf("error occurred while closing writer. err: ", err)
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			err := c.conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				c.logger.Errorf("error occurred while ping. err: ", err)
				return
			}
		}
	}

}
