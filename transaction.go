package main

import "time"

type itemtransaction struct{
	IDbarang 	int 	`json:"IDbarang"`
	Nama 		string 	`json:"Nama"`
	HargaJual 	int 	`json:"HargaJual"`
	Jumlah 		int 	`json:"Jumlah"`
}

type transaction struct {
	Idtransaksi int 				`json:"Idtransaksi"`
	Waktu 		time.Time			`json:"Waktu"`
	Items 		[]itemtransaction 	`json:"Items"`
	Total		float64				`json:"Total"`
}