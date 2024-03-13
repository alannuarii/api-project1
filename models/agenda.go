package models

type Agenda struct {
    ID          		int    `db:"id"`
    Tanggal     		string `db:"tanggal"`
    Waktu       		string `db:"waktu"`
    Tempat   			string `db:"tempat"`
    NamaAgenda        	string `db:"nama_agenda"`
    DasarPembahasan  	string `db:"dasar_pembahasan"`
    Notulis        		string `db:"notulis"`
	Foto        		string `db:"foto"`
	Kode        		string `db:"kode"`
}