package main

import (
	"fmt"
)

const SIZE int = 100

type Murid struct { //hanya berisi biodata murid
	nama, NIS, kelas   string
	username, password string
}

type Guru struct { //hanya berisi biodata guru
	nama, NIP          string
	username, password string
	namaMatkul         string
}

type Matkul struct { //hanya berisi rincian matkul
	namaPengajar, namaMatkul string
	namaTugas                [SIZE]string
	namaQuiz                 [SIZE]string
	pertanyaanTugas          [SIZE]string
	pertanyaanQuiz           [SIZE]string
	jawabanPertanyaanTugas   [SIZE]string
	jawabanPertanyaanQuiz    [SIZE]string
	bobotNilaiTugas          [SIZE]int
	bobotNilaiQuiz           [SIZE]int
	nilaiTugas               [SIZE]int
	nilaiQuiz                [SIZE]int
	nilaiMatkul              float64
}

type Nilai struct { //hanya berisi rincian nilai
	nilaiTugas1, nilaiTugas2, nilaiTugas3 int
	nilaiQuiz1, nilaiQuiz2, nilaiQuiz3    int
	nilaiRataRata                         float64
}

type ArrMurid [SIZE]Murid //membuat tipe data ArrMurid yang merupakan array struct

type ArrGuru [SIZE]Guru //membuat tipe data ArrGuru yang merupakan array struct

type ArrMatkul [SIZE]Matkul //membuat tipe data ArrMatkul yang merupakan array struct

type ArrNilai [SIZE]Nilai //membuat tipe data ArrNilai yang merupakan array struct

var murid ArrMurid //membuat array bernama murid dengan tipe data ArrMurid (array struct)

var guru ArrGuru //membuat array bernama guru dengan tipe data ArrGuru (array struct)

var matkulMatematika ArrMatkul //membuat array bernama matkulMatematika dengan tipe data ArrMatkul (array struct)
var matkulIPA ArrMatkul        //membuat array bernama matkulIPA dengan tipe data ArrMatkul (array struct)
var matkulBIndo ArrMatkul      //membuat array bernama matkulBIndo dengan tipe data ArrMatkul (array struct)

var nilai ArrNilai //membuat array bernama nilai dengan tipe data ArrNilai (array struct)

func main() {
	var dataMurid *Murid
	var dataGuru *Guru
	var dataMatkul *Matkul
	var pilihanLogin int
	var username, password string
	var pilihanMenu int

	//isi data array guru
	guru[0] = Guru{"Dhimas Hafizh", "NIP123", "GuruDhimas", "dhimas", "Matematika"}
	guru[1] = Guru{"Satrio Wibowo", "NIP456", "GuruSatrio", "satrio", "IPA"}
	guru[2] = Guru{"Azaria Nanda", "NIP789", "GuruNanda", "nanda", "Bahasa Indonesia"}

	//isi data array murid
	murid[0] = Murid{"Dhimas Hafizh", "NIS123", "XII-IPA-1", "dhimas", "dhimas"}
	murid[1] = Murid{"Satrio Wibowo", "NIS456", "XII-IPA-1", "satrio", "satrio"}
	murid[2] = Murid{"Azaria Nanda", "NIS789", "XII-IPA-1", "nanda", "nanda"}

MenuLogin:
	fmt.Println("--- LMS SEDERHANA ---")
	fmt.Println("1. Login Guru")
	fmt.Println("2. Login Murid")
	fmt.Print("Masukkan pilihan anda : ")
	fmt.Scan(&pilihanLogin)

	switch pilihanLogin {
	case 1:
		fmt.Print("Masukkan username : ")
		fmt.Scan(&username)
		fmt.Print("Masukkan password : ")
		fmt.Scan(&password)

		// Cari data guru
		for i := 0; i < SIZE; i++ {
			if guru[i].username == username && guru[i].password == password { //cek login
				dataGuru = &guru[i] //Referensikan dataGuru ke array guru, selanjutnya tinggal gunakan dataGuru.[komponen]
				break
			}
		}

		if dataGuru != nil {
			fmt.Println()
			fmt.Println("Selamat Datang", dataGuru.nama)
			fmt.Println("Mata kuliah yang anda ajar :", dataGuru.namaMatkul)
		} else {
			fmt.Println()
			fmt.Println("Username atau password salah")
			fmt.Println()
			goto MenuLogin
		}

	MenuGuru:
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Tambah Quiz")
		fmt.Println("3. Edit Tugas")
		fmt.Println("4. Edit Quiz")
		fmt.Println("5. Hapus Tugas")
		fmt.Println("6. Hapus Quiz")
		fmt.Println("7. Tampilkan list Tugas")
		fmt.Println("8. Tampilkan List Quiz")
		fmt.Println("9. Forum Mata Kuliah")
		fmt.Println("10. Lihat List Nilai Murid")
		fmt.Print("Pilihan anda : ")
		fmt.Scan(&pilihanMenu)

		switch pilihanMenu {
		case 1: //tampilan jika memilih menu 1. Tambah Tugas

		case 2: //tampilan jika memilih menu 2. Tambah Quiz
			// Tentukan dataMatkul berdasarkan nama guru
			if dataGuru.nama == "Dhimas Hafizh" {
				dataMatkul = &matkulMatematika[0]
			} else if dataGuru.nama == "Satrio Wibowo" {
				dataMatkul = &matkulIPA[0]
			} else if dataGuru.nama == "Azaria Nanda" {
				dataMatkul = &matkulBIndo[0]
			}

			for i := 0; i < SIZE; {
				if dataMatkul.namaQuiz[i] == "" { // Jika indeks kosong
					fmt.Print("Masukkan nama Quiz", i+1, " : ")
					fmt.Scan(&dataMatkul.namaQuiz[i])
					break
				} else {
					i++ // Pindah ke indeks berikutnya jika tidak kosong
				}
			}

			for i := 0; i < SIZE; i++ {
				if dataMatkul.namaQuiz[i] != "" { // Jika indeks tidak kosong
					var tempPertanyaan string
					fmt.Print("Masukkan pertanyaan ke-", i+1, " (masukkan stop untuk berhenti) : ")
					fmt.Scan(&tempPertanyaan)
					if tempPertanyaan == "stop" {
						break
					}
					dataMatkul.pertanyaanQuiz[i] = tempPertanyaan
					fmt.Print("Masukkan kunci jawaban pertanyaan ke-", i+1, " : ")
					fmt.Scan(&dataMatkul.jawabanPertanyaanQuiz[i])
					fmt.Print("Masukkan bobot nilai pertanyaan ke-", i+1, " : ")
					fmt.Scan(&dataMatkul.bobotNilaiQuiz[i])

				}
			}

		case 3: //tampilan jika memilih menu 3. Edit Tugas

		case 4: //tampilan jika memilih menu 4. Edit Quiz

		case 5: //tampilan jika memilih menu 5. Hapus Tugas

		case 6: //tampilan jika memilih menu 6. Hapus Quiz

		case 7: //tampilan jika memilih menu 7. Tampilkan List Tugas

		case 8: //tampilan jika memilih menu 8. Tampilkan List Quiz
			fmt.Println("Data Quiz yang tersimpan")
			for i := 0; i < SIZE; i++ {
				fmt.Println(dataMatkul.namaQuiz[i])
			}

		case 9: //tampilan jika memilih menu 9. Forum Mata Kuliah

		case 10: //tampilan jika memilih menu 10. Lihat List Nilai Murid

		default: //tampilan jika memilih selain menu 1-10
			fmt.Println("Pilihan yang dimasukkan tidak valid")
			fmt.Println()
			goto MenuGuru
		}

	case 2:
		fmt.Print("Masukkan username : ")
		fmt.Scan(&username)
		fmt.Print("Masukkan password : ")
		fmt.Scan(&password)
		// Panggil loginGuru dan tampilkan nama guru
		// Cari data murid
		for i := 0; i < SIZE; i++ {
			if murid[i].username == username && murid[i].password == password {
				dataMurid = &murid[i] // Referensikan ke elemen array
				break
			}
		}

		if dataMurid != nil {
			fmt.Println()
			fmt.Println("Selamat Datang", dataMurid.nama)
		} else {
			fmt.Println()
			fmt.Println("Username atau password salah")
			fmt.Println()
			goto MenuLogin
		}

	default:
		fmt.Println("Pilihan yang dimasukkan tidak valid")
		fmt.Println()
		goto MenuLogin
	}
}
