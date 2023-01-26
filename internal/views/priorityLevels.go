package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levelmeup/levelmeup/internal/db"
	"github.com/levelmeup/levelmeup/internal/db/models"
)

func GetPriorityLevels(c *gin.Context) {
	const query = "SELECT * FROM priorityLevels"
	repo := db.GetRepo()

	result, err := repo.Query(query)
	var levels []models.PriorityLevel

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var level models.PriorityLevel
		err = result.Scan(&level.ID, &level.Name, &level.Value, &level.IsActive, &level.CreatedAt, &level.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		levels = append(levels, level)
	}
	c.IndentedJSON(http.StatusOK, levels)
}

func GetPriorityLevel(c *gin.Context) {
	const query = "SELECT * FROM priorityLevels WHERE id=?"
	repo := db.GetRepo()
	id := c.Param("id")
	var level models.PriorityLevel

	result, err := repo.Query(query, id)

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err = result.Scan(&level.ID, &level.Name, &level.Value, &level.IsActive, &level.CreatedAt, &level.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
	}
	c.IndentedJSON(http.StatusOK, level)
}
