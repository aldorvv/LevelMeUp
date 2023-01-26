package main

import (
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)

type Skill struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var DefaultSkills = []Skill{
	{ID: 1, Name: "COOKING", IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, Name: "EXCERCISING", IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 3, Name: "READING", IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 4, Name: "CODING", IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 5, Name: "PAINTING", IsActive: true, CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func getSkills(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, DefaultSkills)
}

func main() {
	router := gin.Default()
	router.GET("/skills", getSkills)

	router.Run("localhost:8080")
}
