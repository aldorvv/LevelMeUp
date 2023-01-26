package main

import (
	"github.com/gin-gonic/gin"
	"github.com/levelmeup/levelmeup/internal/views"
)

func main() {
	router := gin.Default()
	router.GET("/skills", views.GetSkills)
	router.GET("/skills/:id", views.GetSkill)

	router.GET("/priorityLevels", views.GetPriorityLevels)
	router.GET("/priorityLevels/:id", views.GetPriorityLevel)

	router.GET("/quests", views.GetQuests)
	router.GET("/quests/:id", views.GetQuest)

	router.Run("localhost:8080")
}
