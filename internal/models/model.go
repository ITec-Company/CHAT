package models

import "time"

const (
	OrderASC  Order = "asc"
	OrderDESC Order = "desc"
)

type Order string

type User struct {
	ID           int       `json:"id"`
	ProfileID    int       `json:"profile_id"`
	Name         string    `json:"name"`
	LastActivity time.Time `json:"last_activity"`
	Role         UserRole  `json:"role"`
	Status       Status    `json:"status"`
}

type CreateUser struct {
	ProfileID    int       `json:"profile_id"`
	Name         string    `json:"name"`
	LastActivity time.Time `json:"last_activity"`
	RoleID       int       `json:"role_id"`
	StatusID     int       `json:"status_id"`
}

type UpdateUser struct {
	ID           int       `json:"id"`
	ProfileID    int       `json:"profile_id"`
	Name         string    `json:"name"`
	LastActivity time.Time `json:"last_activity"`
	RoleID       int       `json:"role_id"`
	StatusID     int       `json:"status_id"`
}

type Chat struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	PhotoURL  string    `json:"photo_url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ChatByUser struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	PhotoURL string `json:"photo_url,omitempty"`
	IsAdmin  bool   `json:"is_admin"`
}

type CreateChat struct {
	Name     string `json:"name"`
	PhotoURL string `json:"photo_url,omitempty"`
}

type UpdateChat struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	PhotoURL string `json:"photo_url,omitempty"`
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
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateMessage struct {
	ChatID int    `json:"chat_id"`
	UserID int    `json:"user_id"`
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
	MessageID int    `json:"message_id"`
	URL       string `json:"url"`
}

type UpdateFile struct {
	ID  int    `json:"id,omitempty"`
	URL string `json:"url,omitempty"`
}

type SortOptions struct {
	SortBy         string                 `json:"sort_by,omitempty"`
	Order          Order                  `json:"order,omitempty"`
	FiltersAndArgs map[string]interface{} `json:"filters_and_args,omitempty"`
	Limit          uint64                 `json:"limit"`
	Page           uint64                 `json:"page"`
}
