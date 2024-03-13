package controllers

import (
	"api-project1/data"
	"api-project1/db"
	"api-project1/utils"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func PostNotulen(c *gin.Context){

	db := db.DB

	kode := c.Request.FormValue("kode")
	tanggal := c.Request.FormValue("tanggal")
	tempat := c.Request.FormValue("tempat")
	waktu := c.Request.FormValue("waktu")
	agenda := c.Request.FormValue("agenda")
	dasar := c.Request.FormValue("dasar")
	notulis := c.Request.FormValue("notulis")
	foto := c.Request.FormValue("foto")
	namaPegawai := c.Request.FormValue("namaPeg")
	namaOphar := c.Request.FormValue("namaOH")
	namaCs := c.Request.FormValue("namaCs")
	namaSecurity := c.Request.FormValue("namaSec")
	dmn := c.Request.FormValue("dmn")
	status := c.Request.FormValue("status")
	tp := c.Request.FormValue("tp")
	th := c.Request.FormValue("th")
	k3kl := c.Request.FormValue("k3kl")
	adm := c.Request.FormValue("adm")
	kegiatan := c.Request.FormValue("kegiatan")
	pic := c.Request.FormValue("pic")
	target := c.Request.FormValue("target")

	rename := fmt.Sprintf("%s.png", kode)

	// Agenda 
	queryAgenda := `INSERT INTO agenda (tanggal, waktu, tempat, nama_agenda, dasar_pembahasan, notulis, foto, kode) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	var agendaID int
	err := db.QueryRow(queryAgenda, tanggal, waktu, tempat, agenda, dasar, notulis, rename, kode).Scan(&agendaID)
	if err != nil{
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Daftar Hadir 
	namaPegawaiSlice := strings.Split(namaPegawai, ",")
	namaOpharSlice := strings.Split(namaOphar, ",")
	namaCsSlice := strings.Split(namaCs, ",")
	namaSecuritySlice := strings.Split(namaSecurity, ",")

	var person []string
	person = append(person, namaPegawaiSlice...)
	person = append(person, namaOpharSlice...)
	person = append(person, namaCsSlice...)
	person = append(person, namaSecuritySlice...)

	var daftarHadir []data.Person

	for _, nama := range person{
		for _, p := range data.DataPerson{
			if p.Nama == nama{
				daftarHadir = append(daftarHadir, p)
				break
			}
		}
	}
	queryDaftarHadir := `INSERT INTO daftar_hadir (nama, nid, instansi, unit, jabatan, kode) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	for _, p := range daftarHadir{
		var daftarHadirID int
		err = db.QueryRow(queryDaftarHadir, p.Nama, p.NID, p.Instansi, p.Unit, p.Jabatan, kode).Scan(&daftarHadirID)
		if err != nil{
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	// Pembangkit
	dmnSlice := strings.Split(dmn, ",")
	statusSlice := strings.Split(status, ",")

	var dmnIntSlice []int

	for _, s := range dmnSlice{
		i, _ := strconv.Atoi(s)
		dmnIntSlice = append(dmnIntSlice, i)
	}

	var kondisiPembangkit []data.Engine
	for i, e := range data.DataEngine{
		e.DMN = dmnIntSlice[i]
		e.Status = statusSlice[i]
		kondisiPembangkit = append(kondisiPembangkit, e)
	}

	queryPembangkit := `INSERT INTO pembangkit (unit, mesin, tipe, seri, dtp, dmn, status, kode) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	for _, p := range kondisiPembangkit{
		var pembangkitID int
		err = db.QueryRow(queryPembangkit, p.Unit, p.Mesin, p.Tipe, p.Seri, p.DTP, p.DMN, p.Status, kode).Scan(&pembangkitID)
		if err != nil{
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	// BBM 
	tpSlice := strings.Split(tp, ",")
	thSlice := strings.Split(th, ",")

	var tpFloatSlice []float32
	var thFloatSlice []float32

	for _, s := range tpSlice{
		f32, _ := strconv.ParseFloat(s, 32)
		tpFloatSlice = append(tpFloatSlice, float32(f32))
	}

	for _, s := range thSlice{
		f32, _ := strconv.ParseFloat(s, 32)
		thFloatSlice = append(thFloatSlice, float32(f32))
	}

	var dataPengukuranTP []data.Fuel
	for i, p := range data.TP{
		p.Pengukuran = tpFloatSlice[i]
		dataPengukuranTP = append(dataPengukuranTP, p)
	}

	var dataPengukuranTH []data.Fuel
	for i, p := range data.TH{
		p.Pengukuran = thFloatSlice[i]
		dataPengukuranTH = append(dataPengukuranTH, p)
	}

	queryTP := `INSERT INTO bbm_tp (kode_tangki, pengukuran, kode) VALUES ($1, $2, $3) RETURNING id`
	queryTH := `INSERT INTO bbm_th (kode_tangki, pengukuran, kode) VALUES ($1, $2, $3) RETURNING id`

	for _, p := range dataPengukuranTP{
		var tpID int
		err = db.QueryRow(queryTP, p.KodeTangki, p.Pengukuran, kode).Scan(&tpID)
		if err != nil{
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	for _, p := range dataPengukuranTH{
		var thID int
		err = db.QueryRow(queryTH, p.KodeTangki, p.Pengukuran, kode).Scan(&thID)
		if err != nil{
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	// K3Kl 
	k3klSlice := strings.Split(k3kl, ",")

	queryK3kl := `INSERT INTO k3kl (info, kode) VALUES ($1, $2) RETURNING id`

	for i := range k3klSlice{
		var k3klID int
		err = db.QueryRow(queryK3kl, k3klSlice[i], kode).Scan(&k3klID)
		if err != nil{
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	// Adm 
	admSlice := strings.Split(adm, ",")

	queryAdm := `INSERT INTO adm (info, kode) VALUES ($1, $2) RETURNING id`

	for i := range admSlice{
		var admID int
		err = db.QueryRow(queryAdm, admSlice[i], kode).Scan(&admID)
		if err != nil{
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	// Kegiatan 
	kegiatanSlice := strings.Split(kegiatan, ",")
	picSlice := strings.Split(pic, ",")
	targetSlice := strings.Split(target, ",")

	queryKegiatan := `INSERT INTO kegiatan (nama_kegiatan, pic, target, kode) VALUES ($1, $2, $3, $4) RETURNING id`

	for i := range kegiatanSlice{
		var kegiatanID int
		err = db.QueryRow(queryKegiatan, kegiatanSlice[i], picSlice[i], targetSlice[i], kode).Scan(&kegiatanID)
		if err != nil{
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	// Foto 
	destination := "static/img"
	outputPath := filepath.Join(destination,rename)
	err = utils.Base64topng(foto, outputPath)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Data berhasil dikirim"})
}


