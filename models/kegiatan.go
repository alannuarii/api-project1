package models

type Kegiatan struct {
    ID          		int    `db:"id"`
    NamaKegiatan     	string `db:"nama_kegiatan"`
    PIC       			string `db:"pic"`
    Target   			string `db:"target"`
	Kode        		string `db:"kode"`
}