package views

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levelmeup/levelmeup/internal/db"
	"github.com/levelmeup/levelmeup/internal/db/models"
)

func getPriorityLevel(id int, repo *sql.DB) models.PriorityLevel {
	const queryLevel = "SELECT * FROM priorityLevels WHERE id=?"
	var level models.PriorityLevel

	result, err := repo.Query(queryLevel, id)

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err = result.Scan(&level.ID, &level.Name, &level.Value, &level.IsActive, &level.CreatedAt, &level.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
	}
	return level
}

func GetQuests(c *gin.Context) {
	const query = "SELECT * FROM quests"
	repo := db.GetRepo()
	var priorityID int

	result, err := repo.Query(query)
	var quests []models.Quest

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var q models.Quest
		err = result.Scan(&q.ID, &priorityID, &q.Name, &q.Expoints, &q.IsActive, &q.CreatedAt, &q.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}

		q.Plevel = getPriorityLevel(priorityID, repo)
		quests = append(quests, q)
	}

	c.IndentedJSON(http.StatusOK, quests)
}

func GetQuest(c *gin.Context) {
	const query = "SELECT * FROM quests WHERE id=?"
	repo := db.GetRepo()
	id := c.Param("id")
	var q models.Quest
	var priorityID int

	result, err := repo.Query(query, id)

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err = result.Scan(&q.ID, &priorityID, &q.Name, &q.Expoints, &q.IsActive, &q.CreatedAt, &q.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
	}

	q.Plevel = getPriorityLevel(priorityID, repo)
	c.IndentedJSON(http.StatusOK, q)
}
