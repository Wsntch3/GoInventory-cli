package main

import (
	"fmt"
)
func inputbarang() barang {
	var b barang
	fmt.Println("Silahkan Masukkan barang apa yang ditambahkan: ")
	fmt.Print("Nama Barang: ")
	fmt.Scan(&b.Nama)
	fmt.Print("Harga Barang: ")
	fmt.Scan(&b.Harga)
	fmt.Print("Masukkan Stok: ")
	fmt.Scan(&b.Stok)
	fmt.Println("Barang dengan Nama", b.Nama,"Harga", b.Harga,"Sejumlah: ",b.Stok,"Telah ditambahkan")
	return b
}

func main(){
	store, err := loaddata("data/data.json")
	if err != nil {
		fmt.Println("GAGAL MEMUAT DATA....")
		return
	}
	for{
		fmt.Println("Selamat datang di inventory")
		fmt.Println("Pilih Layanan: ")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Hapus Barang")
		fmt.Println("3. Cari Barang")
		fmt.Println("4. Catat Penjualan")
		fmt.Println("5. Tampilkan laporan")
		fmt.Println("0. Exit")

		fmt.Print("Masukkan Pilihan Layanan: ")
		var c int
		fmt.Scan(&c)
		switch(c){
		case 0:
			fmt.Println("Operasi Selesai")
			return
		case 1:
			b := inputbarang()
			store.tambahbarang(b)
			if err := savedata("data/data.json", store);err != nil {
			fmt.Println("BARANG GAGAL DIMUAT....")
		}
		case 2:
			for _,b := range store.Barang{
				fmt.Println(b.Id,b.Nama,b.Harga)
			}
			fmt.Println("Masukkan id barang: ")
			var c int
			fmt.Scan(&c)
			if err = store.hapusbarang(c);err != nil{
				fmt.Println("ID Barang tidak ada")
			}
			if err := savedata("data/data.json", store);err != nil {
				fmt.Println("BARANG GAGAL DIMUAT....")
			}
			fmt.Println("Barang dengan ID ",c, "Telah Dihapus")
		case 3:
			fmt.Println("Masukkan ID barang: ")
			var d int
			fmt.Scan(&d)
			b,err := store.caribarang(d)
			if err != nil{
				fmt.Println("BARANG YANG DICARI TIDAK TERSEDIA")
			}
			fmt.Println("Barang Ditemukan: ",b.Id,b.Nama,b.Harga)
		case 4:
			var a int
			var b int
			fmt.Println("Masukkan ID barang: ")
			fmt.Scan(&a)
			fmt.Println("Masukkan Jumlah yang dijual: ")
			fmt.Scan(&b)
			t, err := store.transaksi(a,b)
			if err != nil{
				fmt.Println("Gagal melakukan transaksi")	
			}
			if err := savedata("data/data.json", store);err != nil {
				fmt.Println("BARANG GAGAL DIMUAT....")
			}
			fmt.Println("Transaksi dengan id ",t.Idtransaksi,"dengan Total Harga: ",t.Total,"Telah sukses")
		case 5:
			fmt.Println("TOTAL PENDAPATAN: ", store.totalpendapatan())
			fmt.Println("BARANG TERLARIS: ")
			for Id,items := range store.barangterlaris(){
				fmt.Println("ID","Total Items")
				fmt.Println(Id,items)
			}

		default:
			fmt.Println("Layanan yang dipilih tidak tersedia")
		}
	}
}