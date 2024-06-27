package controllers

import (
	"api-project1/db"
	"api-project1/models"
	"github.com/gin-gonic/gin"
)


func GetAllAbsensi(c *gin.Context){
	db := db.DB

	daftar_hadir := []models.AgendaDaftarHadir{}

	query := `SELECT a.tanggal, a.tempat, a.nama_agenda, a.kode FROM agenda AS a JOIN daftar_hadir AS b ON a.kode = b.kode GROUP BY a.tanggal, a.tempat, a.nama_agenda, a.kode`

	err := db.Select(&daftar_hadir, query)
	if err != nil{
		c.JSON(500, gin.H{"message": err.Error()})
        return
	}

	c.JSON(200, gin.H{"message": "Sukses", "data": daftar_hadir})
}


func GetAbsensi(c *gin.Context){
	db := db.DB

	kode := c.Param("kode")

	// Agenda 
	agenda := []models.Agenda{}

	query := `SELECT * FROM agenda WHERE kode = $1`

	err := db.Select(&agenda, query, kode)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

    // Daftar Hadir 
    daftarHadir := []models.DaftarHadir{}

    query = `SELECT * FROM daftar_hadir WHERE kode = $1`

    err = db.Select(&daftarHadir, query, kode)
    if err != nil {
        c.JSON(500, gin.H{"message": err.Error()})
        return
    }

	data := map[string]interface{}{
		"agenda" : agenda,
		"daftar_hadir" : daftarHadir,
	}

	c.JSON(200, gin.H{"message": "Sukses", "data": data})
}