package models

type Engine struct {
    ID      int    `db:"id"`
    Unit    string `db:"unit"`
    Mesin   string `db:"mesin"`
    Tipe   	string `db:"tipe"`
    Seri    string `db:"seri"`
    DTP 	int `db:"dtp"`
	DMN 	int `db:"dmn"`
	Status 	string `db:"status"`
	Kode    string `db:"kode"`
}