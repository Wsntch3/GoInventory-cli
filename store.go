package main

import (
	"errors"
	"time"
)

type Store struct{
	Barang []barang 	`json:"barang"`
	Transaction []transaction	`json:"Transaction"`
	NewID int			`json:"NewID"`
	Newidtransaksi int	`json:"NewIDtransaksi"`
}

func DefaultStore () *Store{
	return &Store{
		Barang: []barang{},
		Transaction: []transaction{},
		NewID: 1,
		Newidtransaksi: 1,
	}
}

func(s *Store) tambahbarang(b barang){
	b.Id = s.NewID
	s.NewID++
	s.Barang = append(s.Barang, b)
}

func (s *Store) hapusbarang(id int) error {
	for i := range s.Barang {
		if s.Barang[i].Id == id {
			s.Barang = append(s.Barang[:i], s.Barang[i+1:]...)
			return nil
		}
	}
	return errors.New("BARANG DENGAN ID TIDAK DITEMUKAN")
}


func (b *barang) kurangistok(jumlah int) error{
	if jumlah > b.Stok{
		return errors.New("Stok Kurang")
	}
	b.Stok -= jumlah
	return nil
}

func (s *Store) transaksi(idBarang, jumlah int) (*transaction, error){
	b, err := s.caribarang(idBarang)
	if err != nil{
		return nil,errors.New("ID TIDAK DITEMUKAN")
	}

	if err := b.kurangistok(jumlah);err != nil{
		return nil, err
	}

	items := itemtransaction{
		IDbarang: b.Id,	
		Nama: b.Nama,
		HargaJual: int(b.Harga),
		Jumlah: jumlah,
	}

	t := transaction{
		Idtransaksi: s.Newidtransaksi,
		Waktu: time.Now(),
		Items: []itemtransaction{items},
		Total: float64(jumlah)*b.Harga,
	}

	s.Newidtransaksi++
	s.Transaction = append(s.Transaction, t)
	return &t, err
}

func (s *Store) caribarang(id int) (*barang, error){
	for i := range s.Barang{
		if s.Barang[i].Id == id {
			return &s.Barang[i], nil
		}
	}
	return nil, errors.New("BARANG DENGAN ID TIDAK DITEMUKAN")
}
