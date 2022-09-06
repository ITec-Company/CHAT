package chatsession

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type chatSession struct {
	ws        *websocket.Conn
	userName  string
	chatId    int
	broadcast chan []byte
}

var Peers = make(map[string]*websocket.Conn)

var ChatSessions = make(map[string]*chatSession)

func NewChatSession(userName string, chatId int, ws *websocket.Conn) *chatSession {

	return &chatSession{
		ws:        ws,
		userName:  userName,
		chatId:    chatId,
		broadcast: make(chan []byte),
	}
}

func (s *chatSession) Start() {
	Peers[s.userName] = s.ws

	s.greetings("Hello to chat")

	go func() {
		log.Println("Client connected:", s.ws.RemoteAddr().String())

		for {
			_, msg, err := s.ws.ReadMessage()
			if err != nil {
				_, ok := err.(*websocket.CloseError)
				if ok {
					log.Println("Connection closed err:", err)
					s.disconnect()
				}
				log.Println("Connection closed err:", err)
				return
			}
			s.broadcast <- msg
		}
	}()

	go s.handleMessages()
}

func (s *chatSession) greetings(msg string) {
	err := s.ws.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		log.Println("failed to write message", err)
	}
}

func (s *chatSession) disconnect() {
	s.ws.WriteMessage(websocket.TextMessage, []byte("User left the chat"))

	s.ws.Close()

	delete(Peers, s.userName)
}

func (s *chatSession) handleMessages() {
	for {
		msg := <-s.broadcast
		for user, ws := range Peers {
			err := ws.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Printf("error: %v", err)
				s.ws.WriteMessage(websocket.TextMessage, []byte("User left the chat"))
				ws.Close()
				delete(Peers, user)
			}
		}
	}
}

func MapInfo() string {
	str := ""
	for key := range Peers {
		m := fmt.Sprint("\nKey:", key)
		str = str + m
	}
	return str
}
 