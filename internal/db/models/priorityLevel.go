package models

import "time"

type PriorityLevel struct {
	ID    int    `json:"id"`
	Name  string `json:"levelName"`
	Value string `json:"value"`

	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
