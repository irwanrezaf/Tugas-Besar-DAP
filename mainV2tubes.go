/*
==== bug uy bug ====
1. nim masih bisa diisi karakter √
2. nama belum bisa di spasi √
3. nim yang sama masih di anggap valid √
4. beberapa inputan tipe data int masih error jika di isi karakter √
5. belum ada IPK dan GRADE √
6. Sorting belum dibuat (tapi belum lengkap) √
7. IPK malah 0 √

===== spesifikasi =====
1. hapus matkul
2. mengurutkan sesuai ipk √
3. menampilkan sesuai nilai dan sks (di pilihan 5) √
4. transkrip nilai


												==== Program ====
1. Untuk input data mahasiswa seperti nama, nim, matkul yang diambil, dan nilai[quiz, uts, uas])
2. untuk mengedit data mahasiswa ada dua pilihan, pertama tama kita cari nama mahasiswa yang akan di edit datanya
    [1] untuk mengganti matkul yang sudah di inputkan
	   terdapat daftar matakuliah yang diambil oleh mahasiswa yang telah di cari, lalu pilih nomor mata kuliah yang akan diganti
    [2] untuk menambah matkul yang di ambil
	   disini akan ditanyakan jumlah mata kuliah yang akan ditambahkan, lalu inputkan seperti no 1
3. Sorting data masih kosong
4. Untuk mencari data berdasatkan :
   [1] mencari nama mahasiswa dan menampilkan mata kuliah yang di ambilnya
   [2] mencari mata kuliah dan menampilkan nama mahasiswa yang mengambil mata kuliah tersebut
5. Untuk melihat semua data mahasiswa yang diinputkan
6. untuk menghapus data mahasiswa dengan cara menggeser array data setelahnya ke data yang dihapus, pergeseran dilakukan sebanyak
   jumlah mahasiswa dikurangi indexk mahasiswa yang di hapus
0. keluar dari program (inputan akan tereset)
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	N       = 1000
	Warning = "\033[1;31m%s\033[0m"
)

type (
	Mahasiswa struct {
		nama   string
		nim    string
		matkul [100]Matakuliah
		nummk  int
		ipk    float64
		rerata float64
	}

	data [N]Mahasiswa

	Matakuliah struct {
		kode   string
		nilai  [3]float64
		ratamk float64
		sks    int
		grade  float64
		index  string
	}
)

var (
	mhs           data
	nummhs        int
	consoleReader = bufio.NewReader(os.Stdin)
)

func main() {
	fmt.Println("Nama : IRWAN REZA FIRMANSYAH")
	var (
		pil, key, pause string
		pilInt          int
	)

	// inisiasi
	nummhs = 0

	// main
	fmt.Println("Selamat Datang Di Aplikasi Nilai Mahasiswa")
	menu()
	fmt.Print("Pilihan : ")
	fmt.Scanln(&pil)
	for pil != "0" {
		switch pil {
		case "1":
			// input data mhs
			back("Masukan data mahasiswa [Y/N] : ", &key)
			for (key == "y" || key == "Y") && (key != "n" || key != "N") {
				loadMhs(nummhs)

				back("Masukan mata kuliah [Y/N] : ", &key)
				for i := 0; key != "N" && key != "n"; i++ {
					loadMatkul(nummhs, mhs[nummhs].nummk)
					mhs[nummhs].nummk++
					back("Tambah mata kuliah [Y/N] : ", &key)
				}
				nummhs++
				back("Tambah data mahasiswa lagi?[Y/N] : ", &key)
			}
		case "2":
			// edit data mhs
			fmt.Println("1. Ganti Data mahasiswa")
			fmt.Println("2. Ganti mata kuliah")
			fmt.Println("3. Tambah mata kuliah")
			fmt.Println("0. Kembali")
			fmt.Print("Pilihan : ")
			fmt.Scan(&pil)
			if pil == "1" {
				fmt.Print("NIM mahasiswa : ")
				fmt.Scan(&key)

				idxmhs := cariNIM(key)
				mhs[idxmhs].nama = ""
				mhs[idxmhs].nim = ""
				loadMhs(idxmhs)
			} else if pil == "2" {
				fmt.Print("NIM mahasiswa : ")
				fmt.Scanln(&key)
				idxmhs := cariNIM(key)
				if idxmhs != -1 {
					fmt.Println("Nama	:", mhs[idxmhs].nama)
					fmt.Println("Matkul yang akan diganti : ")
					for i := 0; i < mhs[idxmhs].nummk; i++ {
						fmt.Printf("%d. %s\n", i+1, mhs[idxmhs].matkul[i].kode)
					}
					fmt.Print("Pilihan : ")
					convToInt("Pilihan : ", 1, mhs[idxmhs].nummk, &pilInt)
					if pilInt > mhs[idxmhs].nummk || pilInt < mhs[idxmhs].nummk {
						fmt.Printf("Tidak ada mata kuliah ke %d \n", pilInt)
					} else {
						loadMatkul(idxmhs, pilInt-1)
					}
				} else {
					fmt.Printf("%s tidak ditemukan\n", key)
				}
			} else if pil == "3" {
				fmt.Print("NIM mahasiswa : ")
				fmt.Scanln(&key)
				idxmhs := cariNIM(key)
				if idxmhs != -1 {

					back("Tambah mata kuliah?[Y/N] : ", &pil)
					for pil == "y" || pil == "Y" && pil != "n" || pil != "N" {
						loadMatkul(idxmhs, mhs[idxmhs].nummk)
						mhs[idxmhs].nummk++
						back("Tambah mata kuliah lagi?[Y/N] : ", &pil)
					}
				} else {
					fmt.Printf("NIM : %s tidak ditemukan\n", key)
				}
			} else if pil == "0" {

			} else {
				fmt.Println("Pilihan tidak valid")
			}
		case "3":
			// cari data mhs
			fmt.Println("1. Cari berdasarkan NIM mahasiswa")
			fmt.Println("2. Cari berdasarkan mata kuliah")
			fmt.Println("0. Kembali")
			fmt.Print("Pilihan : ")
			fmt.Scanln(&pil)

			if pil == "1" {
				fmt.Print("NIM mahasiswa : ")
				fmt.Scanln(&key)
				idxmhs := cariNIM(key)
				if idxmhs != -1 {
					viewData(mhs, idxmhs)
				} else {
					fmt.Printf("NIM : %s tidak ditemukan\n", key)
				}
			} else if pil == "2" {
				fmt.Print("Mata kuliah : ")
				fmt.Scanln(&key)
				viewMatkul(key)
			} else if pil == "0" {

			} else {
				fmt.Println("Pilihan tidak valid")
			}
		case "4":
			// show data mhs
			if mhs[0].nim != "" {
				viewMhs(mhs)
			} else {
				fmt.Println("Belum ada data")
			}

		case "5":
			// hapus data mhs
			fmt.Println("1. Hapus Mahasiswa")
			fmt.Println("2. Hapus mata kuliah")
			fmt.Println("0. Kembali")
			fmt.Print("Pilihan : ")
			fmt.Scanln(&pil)
			if pil == "1" {
				// hapus mahasiswa
				fmt.Print("NIM mahasiswa : ")
				fmt.Scanln(&key)
				idxmhs := cariNIM(key)
				if idxmhs != -1 {
					delete(idxmhs)
				} else {
					fmt.Printf("NIM : %s tidak ditemukan\n", key)
				}
			} else if pil == "2" {
				// hapus matkul
				fmt.Print("NIM mahasiswa : ")
				fmt.Scanln(&key)
				idxmhs := cariNIM(key)
				if idxmhs != -1 {
					delete(idxmhs)
				} else {
					fmt.Printf("NIM : %s tidak ditemukan\n", key)
				}
			} else if pil == "0" {

			} else {
				fmt.Println("pilihan tidak valid")
			}
		case "6":
			for i := 0; i < nummhs; i++ {
				viewTranskrip(i)
			}
		default:
			fmt.Println("pilihan tidak valid")
		}
		fmt.Print("Press enter to continue...")
		fmt.Scanln(&pause)
		fmt.Println()

		menu()
		fmt.Print("Pilihan : ")
		fmt.Scanln(&pil)
	}
}

// ==========================================================================================================================================
// ==========================================================================================================================================
// ============================================================= BATAS FUNCTION =============================================================
// ==========================================================================================================================================
// ==========================================================================================================================================

func menu() {
	fmt.Println("------------- Menu -------------")
	fmt.Println("1. Input Data Mahasiswa")
	fmt.Println("2. Edit Data Mahasiwa")
	fmt.Println("3. Cari Data Mahasiswa")
	fmt.Println("4. Lihat Semua Data Mahasiswa")
	fmt.Println("5. Hapus Data Mahasiswa")
	fmt.Println("6. Lihat Transkrip Nilai Mahasiswa")
	fmt.Println("0. Keluar")
}

func loadMhs(idxmhs int) {
	var nim string
	fmt.Scanln(&nim)
	fmt.Print("Nama	: ")
	mhs[idxmhs].nama, _ = consoleReader.ReadString('\n')
	mhs[idxmhs].nama = strings.TrimSpace(strings.ToUpper(mhs[idxmhs].nama))

	fmt.Print("NIM	: ")
	fmt.Scanln(&nim)
	_, err := strconv.Atoi(nim)
	for err != nil || cariNIM(nim) != -1 {
		fmt.Printf(Warning, "{NIM telah terdaftar/tidak valid}\n")
		fmt.Print("NIM	: ")
		fmt.Scanln(&nim)
		_, err = strconv.Atoi(nim)
	}
	mhs[idxmhs].nim = nim
}

func loadMatkul(idxmhs, idxmatkul int) {
	var key string

	fmt.Printf("Nama matkul ke %d  : ", idxmatkul+1)
	fmt.Scanln(&key)
	idxmk := cariMatkul(idxmhs, key)
	for idxmk != -1 {
		fmt.Println("{Mata kuliah telah terdaftar}")
		fmt.Printf("Nama matkul ke %d  : ", idxmatkul+1)
		fmt.Scanln(&key)
		idxmk = cariMatkul(idxmhs, key)
	}
	mhs[idxmhs].matkul[idxmatkul].kode = key

	fmt.Print("Jumlah SKS   : ")
	convToInt("Jumlah SKS   : ", 1, 5, &mhs[idxmhs].matkul[idxmatkul].sks)

	fmt.Print("Nilai Quiz   : ")
	convToFloat("Nilai Quiz   : ", 0, 100, &mhs[idxmhs].matkul[idxmatkul].nilai[0])

	fmt.Print("Nilai UTS    : ")
	convToFloat("Nilai UTS   : ", 0, 100, &mhs[idxmhs].matkul[idxmatkul].nilai[1])

	fmt.Print("Nilai UAS    : ")
	convToFloat("Nilai UAS   : ", 0, 100, &mhs[idxmhs].matkul[idxmatkul].nilai[2])

	cekGrade(idxmhs, idxmatkul)
	cekIpk(idxmhs)

}

// fungsi sederhananya
func hitungSKS(idxmhs int) int {
	sum := 0
	for i := 0; i < mhs[idxmhs].nummk; i++ {
		sum = sum + mhs[idxmhs].matkul[i].sks
	}
	return sum
}

func hitungRata(idxmhs, idxmk int) float64 {
	rata := float64(mhs[idxmhs].matkul[idxmk].nilai[0]+mhs[idxmhs].matkul[idxmk].nilai[1]+mhs[idxmhs].matkul[idxmk].nilai[2]) / 3.0
	return rata
}

func cekGrade(idxmhs, idxmk int) {
	rerata := hitungRata(idxmhs, idxmk)
	switch {
	case rerata >= 90:
		mhs[idxmhs].matkul[idxmk].index = "A"
		mhs[idxmhs].matkul[idxmk].grade = 4.0
	case rerata >= 80:
		mhs[idxmhs].matkul[idxmk].index = "AB"
		mhs[idxmhs].matkul[idxmk].grade = 3.5
	case rerata >= 70:
		mhs[idxmhs].matkul[idxmk].index = "B"
		mhs[idxmhs].matkul[idxmk].grade = 3.0
	case rerata >= 60:
		mhs[idxmhs].matkul[idxmk].index = "BC"
		mhs[idxmhs].matkul[idxmk].grade = 2.5
	case rerata >= 50:
		mhs[idxmhs].matkul[idxmk].index = "C"
		mhs[idxmhs].matkul[idxmk].grade = 2.0
	case rerata >= 40:
		mhs[idxmhs].matkul[idxmk].index = "CD"
		mhs[idxmhs].matkul[idxmk].grade = 1.5
	case rerata >= 30:
		mhs[idxmhs].matkul[idxmk].index = "D"
		mhs[idxmhs].matkul[idxmk].grade = 1.0
	case rerata >= 15:
		mhs[idxmhs].matkul[idxmk].index = "DE"
		mhs[idxmhs].matkul[idxmk].grade = 0.5
	case rerata >= 0:
		mhs[idxmhs].matkul[idxmk].index = "E"
		mhs[idxmhs].matkul[idxmk].grade = 0
	}
}

func cekIpk(idxmhs int) {

	sum := 0.0
	for i := 0; i < mhs[idxmhs].nummk; i++ {
		sum = sum + mhs[idxmhs].matkul[i].grade*float64(mhs[idxmhs].matkul[i].sks)
		fmt.Println(sum)
	}

	mhs[idxmhs].ipk = float64(sum) / float64(hitungSKS(idxmhs))
}

func cariNIM(key string) int {
	for i := 0; i < nummhs; i++ {
		if mhs[i].nim == key {
			return i
		}
	}
	return -1

}

func cariMatkul(idxmhs int, key string) int {
	var i int

	for i = 0; strings.ToLower(key) != strings.ToLower(mhs[idxmhs].matkul[i].kode) && i < mhs[idxmhs].nummk; i++ {
	}

	if strings.ToLower(key) == strings.ToLower(mhs[idxmhs].matkul[i].kode) {
		return i
	} else {
		return -1
	}
}

func delete(idxmhs int) {
	for idxmhs < nummhs {
		mhs[idxmhs] = mhs[idxmhs+1]
		idxmhs++
	}
	nummhs--
}

func deleteMk(idxmhs, idxmk int) {
	for idxmk < mhs[idxmhs].nummk {
		mhs[idxmhs].matkul[idxmk] = mhs[idxmhs].matkul[idxmk+1]
		idxmk++
	}
	mhs[idxmhs].nummk--
}

func viewData(mahas data, idxmhs int) {
	fmt.Println("Nama		: ", mahas[idxmhs].nama)
	fmt.Println("Nim		:", mahas[idxmhs].nim)

	fmt.Print("Mata kuliah	: ")
	fmt.Printf("%s", mahas[idxmhs].matkul[0].kode)
	for j := 1; j < mahas[idxmhs].nummk; j++ {
		fmt.Printf(", %s", mahas[idxmhs].matkul[j].kode)
	}
	fmt.Println()
	fmt.Println("Jumlah SKS	:", hitungSKS(idxmhs))
}

func viewMatkul(key string) {
	for i := 0; i < nummhs; i++ {
		for j := 0; j < mhs[i].nummk; j++ {
			if strings.ToLower(mhs[i].matkul[j].kode) == strings.ToLower(key) {
				fmt.Printf("%d. %s", i+1, mhs[i].nama)
			}
		}
	}
}

func viewMhs(mahas data) {

	for i := 1; i < nummhs; i++ {
		for j := i; j > 0 && mahas[j-1].ipk < mahas[j].ipk; j-- {
			temp := mahas[j-1]
			mahas[j-1] = mahas[j]
			mahas[j] = temp
		}
	}

	for i := 0; i < nummhs; i++ {
		fmt.Printf("========= IPK TERTINGGI KE %d =========\n", i+1)
		viewData(mahas, i)
		fmt.Printf("IPK		: %.2f\n", mahas[i].ipk)
	}
}

func viewTranskrip(idxmhs int) {
	fmt.Println("Nama 		:", mhs[idxmhs].nama)
	fmt.Println("NIM 		:", mhs[idxmhs].nim)
	for i := 0; i < mhs[idxmhs].nummk; i++ {
		fmt.Println()
		fmt.Printf("=============== %s ===============\n", mhs[idxmhs].matkul[i].kode)
		fmt.Println("Sks		:", mhs[idxmhs].matkul[i].sks)
		fmt.Println("Nilai		:", mhs[idxmhs].matkul[i].index)
		fmt.Println("Mutu		:", mhs[idxmhs].matkul[i].grade)
		fmt.Println("SKS x Mutu	:", mhs[idxmhs].matkul[i].grade*float64(mhs[idxmhs].matkul[i].sks))
	}
	fmt.Println()
	fmt.Println("Jumlah SKS 		:", hitungSKS(idxmhs))
	fmt.Println("Jumlah SKS x Mutu	:", hitungSKS(idxmhs)*1)
	fmt.Println("Indeks Prestasi		:", mhs[idxmhs].ipk)
	fmt.Println("========================================")
}

// ==================== error handling =====================

func convToInt(print string, bawah, atas int, outInt *int) {
	var str string

	fmt.Scanln(&str)
	value, err := strconv.Atoi(str)
	for err != nil || value < bawah || value > atas {
		fmt.Printf(Warning, "Error : inputan tidak valid\n")
		fmt.Print(print)
		fmt.Scanln(&str)
		value, err = strconv.Atoi(str)
	}
	*outInt = value
}

func convToFloat(print string, bawah, atas float64, outInt *float64) {
	var str string

	fmt.Scanln(&str)
	value, err := strconv.ParseFloat(str, 64)
	for err != nil || value < bawah || value > atas {
		fmt.Printf(Warning, "Error : inputan tidak valid\n")
		fmt.Print(print)
		fmt.Scanln(&str)
		value, err = strconv.ParseFloat(str, 64)
	}
	*outInt = value
}

func back(print string, outStr *string) {
	var str string

	fmt.Print(print)
	fmt.Scanln(&str)
	for str != "Y" && str != "y" && str != "N" && str != "n" {
		fmt.Printf(Warning, "Error : inputan tidak valid\n")
		fmt.Print(print)
		fmt.Scanln(&str)
	}
	*outStr = str
}
