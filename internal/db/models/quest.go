package models

import "time"

type Quest struct {
	ID       int           `json:"id"`
	Plevel   PriorityLevel `json:"priority"`
	Name     string        `json:"levelName"`
	Expoints int           `json:"Expoints"`

	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
