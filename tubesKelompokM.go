package main

import "fmt"

const NMAX int = 100
const MAX_RIWAYAT int = 50

type Kontak struct {
	Telepon string
	Email   string
}

type Riwayat struct {
	Tanggal   string
	Deskripsi string
}

type Supplier struct {
	NamaPerusahaan string
	Lokasi         string
	JenisMaterial  string
	RatingPerforma float64
	DetailKontak   Kontak
	RiwayatLayanan [MAX_RIWAYAT]Riwayat
	JumRiwayat     int
}

type TabSupplier [NMAX]Supplier

func tambahSupplier(A *TabSupplier, n *int) {
	if *n < NMAX {
		fmt.Print("Nama Perusahaan (Tanpa Spasi) : ")
		fmt.Scan(&A[*n].NamaPerusahaan)
		fmt.Print("Lokasi/Kota (Tanpa Spasi)   : ")
		fmt.Scan(&A[*n].Lokasi)
		fmt.Print("Jenis Material (Tanpa Spasi): ")
		fmt.Scan(&A[*n].JenisMaterial)
		fmt.Print("Rating Performa (0.0 - 5.0) : ")
		fmt.Scan(&A[*n].RatingPerforma)
		fmt.Print("No Telepon                  : ")
		fmt.Scan(&A[*n].DetailKontak.Telepon)
		fmt.Print("Email                       : ")
		fmt.Scan(&A[*n].DetailKontak.Email)
		A[*n].JumRiwayat = 0
		*n = *n + 1
		fmt.Println("Data berhasil ditambahkan!")
	} else {
		fmt.Println("Kapasitas penyimpanan penuh!")
	}
}

func tampilSemuaSupplier(A TabSupplier, n int) {
	if n == 0 {
		fmt.Println("Data supplier masih kosong.")
	} else {
		for i := 0; i < n; i++ {
			fmt.Printf("\n--- Supplier %d ---\n", i+1)
			fmt.Printf("Nama Perusahaan : %s\n", A[i].NamaPerusahaan)
			fmt.Printf("Lokasi          : %s\n", A[i].Lokasi)
			fmt.Printf("Jenis Material  : %s\n", A[i].JenisMaterial)
			fmt.Printf("Rating Performa : %.2f\n", A[i].RatingPerforma)
			fmt.Printf("Kontak          : %s | %s\n", A[i].DetailKontak.Telepon, A[i].DetailKontak.Email)
			fmt.Printf("Jumlah Riwayat  : %d\n", A[i].JumRiwayat)
			for j := 0; j < A[i].JumRiwayat; j++ {
				fmt.Printf("  - [%s] %s\n", A[i].RiwayatLayanan[j].Tanggal, A[i].RiwayatLayanan[j].Deskripsi)
			}
		}
	}
}

func editSupplier(A *TabSupplier, n int) {
	var target string
	var idx int = -1

	fmt.Print("Masukkan Nama Perusahaan yang ingin diedit: ")
	fmt.Scan(&target)

	for i := 0; i < n; i++ {
		if A[i].NamaPerusahaan == target {
			idx = i
		}
	}

	if idx != -1 {
		var pilihanEdit int
		fmt.Println("\n--- MENU EDIT DATA ---")
		fmt.Println("1. Ubah Nama Perusahaan")
		fmt.Println("2. Ubah Lokasi")
		fmt.Println("3. Ubah Jenis Material")
		fmt.Println("4. Ubah Rating Performa")
		fmt.Println("5. Tambah Riwayat Pelayanan")
		fmt.Print("Pilih data yang ingin diedit (1-5): ")
		fmt.Scan(&pilihanEdit)

		switch pilihanEdit {
		case 1:
			fmt.Print("Masukkan Nama Perusahaan Baru (Tanpa Spasi): ")
			fmt.Scan(&A[idx].NamaPerusahaan)
			fmt.Println("Nama Perusahaan berhasil diubah!")
		case 2:
			fmt.Print("Masukkan Lokasi Baru (Tanpa Spasi): ")
			fmt.Scan(&A[idx].Lokasi)
			fmt.Println("Lokasi berhasil diubah!")
		case 3:
			fmt.Print("Masukkan Jenis Material Baru (Tanpa Spasi): ")
			fmt.Scan(&A[idx].JenisMaterial)
			fmt.Println("Jenis Material berhasil diubah!")
		case 4:
			fmt.Print("Masukkan Rating Baru: ")
			fmt.Scan(&A[idx].RatingPerforma)
			fmt.Println("Rating berhasil diubah!")
		case 5:
			if A[idx].JumRiwayat < MAX_RIWAYAT {
				riwayatKe := A[idx].JumRiwayat
				fmt.Print("Masukkan Tanggal (DD-MM-YYYY): ")
				fmt.Scan(&A[idx].RiwayatLayanan[riwayatKe].Tanggal)
				fmt.Print("Masukkan Deskripsi Pelayanan (Tanpa Spasi): ")
				fmt.Scan(&A[idx].RiwayatLayanan[riwayatKe].Deskripsi)
				A[idx].JumRiwayat++
				fmt.Println("Riwayat pelayanan berhasil ditambahkan!")
			} else {
				fmt.Println("Data riwayat pelayanan penuh!")
			}
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	} else {
		fmt.Println("Nama Perusahaan tidak ditemukan!")
	}
}

func hapusSupplier(A *TabSupplier, n *int) {
	var target string
	var idx int = -1

	fmt.Print("Masukkan Nama Perusahaan yang ingin dihapus: ")
	fmt.Scan(&target)

	for i := 0; i < *n; i++ {
		if A[i].NamaPerusahaan == target {
			idx = i
		}
	}

	if idx != -1 {
		for i := idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n = *n - 1
		fmt.Println("Data berhasil dihapus!")
	} else {
		fmt.Println("Nama Perusahaan tidak ditemukan!")
	}
}

func cariBerdasarkanLokasi(A TabSupplier, n int, targetLokasi string) {
	var i int
	var ditemukan bool = false

	for i = 0; i < n; i++ {
		if A[i].Lokasi == targetLokasi {
			fmt.Println("\nData ditemukan:")
			fmt.Println("Nama Perusahaan :", A[i].NamaPerusahaan)
			fmt.Println("Lokasi          :", A[i].Lokasi)
			fmt.Println("Rating          :", A[i].RatingPerforma)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Printf("\nSupplier dengan lokasi '%s' tidak ditemukan\n", targetLokasi)
	}
}

func cariBerdasarkanNama(A *TabSupplier, n int, targetNama string) int {
	urutkanBerdasarkanNama(A, n)

	var left int = 0
	var right int = n - 1
	var mid int

	for left <= right {
		mid = (left + right) / 2

		if A[mid].NamaPerusahaan == targetNama {
			return mid
		} else if A[mid].NamaPerusahaan < targetNama {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func urutkanRatingSelection(A *TabSupplier, n int, ascending bool) {
	var i, idxTarget, j int
	var temp Supplier

	for i = 0; i < n-1; i++ {
		idxTarget = i
		for j = i + 1; j < n; j++ {
			if ascending {
				if A[j].RatingPerforma < A[idxTarget].RatingPerforma {
					idxTarget = j
				}
			} else {
				if A[j].RatingPerforma > A[idxTarget].RatingPerforma {
					idxTarget = j
				}
			}
		}
		temp = A[i]
		A[i] = A[idxTarget]
		A[idxTarget] = temp
	}
}

func urutkanRatingInsertion(A *TabSupplier, n int, ascending bool) {
	var i, j int
	var temp Supplier

	for i = 1; i < n; i++ {
		temp = A[i]
		j = i - 1
		
		if ascending {
			for j >= 0 && A[j].RatingPerforma > temp.RatingPerforma {
				A[j+1] = A[j]
				j--
			}
		} else {
			for j >= 0 && A[j].RatingPerforma < temp.RatingPerforma {
				A[j+1] = A[j]
				j--
			}
		}
		A[j+1] = temp
	}
}

func urutkanBerdasarkanNama(A *TabSupplier, n int) {
	var i, j int
	var temp Supplier

	for i = 1; i < n; i++ {
		temp = A[i]
		j = i - 1
		for j >= 0 && A[j].NamaPerusahaan > temp.NamaPerusahaan {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}
}

func hitungSupplierPerWilayah(A TabSupplier, n int, targetWilayah string) int {
	var i, jumlah int
	jumlah = 0
	for i = 0; i < n; i++ {
		if A[i].Lokasi == targetWilayah {
			jumlah++
		}
	}
	return jumlah
}

func hitungRataRataKepuasan(A TabSupplier, n int) float64 {
	var i int
	var total float64
	total = 0
	for i = 0; i < n; i++ {
		total += A[i].RatingPerforma
	}
	if n == 0 {
		return 0
	}
	return total / float64(n)
}

func main() {
	var dataMitra TabSupplier
	var nData int = 0
	var pilihan int
	var jalan bool = true

	for jalan {
		fmt.Println("\n======================================================")
		fmt.Println("         SISTEM INFORMASI BANGUNIN (LOGISTIK)         ")
		fmt.Println("======================================================")
		fmt.Println("[ FITUR PENGELOLAAN DATA ]")
		fmt.Println("1. Tambah Data Supplier Baru")
		fmt.Println("2. Tampilkan Seluruh Data Supplier")
		fmt.Println("3. Ubah (Edit) Data Supplier")
		fmt.Println("4. Hapus Data Supplier")
		fmt.Println("\n[ FITUR PENCARIAN DATA ]")
		fmt.Println("5. Cari Supplier berdasarkan Nama (Binary Search)")
		fmt.Println("6. Cari Supplier berdasarkan Lokasi (Sequential Search)")
		fmt.Println("\n[ FITUR PENGURUTAN DATA ]")
		fmt.Println("7. Urutkan Supplier berdasar Rating (Selection Sort)")
		fmt.Println("8. Urutkan Supplier berdasar Rating (Insertion Sort)")
		fmt.Println("\n[ FITUR LAPORAN ]")
		fmt.Println("9. Tampilkan Statistik")
		fmt.Println("0. Keluar Aplikasi")
		fmt.Println("======================================================")
		fmt.Print("Pilih menu (0-9): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahSupplier(&dataMitra, &nData)
		case 2:
			tampilSemuaSupplier(dataMitra, nData)
		case 3:
			editSupplier(&dataMitra, nData)
		case 4:
			hapusSupplier(&dataMitra, &nData)
		case 5:
			var targetNama string
			fmt.Print("Masukkan Nama Perusahaan yang dicari: ")
			fmt.Scan(&targetNama)

			idx := cariBerdasarkanNama(&dataMitra, nData, targetNama)

			if idx != -1 {
				fmt.Println("\n--- Data Ditemukan ---")
				fmt.Printf("Nama Perusahaan : %s\n", dataMitra[idx].NamaPerusahaan)
				fmt.Printf("Lokasi          : %s\n", dataMitra[idx].Lokasi)
				fmt.Printf("Jenis Material  : %s\n", dataMitra[idx].JenisMaterial)
				fmt.Printf("Rating Performa : %.2f\n", dataMitra[idx].RatingPerforma)
				fmt.Printf("Kontak          : %s | %s\n", dataMitra[idx].DetailKontak.Telepon, dataMitra[idx].DetailKontak.Email)
			} else {
				fmt.Printf("\nSupplier dengan nama '%s' tidak ditemukan!\n", targetNama)
			}

		case 6:
			var targetLokasi string
			fmt.Print("Masukkan Lokasi/Kota yang dicari: ")
			fmt.Scan(&targetLokasi)
			cariBerdasarkanLokasi(dataMitra, nData, targetLokasi)

		case 7:
			var urutan int
			fmt.Println("\n>> SELECTION SORT (Berdasarkan Rating)")
			fmt.Println("1. Ascending (Terendah ke Tertinggi)")
			fmt.Println("2. Descending (Tertinggi ke Terendah)")
			fmt.Print("Pilih urutan (1/2): ")
			fmt.Scan(&urutan)

			if urutan == 1 {
				urutkanRatingSelection(&dataMitra, nData, true)
				fmt.Println("\n[ Berhasil diurutkan secara Ascending ]")
			} else if urutan == 2 {
				urutkanRatingSelection(&dataMitra, nData, false)
				fmt.Println("\n[ Berhasil diurutkan secara Descending ]")
			} else {
				fmt.Println("\nPilihan urutan tidak valid!")
			}
			tampilSemuaSupplier(dataMitra, nData)

		case 8:
			var urutan int
			fmt.Println("\n>> INSERTION SORT (Berdasarkan Rating)")
			fmt.Println("1. Ascending (Terendah ke Tertinggi)")
			fmt.Println("2. Descending (Tertinggi ke Terendah)")
			fmt.Print("Pilih urutan (1/2): ")
			fmt.Scan(&urutan)

			if urutan == 1 {
				urutkanRatingInsertion(&dataMitra, nData, true)
				fmt.Println("\n[ Berhasil diurutkan secara Ascending ]")
			} else if urutan == 2 {
				urutkanRatingInsertion(&dataMitra, nData, false)
				fmt.Println("\n[ Berhasil diurutkan secara Descending ]")
			} else {
				fmt.Println("\nPilihan urutan tidak valid!")
			}
			tampilSemuaSupplier(dataMitra, nData)

		case 9:
			if nData == 0 {
				fmt.Println("\nData supplier masih kosong, statistik tidak dapat dihitung.")
			} else {
				var targetWilayah string
				fmt.Println("\n--- LAPORAN STATISTIK ---")

				rataRata := hitungRataRataKepuasan(dataMitra, nData)
				fmt.Printf("1. Rata-rata Rating Kepuasan Keseluruhan : %.2f / 5.0\n", rataRata)

				fmt.Print("\nMasukkan Nama Wilayah/Kota untuk dicek : ")
				fmt.Scan(&targetWilayah)
				jumlahWilayah := hitungSupplierPerWilayah(dataMitra, nData, targetWilayah)
				fmt.Printf("2. Terdapat %d mitra supplier di wilayah %s.\n", jumlahWilayah, targetWilayah)
			}
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi BangunIn!")
			jalan = false
		default:
			fmt.Println("Pilihan tidak valid! Masukkan angka 0-9.")
		}
	}
}
