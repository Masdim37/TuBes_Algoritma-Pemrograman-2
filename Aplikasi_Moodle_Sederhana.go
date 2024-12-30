package main

import "fmt"

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
	namaPengajar, namaMatkul           string
	namaTugas1, namaTugas2, namaTugas3 string
	namaQuiz1, namaQuiz2, namaQuiz3    string
}

type Tugas struct { //hanya berisi rincian tugas
	namaTugas                                                                                                                                                                                               string
	pertanyaan1, pertanyaan2, pertanyaan3, pertanyaan4, pertanyaan5, pertanyaan6, pertanyaan7, pertanyaan8, pertanyaan9, pertanyaan10                                                                       string
	jawabanPertanyaan1, jawabanPertanyaan2, jawabanPertanyaan3, jawabanPertanyaan4, jawabanPertanyaan5, jawabanPertanyaan6, jawabanPertanyaan7, jawabanPertanyaan8, jawabanPertanyaan9, jawabanPertanyaan10 string
	nilaiPertanyaan1, nilaiPertanyaan2, nilaiPertanyaan3, nilaiPertanyaan4, nilaiPertanyaan5, nilaiPertanyaan6, nilaiPertanyaan7, nilaiPertanyaan8, nilaiPertanyaan9, nilaiPertanyaan10 int
}

type Quiz struct { //hanya berisi rincian quiz
	namaQuiz                                                                                                                                                                                                string
	pertanyaan1, pertanyaan2, pertanyaan3, pertanyaan4, pertanyaan5, pertanyaan6, pertanyaan7, pertanyaan8, pertanyaan9, pertanyaan10                                                                       string
	jawabanPertanyaan1, jawabanPertanyaan2, jawabanPertanyaan3, jawabanPertanyaan4, jawabanPertanyaan5, jawabanPertanyaan6, jawabanPertanyaan7, jawabanPertanyaan8, jawabanPertanyaan9, jawabanPertanyaan10 string
	nilaiPertanyaan1, nilaiPertanyaan2, nilaiPertanyaan3, nilaiPertanyaan4, nilaiPertanyaan5, nilaiPertanyaan6, nilaiPertanyaan7, nilaiPertanyaan8, nilaiPertanyaan9, nilaiPertanyaan10 int                                                                                                                                                                                                  int
}

type Nilai struct { //hanya berisi rincian nilai
	nilaiTugas1, nilaiTugas2, nilaiTugas3 int
	nilaiQuiz1, nilaiQuiz2, nilaiQuiz3    int
	nilaiRataRata                         float64
}

type ArrMurid [SIZE]Murid //membuat tipe data ArrMurid yang merupakan array struct

type ArrGuru [SIZE]Guru //membuat tipe data ArrGuru yang merupakan array struct

type ArrMatkul [SIZE]Matkul //membuat tipe data ArrMatkul yang merupakan array struct

type ArrTugas [SIZE]Tugas //membuat tipe data ArrTugas yang merupakan array struct

type ArrQuiz [SIZE]Quiz //membuat tipe data ArrQuiz yang merupakan array struct

type ArrNilai [SIZE]Nilai //membuat tipe data ArrNilai yang merupakan array struct

var murid ArrMurid //membuat array bernama murid dengan tipe data ArrMurid (array struct)

var guru ArrGuru //membuat array bernama guru dengan tipe data ArrGuru (array struct)

var matkul ArrMatkul //membuat array bernama matkul dengan tipe data ArrMatkul (array struct)

var tugasMatkulMatematika ArrTugas //membuat array bernama tugas dengan tipe data ArrTugas (array struct)
var tugasMatkulIPA ArrTugas //membuat array bernama tugas dengan tipe data ArrTugas (array struct)
var tugasMatkulBIndo ArrTugas //membuat array bernama tugas dengan tipe data ArrTugas (array struct)

var quizMatkulMatematika ArrQuiz //membuat array bernama quiz dengan tipe data ArrQuiz (array struct)
var quizMatkulIPA ArrQuiz //membuat array bernama quiz dengan tipe data ArrQuiz (array struct)
var quizMatkulBIndo ArrQuiz //membuat array bernama quiz dengan tipe data ArrQuiz (array struct)

var nilai ArrNilai //membuat array bernama nilai dengan tipe data ArrNilai (array struct)

func main() {
	var dataMurid *Murid
	var dataGuru *Guru
	var dataMatkul *Matkul
	var dataTugas *Tugas
	var dataQuiz *Quiz
	var dataNilai *Nilai
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
				dataGuru = &guru[i]     //Referensikan dataGuru ke array guru, selanjutnya tinggal gunakan dataGuru.[komponen]
				dataMatkul = &matkul[i] //Referensikan dataMatkul ke array matkul, selanjutnya tinggal gunakan dataMatkul.[komponen]
				dataTugas = &tugas[i]   //Referensikan dataTugas ke array tugas, selanjutnya tinggal gunakan dataTugas.[komponen]
				dataQuiz = &quiz[i]     //Referensikan dataQuiz ke array quiz, selanjutnya tinggal gunakan dataQuiz.[komponen]
				dataNilai = &nilai[i]   //Referensikan dataNilai ke array nilai, selanjutnya tinggal gunakan dataNilai.[komponen]
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
		if 	
		fmt.Print("Masukkan Nama Quiz : ")
			fmt.Scan(&dataQuiz.namaQuiz)
		case 3: //tampilan jika memilih menu 3. Edit Tugas

		case 4: //tampilan jika memilih menu 4. Edit Quiz

		case 5: //tampilan jika memilih menu 5. Hapus Tugas

		case 6: //tampilan jika memilih menu 6. Hapus Quiz

		case 7: //tampilan jika memilih menu 7. Tampilkan List Tugas

		case 8: //tampilan jika memilih menu 8. Tampilkan List Quiz

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
