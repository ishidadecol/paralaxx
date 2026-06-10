package entity

import "time"

type Entity struct {
	ID   string `json:"id"`
	Type string `json:"type"`

	DisplayName string `json:"display_name"`

	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
