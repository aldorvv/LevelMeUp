package models

import "time"

type Skill struct {
	ID   int    `json:"id"`
	Name string `json:"skillName"`

	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
