package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/AndrewO777/BandPracticeAPI/models"
	"github.com/AndrewO777/BandPracticeAPI/db"
)

func RegisterSongRoutes(r *gin.Engine) {
	r.GET("/songs", GetSongs)
}

func GetSongs(c *gin.Context) {
	rows, err := db.DB.Query("SELECT * FROM LearnedSongs")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var song models.Song
		if err := rows.Scan(&song.ID, &song.SongName, &song.BandName, &song.PlayCount, &song.CurrentConfidence, &song.LastPlayed); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		songs = append(songs, song)
	}
	c.JSON(http.StatusOK, songs)
}
