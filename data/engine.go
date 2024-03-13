package data

type Engine struct {
	Unit  	string
	Mesin 	string
	Tipe  	string
	Seri  	string
	DTP   	int
	DMN   	int
	Status	string
}

var DataEngine = []Engine{
	{Unit: "1", Mesin: "Daihatsu", Tipe: "6PSHTc 26 Dm", Seri: "6265146", DTP: 500},
	{Unit: "2", Mesin: "Daihatsu", Tipe: "6PSHTc 26 D", Seri: "6263607", DTP: 500},
	{Unit: "3", Mesin: "SWD", Tipe: "8 FHD 240", Seri: "70590-1", DTP: 1000},
	{Unit: "4", Mesin: "SWD", Tipe: "8 FHD 240", Seri: "70590-2", DTP: 1000},
	{Unit: "5", Mesin: "MAK", Tipe: "8M 453 AK", Seri: "26877", DTP: 2500},
	{Unit: "6", Mesin: "MAK", Tipe: "8M 453 AK", Seri: "26878", DTP: 2500},
}