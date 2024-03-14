package controllers

import (
	"api-project1/db"
	"api-project1/models"

	// "fmt"
	"github.com/gin-gonic/gin"
)


func GetNotulen(c *gin.Context) {
    db := db.DB

    // Agenda 
    agenda := []models.Agenda{}

    query := `SELECT * FROM agenda WHERE nama_agenda = $1 ORDER BY tanggal DESC, id DESC LIMIT 1`

    err := db.Select(&agenda, query, "Daily Meeting")
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    if len(agenda) == 0 {
        c.JSON(404, gin.H{"error": "Agenda tidak ditemukan"})
        return
    }

    kode := agenda[0].Kode

    // Daftar Hadir 
    daftarHadir := []models.DaftarHadir{}

    query = `SELECT * FROM daftar_hadir WHERE kode = $1`

    err = db.Select(&daftarHadir, query, kode)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    // Pembangkit 
    pembangkit := []models.Engine{}

    query = `SELECT * FROM pembangkit WHERE kode = $1`

    err = db.Select(&pembangkit, query, kode)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    // BBM 
    tp := []models.Fuel{}
    th := []models.Fuel{}

    query = `SELECT * FROM bbm_tp WHERE kode = $1`

    err = db.Select(&tp, query, kode)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    query = `SELECT * FROM bbm_th WHERE kode = $1`

    err = db.Select(&th, query, kode)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    // k3kl 
    k3kl := []models.Info{}
    adm := []models.Info{}

    query = `SELECT * FROM k3kl WHERE kode = $1`

    err = db.Select(&k3kl, query, kode)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    query = `SELECT * FROM adm WHERE kode = $1`

    err = db.Select(&adm, query, kode)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    // Kegiatan 
    kegiatan := []models.Kegiatan{}

    query = `SELECT * FROM kegiatan WHERE kode = $1`

    err = db.Select(&kegiatan, query, kode)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    data := map[string]interface{}{
        "agenda": agenda,
        "daftar_hadir": daftarHadir,
        "pembangkit": pembangkit,
        "tp": tp,
        "th": th,
        "k3kl": k3kl,
        "adm": adm,
        "kegiatan": kegiatan,
    }

    c.JSON(200, gin.H{"message": "Data berhasil dikirim", "data": data})

}