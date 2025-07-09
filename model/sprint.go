package model

import "time"

type Sprint struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Goals     string    `json:"goals,omitempty"`
	Notes     string    `json:"notes,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

