package models

type DaftarHadir struct {
    ID          int    `db:"id"`
    Nama     	string `db:"nama"`
    NID       	string `db:"nid"`
    Instansi   	string `db:"instansi"`
    Unit        string `db:"unit"`
    Jabatan  	string `db:"jabatan"`
	Kode        string `db:"kode"`
}

type AgendaDaftarHadir struct {
    Tanggal     		string `db:"tanggal"`
    Tempat       		string `db:"tempat"`
    NamaAgenda        	string `db:"nama_agenda"`
    Kode                string `db:"kode"`
}