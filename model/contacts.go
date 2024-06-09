package model

import "time"

type Contacts struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InsertContacts struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateContacts struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
