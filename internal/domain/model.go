package domain

import "time"

type User struct {
}

type Chat struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	PhotoURL  string    `json:"photo-url,omitempty"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
}

type CreateChat struct {
	Name     string `json:"name"`
	PhotoURL string `json:"photo-url,omitempty"`
}

type UpdateChat struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	PhotoURL string `json:"photo-url,omitempty"`
}

type Status struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CreateStatus struct {
	Name string `json:"name"`
}

type UpdateStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserRole struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CreateUserRole struct {
	Name string `json:"name"`
}

type UpdateUserRole struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	ID        int       `json:"id,omitempty"`
	Chat      Chat      `json:"chat"`
	User      User      `json:"user"`
	Body      string    `json:"body,omitempty"`
	CreatedAt time.Time `json:"created-at"`
	UpdatedAt time.Time `json:"updated-at"`
}

type CreateMessage struct {
	ChatID int    `json:"chat-id"`
	UserID int    `json:"user-id"`
	Body   string `json:"body"`
}

type UpdateMessage struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}

type File struct {
	ID      int     `json:"id,omitempty"`
	Message Message `json:"message"`
	URL     string  `json:"url,omitempty"`
}

type CreateFile struct {
	MessageID int    `json:"message-id"`
	URL       string `json:"url"`
}

type UpdateFile struct {
	ID  int    `json:"id,omitempty"`
	URL string `json:"url,omitempty"`
}
