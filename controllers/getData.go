package controllers

import (
	"api-project1/db"
	"api-project1/models"

	"github.com/gin-gonic/gin"
)


func GetNotulen(c *gin.Context) {
    db := db.DB

    agenda := []models.Agenda{}

    query := `SELECT * FROM agenda WHERE nama_agenda = 'Daily Meeting' ORDER BY tanggal DESC LIMIT 1`

    err := db.Select(&agenda, query)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "Data berhasil dikirim", "data": agenda})
}