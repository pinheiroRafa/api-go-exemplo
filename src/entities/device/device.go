package entities

import "time"

type Device struct {
	Id        string     `json:"id"`
	UserAgent string     `json:"user_agent"`
	CreatedAt time.Time  `json:"created_at"`
	LastUsed  *time.Time `json:"last_used"`
	UserId    string     `json:"user_id"`
}
