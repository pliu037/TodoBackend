package controllers

/*
import (
	"TodoBackend/models"
	"TodoBackend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetEpics(c *gin.Context) {
	id := c.Params.ByName("id")
	var epics []models.Epic
	_, err := models.Dbmap.Select(&epics, "SELECT * FROM Epic WHERE id IN (SELECT epicid FROM EpicUserMap WHERE userid=?)", id)

	if err == nil {
		c.JSON(http.StatusOK, epics)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Can't find associated epics"})
	}
}

func PostEpic(c *gin.Context) {
	id := c.Params.ByName("id")
	if UserExists(id) {
		var epic models.Epic
		c.Bind(&epic)

		if epic.Name != "" {

			if insert, _ := models.Dbmap.Exec(`INSERT INTO Epic (name) VALUES (?)`, epic.Name); insert != nil {
				epic_id, err := insert.LastInsertId()
				if err == nil {
					models.Dbmap.Exec(`INSERT INTO EpicUserMap (userid, epicid) VALUES (?, ?)`, id, epic_id)
					content := &models.Epic{
						Id:   epic_id,
						Name: epic.Name,
					}
					c.JSON(http.StatusCreated, content)
				} else {
					utils.CheckErr(err, "Insert epic failed")
				}
			}

		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Field(s) is(are) empty"})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "User does not exist"})
	}
}

func UpdateEpic(c *gin.Context) {
	id := c.Params.ByName("id")
	epic_id := c.Params.ByName("epicid")

	if CheckEpicOwnedByUser(id, epic_id) {
		var epic models.Epic
		err := models.Dbmap.SelectOne(&epic, "SELECT * FROM Epic WHERE id=?", epic_id)

		if err == nil {
			var json models.Epic
			c.Bind(&json)

			epic := models.Epic{
				Id:   epic.Id,
				Name: json.Name,
			}

			if epic.Name != "" {
				_, err = models.Dbmap.Update(&epic)

				if err == nil {
					c.JSON(200, epic)
				} else {
					utils.CheckErr(err, "Updated epic failed")
				}

			} else {
				c.JSON(422, gin.H{"error": "Field(s) is(are) empty"})
			}

		} else {
			c.JSON(404, gin.H{"error": "Epic not found"})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Epic not owned by you"})
	}
}

func DeleteEpic(c *gin.Context) {
	id := c.Params.ByName("id")
	epic_id := c.Params.ByName("epicid")

	if CheckEpicOwnedByUser(id, epic_id) {
		var mapping models.EpicUserMap
		err := models.Dbmap.SelectOne(&mapping, "SELECT id FROM EpicUserMap WHERE userid=? AND epicid=?", id, epic_id)

		if err == nil {
			_, err = models.Dbmap.Delete(&mapping)

			if err == nil {
				c.JSON(http.StatusOK, gin.H{"id #" + epic_id: " deleted"})
				RemoveUnownedEpic(mapping.EpicID)
			} else {
				utils.CheckErr(err, "Delete epic failed")
			}

		} else {
			c.JSON(404, gin.H{"error": "Epic not found"})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Epic not owned by you"})
	}
}*/
