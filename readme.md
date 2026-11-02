# Inventory CLI

Aplikasi terminal untuk mengelola inventori dan penjualan toko, dibuat dengan **Go** tanpa library eksternal. Data disimpan dalam file JSON sehingga tidak hilang saat program ditutup.

Proyek ini dibuat sebagai latihan dasar Go (struct, slice, map, pointer, method, error handling, file, dan waktu) dan akan dikembangkan menjadi REST API dengan PostgreSQL.

## Fitur

- **Tambah barang**: nama, harga, dan stok. ID dibuat otomatis.
- **Hapus barang**: berdasarkan ID.
- **Cari barang**: berdasarkan ID.
- **Catat penjualan**: stok berkurang otomatis, total harga dihitung, dan transaksi tercatat lengkap dengan waktu. Penjualan ditolak jika stok tidak cukup.
- **Laporan**: total pendapatan dan barang terlaris.
- **Penyimpanan permanen**: semua data disimpan ke `data/data.json`.
- **Menu berulang**: program kembali ke menu utama setelah setiap operasi sampai pengguna memilih keluar.

## Teknologi

- Go (hanya standard library: `fmt`, `os`, `encoding/json`, `errors`, `time`)

## Struktur Proyek

```
inventory_cli/
├── main.go          # menu dan interaksi dengan pengguna (input/output)
├── barang.go        # struct barang dan method-nya
├── transaction.go   # struct transaksi dan logika penjualan
├── store.go         # penyimpanan data di memori + operasi tambah/hapus/cari
├── file.go          # baca dan tulis data.json (loaddata, savedata)
├── report.go        # fungsi laporan (pendapatan, barang terlaris)
├── data/
│   └── data.json    # file penyimpanan data
├── go.mod
└── README.md
```

Pembagian tugas tiap file:

| File | Tanggung jawab |
|---|---|
| `main.go` | Satu-satunya tempat `fmt.Print` dan `fmt.Scan` |
| `store.go` | Mengubah data di memori, tanpa akses file |
| `file.go` | Satu-satunya tempat yang mengakses file dan JSON |
| `report.go` | Hanya membaca dan menghitung, tidak mengubah data |

Pemisahan ini membuat penyimpanan mudah diganti (misalnya ke PostgreSQL) tanpa mengubah logika lainnya.

## Cara Menjalankan

Pastikan Go sudah terpasang (`go version`).

```bash
git clone https://github.com/Wsntch3/inventory-cli.git
cd inventory-cli
mkdir -p data
go run .
```

> Gunakan `go run .` (dengan titik), bukan `go run main.go`, agar semua file ikut dimuat.

## Contoh Tampilan

```
Selamat datang di inventory

Pilih Layanan:
1. Tambah Barang
2. Hapus Barang
3. Cari Barang
4. Catat Penjualan
5. Tampilkan laporan
0. Keluar
Masukkan Pilihan Layanan: 1
Nama Barang: Gelas
Harga Barang: 20000
Masukkan Stok: 10
Barang dengan Nama Gelas Harga 20000 Sejumlah: 10 Telah ditambahkan
```

## Format Data

Data disimpan dalam satu file `data/data.json`:

```json
{
  "barang": [
    { "id": 1, "nama": "Gelas", "harga": 20000, "stok": 8 }
  ],
  "transaksi": [
    {
      "id": 1,
      "waktu": "2026-11-02T10:00:00+07:00",
      "items": [
        { "id_barang": 1, "nama": "Gelas", "harga": 20000, "jumlah": 2 }
      ],
      "total": 40000
    }
  ],
  "NewID": 2,
  "NewTransaksiID": 2
}
```

Nama dan harga barang disalin ke dalam transaksi, sehingga nota lama tidak berubah walaupun harga barang diubah atau barangnya dihapus.

## Hal yang Dipelajari

- Struct, method, dan receiver pointer
- Slice dan map
- Error handling (mengembalikan `error`, bukan `panic`)
- Membaca dan menulis file dengan `os` dan `encoding/json`
- Pengelolaan waktu dengan `time`
- Memisahkan tanggung jawab antar file dalam satu package

## Lisensi

Proyek ini dibuat untuk tujuan belajar. Bebas dipelajari dan dimodifikasi.