package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levelmeup/levelmeup/internal/db"
	"github.com/levelmeup/levelmeup/internal/db/models"
)

func GetSkills(c *gin.Context) {
	const query = "SELECT * FROM skills"
	repo := db.GetRepo()

	result, err := repo.Query(query)
	var skills []models.Skill

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var skill models.Skill
		err = result.Scan(&skill.ID, &skill.Name, &skill.IsActive, &skill.CreatedAt, &skill.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		skills = append(skills, skill)
	}
	c.IndentedJSON(http.StatusOK, skills)
}

func GetSkill(c *gin.Context) {
	const query = "SELECT * FROM skills WHERE id=?"
	repo := db.GetRepo()
	id := c.Param("id")
	var skill models.Skill

	result, err := repo.Query(query, id)

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		err = result.Scan(&skill.ID, &skill.Name, &skill.IsActive, &skill.CreatedAt, &skill.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
	}
	c.IndentedJSON(http.StatusOK, skill)
}
