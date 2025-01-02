package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	namaTugas              [SIZE]string
	namaQuiz               [SIZE]string
	pertanyaanTugas        [SIZE]string
	pertanyaanQuiz         [SIZE]string
	jawabanPertanyaanTugas [SIZE]string
	jawabanPertanyaanQuiz  [SIZE]string
	bobotNilaiTugas        [SIZE]int
	bobotNilaiQuiz         [SIZE]int
	forumMatkul            [SIZE]string
}

type jawabanMatkul struct { //hanya berisi jawaban murid
	namaMuridTugas         [SIZE]string
	namaMuridQuiz          [SIZE]string
	jawabanPertanyaanTugas [SIZE]string
	jawabanPertanyaanQuiz  [SIZE]string
	nilaiPertanyaanTugas   [SIZE]int
	nilaiPertanyaanQuiz    [SIZE]int
	nilaiTugas             [SIZE]float64
	nilaiQuiz              [SIZE]float64
	nilaiMatkul            float64
}

type forum struct {
	pengirim   [SIZE]string
	waktuKirim [SIZE]string
	pesan      [SIZE]string
}

type ArrMurid [SIZE]Murid //membuat tipe data ArrMurid yang merupakan array struct
type ArrGuru [SIZE]Guru   //membuat tipe data ArrGuru yang merupakan array struct

var murid ArrMurid //membuat array bernama murid dengan tipe data ArrMurid (array struct)

var guru ArrGuru //membuat array bernama guru dengan tipe data ArrGuru (array struct)

var matkulMatematika Matkul //membuat variabel bernama matkulMatematika dengan tipe data Matkul (struct berisi array)
var matkulIPA Matkul        //membuat variabel bernama matkulIPA dengan tipe data Matkul (struct berisi array)
var matkulBIndo Matkul      //membuat variabel bernama matkulBIndo dengan tipe data Matkul (struct berisi array)

var jawabanMatkulMatematika jawabanMatkul //membuat variabel bernama jawabanMatkulMatematika dengan tipe data jawabanMatkul (struct berisi array)
var jawabanMatkulIPA jawabanMatkul        //membuat variabel bernama jawabanMatkulIPA dengan tipe data jawabanMatkul (struct berisi array)
var jawabanMatkulBIndo jawabanMatkul      //membuat variabel bernama jawabanMatkulBIndo dengan tipe data jawabanMatkul (struct berisi array)

var forumMatematika forum //membuat variabel bernama forumMatematika dengan tipe data forum (struct berisi array)
var forumIPA forum        //membuat variabel bernama forumIPA dengan tipe data forum (struct berisi array)
var forumBIndo forum      //membuat variabel bernama forumBIndo dengan tipe data forum (struct berisi array)

// function untuk cek apakah array kosong atau tidak
func isEmpty(arrData []string) bool {
	var cek bool = true
	for i := 0; i < len(arrData); i++ {
		if arrData[i] == "" {
			continue
		} else {
			cek = false
			break
		}
	}
	return cek
}

// function untuk menampilkan isi array
func tampilkan(arrData []string) {
	for i := 0; i < len(arrData); i++ {
		fmt.Printf("%d. %s\n", i+1, arrData[i])
		if arrData[i+1] == "" {
			break
		}
	}
}

// Function untuk membaca input dengan spasi
func bacaInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input) // Menghapus spasi dan newline
}

// function untuk input pertanyaan tugas dan quiz
func InputPertanyaanTugasQuiz(ArrNamaTugasQuiz []string, namaTugasQuiz string, ArrPertanyaanTugasQuiz []string, ArrJawabanPertanyaanTugasQuiz []string, ArrBobotNilaiPertanyaanTugasQuiz []int) {
	var tempPertanyaan string
	for i := 0; i < SIZE; i++ {
		if ArrNamaTugasQuiz[i] == namaTugasQuiz {
			for j := 0; j < 10; j++ {
				tempPertanyaan = bacaInput(fmt.Sprintf("Masukkan pertanyaan ke-%d : ", j+1))

				if tempPertanyaan == "stop" { // Berhenti jika input "stop"
					break
				}

				ArrPertanyaanTugasQuiz[i*10+j] = tempPertanyaan // Masukkan pertanyaan ke array pertanyaanTugas

				ArrJawabanPertanyaanTugasQuiz[i*10+j] = bacaInput(fmt.Sprintf("Masukkan kunci jawaban pertanyaan ke-%d : ", j+1)) // Input jawaban

				fmt.Printf("Masukkan bobot nilai pertanyaan ke-%d : ", j+1) // Input bobot nilai
				fmt.Scan(&ArrBobotNilaiPertanyaanTugasQuiz[i*10+j])         // Simpan bobot nilai
			}
		}
	}
}

// function untuk menghapus tugas quiz
func HapusTugasQuiz(ArrNamaTugasQuiz []string, ArrPertanyaanTugasQuiz []string, ArrJawabanPertanyaanTugasQuiz []string, ArrBobotNilaiPertanyaanTugasQuiz []int) {
	var pilihanHapusQuiz, konfirmasiHapusQuiz string

	pilihanHapusQuiz = bacaInput(fmt.Sprintf("Masukkan nama quiz yang ingin dihapus : "))
	fmt.Print("Apakah anda yakin ingin mengapus quiz ", pilihanHapusQuiz, " ? [y/n] : ")
	fmt.Scan(&konfirmasiHapusQuiz)
	if konfirmasiHapusQuiz == "y" || konfirmasiHapusQuiz == "Y" { //kalo masukin konfirmasi y, lakukan penghapusan
		for i := 0; i < SIZE; i++ {
			if ArrNamaTugasQuiz[i] == pilihanHapusQuiz { //kalo nama quiz yang diinputkan sama dengan nama quiz yang ada di array namaQuiz, maka lakukan penghapusan
				ArrNamaTugasQuiz[i] = "" //hapus nama quiz
				for j := 0; j < 10; j++ {
					ArrPertanyaanTugasQuiz[i*10+j] = ""          //hapus pertanyaan quiz (i = kode quiz, 10 = maks pertanyaan, j = kode pertanyaan quiz)
					ArrJawabanPertanyaanTugasQuiz[i*10+j] = ""   //hapus kunci jawaban quiz (i = kode quiz, 10 = maks pertanyaan, j = kode kunci jawaban pertanyaan quiz)
					ArrBobotNilaiPertanyaanTugasQuiz[i*10+j] = 0 // Menghapus bobot nilai pertanyaan quiz (i = kode quiz, 10 = maks pertanyaan, j = kode bobot nilai pertanyaan quiz)
				}
				fmt.Println("Quiz", pilihanHapusQuiz, "berhasil dihapus.") //tampilkan informasi bahwa quiz yang dipilih berhasil dihapus
				break                                                      //keluar dari loop
			}
		}
	} else if konfirmasiHapusQuiz == "n" || konfirmasiHapusQuiz == "N" { //kalo masukin konfirmasi n, batalkan penghapusan
		fmt.Print("Penghapusan quiz ", pilihanHapusQuiz, " dibatalkan") //tampilkan informasi bahwa quiz yang dipilih ngga jadi dihapus
	}
}

func EditTugasQuiz(ArrNamaTugasQuiz []string, ArrPertanyaanTugasQuiz []string, ArrJawabanPertanyaanTugasQuiz []string, ArrBobotNilaiPertanyaanTugasQuiz []int) {
	var pilihanEditQuiz string
	var pilihanNomorEditQuiz int

	pilihanEditQuiz = bacaInput(fmt.Sprintf("Masukkan nama quiz yang ingin diubah : "))
	fmt.Println("List Pertanyaan Quiz", pilihanEditQuiz, " :")
	for i := 0; i < SIZE; i++ {
		if ArrNamaTugasQuiz[i] == pilihanEditQuiz {
			for j := 0; j < 10; j++ {
				fmt.Println("Pertanyaan ke-", j+1)
				fmt.Println("Pertanyaan : ", ArrPertanyaanTugasQuiz[i*10+j])
				fmt.Println("Kunci Jawaban : ", ArrJawabanPertanyaanTugasQuiz[i*10+j])
				fmt.Println("Bobot Nilai : ", ArrBobotNilaiPertanyaanTugasQuiz[i*10+j])
			}
			fmt.Println()
			fmt.Print("Masukkan nomor pertanyaan yang ingin diubah : ")
			fmt.Scan(&pilihanNomorEditQuiz)
			ArrPertanyaanTugasQuiz[i*10+pilihanNomorEditQuiz-1] = bacaInput(fmt.Sprintf("Masukkan pertanyaan baru : "))
			ArrJawabanPertanyaanTugasQuiz[i*10+pilihanNomorEditQuiz-1] = bacaInput(fmt.Sprintf("Masukkan kunci jawaban baru : "))
			fmt.Print("Masukkan bobot nilai baru : ")
			fmt.Scan(&ArrBobotNilaiPertanyaanTugasQuiz[i*10+pilihanNomorEditQuiz-1])
			fmt.Println("Pertanyaan ke-", pilihanNomorEditQuiz, "pada kuis", pilihanEditQuiz, "berhasil diubah")
			break
		}
	}
}

func KerjakanTugasQuiz(ArrNamaTugasQuiz []string, ArrPertanyaanTugasQuiz []string, ArrInputJawabanPertanyaanTugasQuiz []string, ArrJawabanPertanyaanTugasQuiz []string, ArrInputNilaiPertanyaanTugasQuiz []int, ArrBobotNilaiPertanyaanTugasQuiz []int, ArrNamaMurid []string, namaMurid string, ArrNilaiTugasQuiz []float64) {
	var namaQuiz string
	var totalPoinSoalQuiz int
	var totalPoinJawabanQuiz int

	namaQuiz = bacaInput(fmt.Sprintf("Masukkan nama quiz yang ingin dikerjakan : "))
	fmt.Println()
	for i := 0; i < SIZE; i++ {
		if ArrNamaTugasQuiz[i] == namaQuiz { //kalo nama quiz yang diinputkan sama dengan nama quiz yang ada di array namaQuiz, maka lakukan pengerjaan quiz
			totalPoinSoalQuiz = 0    //ini nilai poin total soal quiz
			totalPoinJawabanQuiz = 0 //ini nilai poin total jawaban quiz siswa

			for j := 0; j < 10; j++ {
				fmt.Println(ArrPertanyaanTugasQuiz[i*10+j])                                              //tampilkan pertanyaan quiz di array pertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode pertanyaan quiz)
				ArrInputJawabanPertanyaanTugasQuiz[i*10+j] = bacaInput(fmt.Sprintf("Jawaban : "))        //input jawaban siswa, simpan di array jawabanPertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode jawaban pertanyaan quiz)
				if ArrInputJawabanPertanyaanTugasQuiz[i*10+j] == ArrJawabanPertanyaanTugasQuiz[i*10+j] { //kalo jawaban siswa sama dengan kunci jawaban yang ada di array jawabanPertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode kunci jawaban pertanyaan quiz)
					ArrInputNilaiPertanyaanTugasQuiz[i*10+j] += ArrBobotNilaiPertanyaanTugasQuiz[i*10+j] //maka nilai jawaban siswa di array nilaiPertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode nilai jawaban pertanyaan quiz) ditambahkan dengan bobot nilai quiz yang ada di array bobotNilaiQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode bobot nilai quiz)
				} else { //kalo jawabannya ga sama kaya kunci jawaban
					ArrInputNilaiPertanyaanTugasQuiz[i*10+j] += 0 //maka nilai jawaban siswa di array nilaiPertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode nilai jawaban pertanyaan quiz) ditambahkan dengan 0
				}
				totalPoinSoalQuiz += ArrBobotNilaiPertanyaanTugasQuiz[i*10+j]    //hitung total poin soal quiz
				totalPoinJawabanQuiz += ArrInputNilaiPertanyaanTugasQuiz[i*10+j] //hitung total poin jawaban quiz siswa
			}

			ArrNamaMurid[i] = namaMurid
			ArrNilaiTugasQuiz[i] = float64(totalPoinJawabanQuiz) / float64(totalPoinSoalQuiz) * 100 //perhitungan nilai akhir quiz siswa, disimpan di array nilaiQuiz indeks ke i (i = kode quiz)
			fmt.Println("Nilai Quiz Anda :", ArrNilaiTugasQuiz[i])                                  //tampilkan nilai quiz
		}
	}
}

func TampilkanNilaiRataRataMatkul(ArrNamaTugas []string, ArrNamaQuiz []string, ArrNilaiTugas []float64, ArrNilaiQuiz []float64, NilaiMatkul float64) {
	var jumlahTugas, RataRataTugas float64
	var JumlahQuiz, RataRataQuiz float64

	for i := 0; i < SIZE; i++ {
		if ArrNamaTugas[i] != "" {
			jumlahTugas++
		} else if ArrNamaQuiz[i] != "" {
			JumlahQuiz++
		}
	}

	for i := 0; i < SIZE; i++ {
		if ArrNilaiTugas[i] != 0 && ArrNilaiQuiz[i] != 0 {
			RataRataTugas += (ArrNilaiTugas[i] / jumlahTugas)
			RataRataQuiz += (ArrNilaiQuiz[i] / JumlahQuiz)
		}
	}

	NilaiMatkul = (RataRataTugas + RataRataQuiz) / 2
	fmt.Println("Nilai Rata-Rata Tugas :", RataRataTugas)
	fmt.Println("Nilai Rata-Rata Quiz :", RataRataQuiz)
	fmt.Println("Nilai Rata-Rata Mata Kuliah :", NilaiMatkul)
}

func PostForumMatkul(ArrPengirim []string, ArrWaktuKirim []string, ArrPesan []string) {
	var tempNamaPengirim, tempWaktuKirim, tempPesan, konfirmasiPost string

	for i := 0; i < SIZE; i++ {
		if ArrPengirim[i] == "" {
			tempNamaPengirim = bacaInput(fmt.Sprintf("Masukkan nama pengirim : "))
			tempWaktuKirim = bacaInput(fmt.Sprintf("Masukkan waktu kirim pesan : "))
			tempPesan = bacaInput(fmt.Sprintf("Masukkan pesan : "))

			fmt.Print("Apakah anda yakin ingin mengirim pesan ini? [y/n] : ")
			fmt.Scan(&konfirmasiPost)
			if konfirmasiPost == "y" || konfirmasiPost == "Y" {
				ArrPengirim[i] = tempNamaPengirim
				ArrWaktuKirim[i] = tempWaktuKirim
				ArrPesan[i] = tempPesan
				fmt.Println("Pesan berhasil diposting")
			} else if konfirmasiPost == "n" || konfirmasiPost == "N" {
				fmt.Println("Posting pesan dibatalkan")
			}
			break
		}
	}
}

func main() {
	var dataMurid *Murid
	var dataGuru *Guru
	var dataMatkul *Matkul
	var dataJawabanMatkul *jawabanMatkul
	var dataForum *forum
	var pilihanLogin int
	var username, password string
	var pilihanMenu, pilihanMenuMatkul int

	//isi data array guru
	guru[0] = Guru{"Dhimas Hafizh", "NIP123", "GuruDhimas", "dhimas", "Matematika"}
	guru[1] = Guru{"Satrio Wibowo", "NIP456", "GuruSatrio", "satrio", "IPA"}
	guru[2] = Guru{"Azaria Nanda", "NIP789", "GuruNanda", "nanda", "Bahasa Indonesia"}

	//isi data array murid
	murid[0] = Murid{"Dhimas Hafizh", "NIS123", "XII-IPA-1", "dhimas", "dhimas"}
	murid[1] = Murid{"Satrio Wibowo", "NIS456", "XII-IPA-1", "saubtrio", "satrio"}
	murid[2] = Murid{"Azaria Nanda", "NIS789", "XII-IPA-1", "nanda", "nanda"}

MenuLogin:
	fmt.Println("--- LMS SEDERHANA ---")
	fmt.Println("1. Login Guru")
	fmt.Println("2. Login Murid")
	fmt.Println("3. Keluar")
	fmt.Print("Masukkan pilihan anda : ")
	fmt.Scan(&pilihanLogin)
	fmt.Println()

	switch pilihanLogin {
	case 1:
		fmt.Print("Masukkan username : ")
		fmt.Scan(&username)
		fmt.Print("Masukkan password : ")
		fmt.Scan(&password)

		// Cari data guru
		for i := 0; i < SIZE; i++ {
			if guru[i].username == username && guru[i].password == password { //cek login
				dataGuru = &guru[i]
				break
			}
		}

		if dataGuru != nil {
			fmt.Println()
			fmt.Println("Selamat Datang", dataGuru.nama)
			fmt.Println("Mata kuliah yang anda ajar :", dataGuru.namaMatkul)
			if dataGuru.nama == "Dhimas Hafizh" {
				dataMatkul = &matkulMatematika
				dataForum = &forumMatematika
			} else if dataGuru.nama == "Satrio Wibowo" {
				dataMatkul = &matkulIPA
				dataForum = &forumIPA
			} else if dataGuru.nama == "Azaria Nanda" {
				dataMatkul = &matkulBIndo
				dataForum = &forumBIndo
			}
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
		fmt.Println("11. Keluar")
		fmt.Print("Pilihan anda : ")
		fmt.Scan(&pilihanMenu)
		fmt.Println()

		switch pilihanMenu {
		case 1: //tampilan jika memilih menu 1. Tambah Tugas
			var tempNamaTugas string
			for i := 0; i < SIZE; i++ {
				if dataMatkul.namaTugas[i] == "" { // Cek apakah ada data kosong pada array namaTugas
					fmt.Print("Masukkan nama Tugas: ")
					fmt.Scan(&dataMatkul.namaTugas[i])
					tempNamaTugas = dataMatkul.namaTugas[i]
					break
				}
			}

			fmt.Println("!! PERHATIAN !!")
			fmt.Println("Jumlah pertanyaan tugas yang bisa dimasukkan adalah 10 pertanyaan.")
			fmt.Println("Masukkan 'stop' untuk berhenti apabila ingin menambahkan pertanyaan tugas kurang dari 10.")

			InputPertanyaanTugasQuiz(dataMatkul.namaTugas[:], tempNamaTugas, dataMatkul.pertanyaanTugas[:], dataMatkul.jawabanPertanyaanTugas[:], dataMatkul.bobotNilaiTugas[:])

			fmt.Println()
			goto MenuGuru

		case 2: //tampilan jika memilih menu 2. Tambah Quiz
			var tempNamaQuiz string
			for i := 0; i < SIZE; i++ {
				if dataMatkul.namaQuiz[i] == "" { // Cek apakah ada data kosong pada array namaQuiz, jika indeks i kosong maka lakukan input
					dataMatkul.namaQuiz[i] = bacaInput(fmt.Sprintf("Masukkan nama Quiz : ")) //input nama quiz
					tempNamaQuiz = dataMatkul.namaQuiz[i]
					break
				}
			}

			fmt.Println("!! PERHATIAN !!")
			fmt.Println("Jumlah pertanyaan quiz yang bisa dimasukkan adalah 10 pertanyaan")
			fmt.Println("Masukkan 'stop' untuk berhenti apabila ingin menambahkan pertanyaan quiz kurang dari 10")

			InputPertanyaanTugasQuiz(dataMatkul.namaQuiz[:], tempNamaQuiz, dataMatkul.pertanyaanQuiz[:], dataMatkul.jawabanPertanyaanQuiz[:], dataMatkul.bobotNilaiQuiz[:])

			fmt.Println()
			goto MenuGuru

		case 3: //tampilan jika memilih menu 3. Edit Tugas

		case 4: //tampilan jika memilih menu 4. Edit Quiz
			if isEmpty(dataMatkul.namaQuiz[:]) { //cek array namaQuiz kosong atau tidak
				fmt.Println("Quiz kosong!")
				fmt.Println("Tambah quiz terlebih dahulu!")

			} else { //kalo ga kosong, tampilkan isi array namaQuiz dan lakukan penghapusan
				fmt.Println("Data Quiz yang tersimpan :")
				tampilkan(dataMatkul.namaQuiz[:])
				fmt.Println()

				EditTugasQuiz(dataMatkul.namaQuiz[:], dataMatkul.pertanyaanQuiz[:], dataMatkul.jawabanPertanyaanQuiz[:], dataMatkul.bobotNilaiQuiz[:])
			}
			fmt.Println()
			goto MenuGuru

		case 5: //tampilan jika memilih menu 5. Hapus Tugas

		case 6: //tampilan jika memilih menu 6. Hapus Quiz
			if isEmpty(dataMatkul.namaQuiz[:]) { //cek array namaQuiz kosong atau tidak
				fmt.Println("Quiz kosong!")
				fmt.Println("Tambah quiz terlebih dahulu!")
			} else { //kalo ga kosong, tampilkan isi array namaQuiz
				fmt.Println("Data Quiz yang tersimpan :")
				tampilkan(dataMatkul.namaQuiz[:])
			}

			HapusTugasQuiz(dataMatkul.namaQuiz[:], dataMatkul.pertanyaanQuiz[:], dataMatkul.jawabanPertanyaanQuiz[:], dataMatkul.bobotNilaiQuiz[:])
			fmt.Println()
			goto MenuGuru

		case 7: //tampilan jika memilih menu 7. Tampilkan List Tugas

		case 8: //tampilan jika memilih menu 8. Tampilkan List Quiz
			if isEmpty(dataMatkul.namaQuiz[:]) { //cek array namaQuiz kosong atau tidak
				fmt.Println("Quiz kosong!")
				fmt.Println("Tambah quiz terlebih dahulu!")
			} else { //kalo ga kosong, tampilkan isi array namaQuiz
				fmt.Println("Data Quiz yang tersimpan :")
				tampilkan(dataMatkul.namaQuiz[:])
			}
			fmt.Println()
			goto MenuGuru

		case 9: //tampilan jika memilih menu 9. Forum Mata Kuliah
			fmt.Println("-- Forum Mata Kuliah", dataGuru.namaMatkul, "--")
			if !isEmpty(dataForum.pengirim[:]) { //kalo forum tidak kosong, tampilkan isi forum
				for i := 0; i < SIZE; i++ {
					if dataForum.pengirim[i+1] != "" {
						fmt.Println("Postingan ke-", i+1)
						fmt.Println("Pengirim : ", dataForum.pengirim[i])
						fmt.Println("Waktu Kirim : ", dataForum.waktuKirim[i])
						fmt.Println("Pesan : ", dataForum.pesan[i])
						fmt.Println()
					} else {
						break
					}
				}
			}
			PostForumMatkul(dataForum.pengirim[:], dataForum.waktuKirim[:], dataForum.pesan[:])
			fmt.Println()
			goto MenuGuru

		case 10: //tampilan jika memilih menu 10. Lihat List Nilai Murid

		case 11: //tampilan jika memilih menu 11. Keluar
			fmt.Println()
			goto MenuLogin

		default: //tampilan jika memilih selain menu 1-10
			fmt.Println("Pilihan yang dimasukkan tidak valid")
			fmt.Println()
			goto MenuGuru
		}

	case 2: //tampilan jika memilih menu 2. Login Murid
		fmt.Print("Masukkan username : ")
		fmt.Scan(&username)
		fmt.Print("Masukkan password : ")
		fmt.Scan(&password)

		for i := 0; i < SIZE; i++ {
			if murid[i].username == username && murid[i].password == password { //cek login
				dataMurid = &murid[i]
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

	MenuHomeMurid:
		fmt.Println("1. Matematika")
		fmt.Println("2. IPA")
		fmt.Println("3. Bahasa Indonesia")
		fmt.Println("4. Keluar")
		fmt.Print("Masukkan pilihan anda : ")
		fmt.Scan(&pilihanMenu)
		fmt.Println()

		switch pilihanMenu {
		case 1: //tampilan jika memilih menu 1. Matematika
			dataJawabanMatkul = &jawabanMatkulMatematika
			dataMatkul = &matkulMatematika
			dataGuru = &guru[0]
		MenuMatkulMatematika:
			fmt.Println("1. Kerjakan Tugas")
			fmt.Println("2. Kerjakan Quiz")
			fmt.Println("3. Forum matkul")
			fmt.Println("4. Tampilkan nilai rata-rata")
			fmt.Println("5. Kembali ke home")
			fmt.Print("Masukkan pilihan anda : ")
			fmt.Scan(&pilihanMenuMatkul)
			fmt.Println()

			switch pilihanMenuMatkul {
			case 1: //tampilan jika memilih menu 1. Kerjakan Tugas

			case 2: //tampilan jika memilih menu 2. Kerjakan Quiz
				if isEmpty(dataMatkul.namaQuiz[:]) { //kalo array namaQuiz kosong, tampilkan informasi bahwa quiz tidak tersedia
					fmt.Println("Quiz Tidak Tersedia!")
					fmt.Println("Hubungi guru anda untuk menambahkan quiz")
				} else { //kalo array namaQuiz tidak kosong, tampilkan list quiz yang tersedia dan lakukan pengerjaan quiz
					fmt.Println("List Quiz Tersededia :")
					tampilkan(dataMatkul.namaQuiz[:])

					KerjakanTugasQuiz(dataMatkul.namaQuiz[:], dataMatkul.pertanyaanQuiz[:], dataJawabanMatkul.jawabanPertanyaanQuiz[:], dataMatkul.jawabanPertanyaanQuiz[:], dataJawabanMatkul.nilaiPertanyaanQuiz[:], dataMatkul.bobotNilaiQuiz[:], dataJawabanMatkul.namaMuridQuiz[:], dataMurid.nama, dataJawabanMatkul.nilaiQuiz[:])
				}
				fmt.Println()
				goto MenuMatkulMatematika

			case 3: //tampilan jika memilih menu 3. Forum matkul
				fmt.Println("-- Forum Mata Kuliah", dataGuru.namaMatkul, "--")
				if !isEmpty(dataForum.pengirim[:]) { //kalo forum tidak kosong, tampilkan isi forum
					for i := 0; i < SIZE; i++ {
						if dataForum.pengirim[i+1] != "" {
							fmt.Println("Postingan ke-", i+1)
							fmt.Println("Pengirim : ", dataForum.pengirim[i])
							fmt.Println("Waktu Kirim : ", dataForum.waktuKirim[i])
							fmt.Println("Pesan : ", dataForum.pesan[i])
							fmt.Println()
						} else {
							break
						}
					}
				}
				PostForumMatkul(dataForum.pengirim[:], dataForum.waktuKirim[:], dataForum.pesan[:])
				fmt.Println()
				goto MenuMatkulMatematika

			case 4: //tampilan jika memilih menu 4. Tampilkan nilai rata-rata
				TampilkanNilaiRataRataMatkul(dataMatkul.namaTugas[:], dataMatkul.namaQuiz[:], dataJawabanMatkul.nilaiTugas[:], dataJawabanMatkul.nilaiQuiz[:], dataJawabanMatkul.nilaiMatkul)
				fmt.Println()
				goto MenuMatkulMatematika

			case 5: //tampilan jika memilih menu 5. Kembali ke home
				fmt.Println()
				goto MenuHomeMurid

			default: //tampilan jika memilih selain menu 1-5
				fmt.Println("Pilihan yang dimasukkan tidak valid")
				fmt.Println()
				goto MenuMatkulMatematika
			}

		case 2: //tampilan jika memilih menu 2. IPA
			dataJawabanMatkul = &jawabanMatkulIPA
			dataMatkul = &matkulIPA
			dataGuru = &guru[1]
		MenuMatkulIPA:
			fmt.Println("1. Kerjakan Tugas")
			fmt.Println("2. Kerjakan Quiz")
			fmt.Println("3. Forum matkul")
			fmt.Println("4. Tampilkan nilai rata-rata")
			fmt.Println("5. Kembali ke home")
			fmt.Print("Masukkan pilihan anda : ")
			fmt.Scan(&pilihanMenuMatkul)
			fmt.Println()

			switch pilihanMenuMatkul {
			case 1: //tampilan jika memilih menu 1. Kerjakan Tugas

			case 2: //tampilan jika memilih menu 2. Kerjakan Quiz
				if isEmpty(dataMatkul.namaQuiz[:]) { //kalo array namaQuiz kosong, tampilkan informasi bahwa quiz tidak tersedia
					fmt.Println("Quiz Tidak Tersedia!")
					fmt.Println("Hubungi guru anda untuk menambahkan quiz")
				} else { //kalo array namaQuiz tidak kosong, tampilkan list quiz yang tersedia dan lakukan pengerjaan quiz
					fmt.Println("List Quiz Tersededia :")
					tampilkan(dataMatkul.namaQuiz[:])

					KerjakanTugasQuiz(dataMatkul.namaQuiz[:], dataMatkul.pertanyaanQuiz[:], dataJawabanMatkul.jawabanPertanyaanQuiz[:], dataMatkul.jawabanPertanyaanQuiz[:], dataJawabanMatkul.nilaiPertanyaanQuiz[:], dataMatkul.bobotNilaiQuiz[:], dataJawabanMatkul.namaMuridQuiz[:], dataMurid.nama, dataJawabanMatkul.nilaiQuiz[:])
				}
				fmt.Println()
				goto MenuMatkulIPA

			case 3: //tampilan jika memilih menu 3. Forum matkul
				fmt.Println("-- Forum Mata Kuliah", dataGuru.namaMatkul, "--")
				if !isEmpty(dataForum.pengirim[:]) { //kalo forum tidak kosong, tampilkan isi forum
					for i := 0; i < SIZE; i++ {
						if dataForum.pengirim[i+1] != "" {
							fmt.Println("Postingan ke-", i+1)
							fmt.Println("Pengirim : ", dataForum.pengirim[i])
							fmt.Println("Waktu Kirim : ", dataForum.waktuKirim[i])
							fmt.Println("Pesan : ", dataForum.pesan[i])
							fmt.Println()
						} else {
							break
						}
					}
				}
				PostForumMatkul(dataForum.pengirim[:], dataForum.waktuKirim[:], dataForum.pesan[:])
				fmt.Println()
				goto MenuMatkulIPA

			case 4: //tampilan jika memilih menu 4. Tampilkan nilai rata-rata
				TampilkanNilaiRataRataMatkul(dataMatkul.namaTugas[:], dataMatkul.namaQuiz[:], dataJawabanMatkul.nilaiTugas[:], dataJawabanMatkul.nilaiQuiz[:], dataJawabanMatkul.nilaiMatkul)
				fmt.Println()
				goto MenuMatkulIPA

			case 5: //tampilan jika memilih menu 5. Kembali ke home
				fmt.Println()
				goto MenuHomeMurid

			default: //tampilan jika memilih selain menu 1-5
				fmt.Println("Pilihan yang dimasukkan tidak valid")
				fmt.Println()
				goto MenuMatkulIPA
			}

		case 3: //tampilan jika memilih menu 3. Bahasa Indonesia
			dataJawabanMatkul = &jawabanMatkulBIndo
			dataMatkul = &matkulBIndo
			dataGuru = &guru[2]
		MenuMatkulBIndo:
			fmt.Println("1. Kerjakan Tugas")
			fmt.Println("2. Kerjakan Quiz")
			fmt.Println("3. Forum matkul")
			fmt.Println("4. Tampilkan nilai rata-rata")
			fmt.Println("5. Kembali ke home")
			fmt.Print("Masukkan pilihan anda : ")
			fmt.Scan(&pilihanMenuMatkul)
			fmt.Println()

			switch pilihanMenuMatkul {
			case 1: //tampilan jika memilih menu 1. Kerjakan Tugas

			case 2: //tampilan jika memilih menu 2. Kerjakan Quiz
				if isEmpty(dataMatkul.namaQuiz[:]) { //kalo array namaQuiz kosong, tampilkan informasi bahwa quiz tidak tersedia
					fmt.Println("Quiz Tidak Tersedia!")
					fmt.Println("Hubungi guru anda untuk menambahkan quiz")
				} else { //kalo array namaQuiz tidak kosong, tampilkan list quiz yang tersedia dan lakukan pengerjaan quiz
					fmt.Println("List Quiz Tersededia :")
					tampilkan(dataMatkul.namaQuiz[:])

					KerjakanTugasQuiz(dataMatkul.namaQuiz[:], dataMatkul.pertanyaanQuiz[:], dataJawabanMatkul.jawabanPertanyaanQuiz[:], dataMatkul.jawabanPertanyaanQuiz[:], dataJawabanMatkul.nilaiPertanyaanQuiz[:], dataMatkul.bobotNilaiQuiz[:], dataJawabanMatkul.namaMuridQuiz[:], dataMurid.nama, dataJawabanMatkul.nilaiQuiz[:])
				}
				fmt.Println()
				goto MenuMatkulBIndo

			case 3: //tampilan jika memilih menu 3. Forum matkul
				fmt.Println("-- Forum Mata Kuliah", dataGuru.namaMatkul, "--")
				if !isEmpty(dataForum.pengirim[:]) { //kalo forum tidak kosong, tampilkan isi forum
					for i := 0; i < SIZE; i++ {
						if dataForum.pengirim[i+1] != "" {
							fmt.Println("Postingan ke-", i+1)
							fmt.Println("Pengirim : ", dataForum.pengirim[i])
							fmt.Println("Waktu Kirim : ", dataForum.waktuKirim[i])
							fmt.Println("Pesan : ", dataForum.pesan[i])
							fmt.Println()
						} else {
							break
						}
					}
				}
				PostForumMatkul(dataForum.pengirim[:], dataForum.waktuKirim[:], dataForum.pesan[:])
				fmt.Println()
				goto MenuMatkulBIndo

			case 4: //tampilan jika memilih menu 4. Tampilkan nilai rata-rata
				TampilkanNilaiRataRataMatkul(dataMatkul.namaTugas[:], dataMatkul.namaQuiz[:], dataJawabanMatkul.nilaiTugas[:], dataJawabanMatkul.nilaiQuiz[:], dataJawabanMatkul.nilaiMatkul)
				fmt.Println()
				goto MenuMatkulBIndo

			case 5: //tampilan jika memilih menu 5. Kembali ke home
				fmt.Println()
				goto MenuHomeMurid

			default: //tampilan jika memilih selain menu 1-5
				fmt.Println("Pilihan yang dimasukkan tidak valid")
				fmt.Println()
				goto MenuMatkulBIndo
			}
		case 4: //tampilan jika memilih menu 4. Keluar
			fmt.Println()
			goto MenuLogin

		default: //tampilan jika memilih selain menu 1-4
			fmt.Println("Pilihan yang dimasukkan tidak valid")
			fmt.Println()
			goto MenuHomeMurid
		}
	case 3: //tampilan jika memilih menu 3. Keluar
		fmt.Println("Terima kasih telah menggunakan LMS")

	default: //tampilan jika memilih selain menu 1-3
		fmt.Println("Pilihan yang dimasukkan tidak valid")
		fmt.Println()
		goto MenuLogin
	}
}
