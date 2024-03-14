package models

type Fuel struct {
    ID      		int    `db:"id"`
    KodeTangki    	string `db:"kode_tangki"`
    Pengukuran   	string `db:"pengukuran"`
	Kode    		string `db:"kode"`
}