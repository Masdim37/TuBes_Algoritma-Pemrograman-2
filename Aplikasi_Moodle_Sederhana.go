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

// Fungsi bacaInput membaca input dari terminal dengan prompt tertentu
func bacaInput(prompt string) string {
	fmt.Print(prompt) // Menampilkan prompt
	//os.Stdout.Sync()               // Flush output buffer secara manual
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // Membaca input hingga newline
	return strings.TrimSpace(input)     // Menghapus spasi atau newline di awal/akhir
}

// function untuk input pertanyaan tugas dan quiz
func InputPertanyaanTugasQuiz(ArrNamaTugasQuiz []string, namaTugasQuiz string, ArrPertanyaanTugasQuiz []string, ArrJawabanPertanyaanTugasQuiz []string, ArrBobotNilaiPertanyaanTugasQuiz []int) {
	var tempPertanyaan string
loopInputPertanyaan:
	for i := 0; i < SIZE; i++ {
		if ArrNamaTugasQuiz[i] == namaTugasQuiz {
			for j := 0; j < 10; j++ {
				tempPertanyaan = bacaInput(fmt.Sprintf("Masukkan pertanyaan ke-%d : ", j+1)) // Input pertanyaan

				if strings.ToLower(tempPertanyaan) == "stop" { // Berhenti jika input "stop"
					break loopInputPertanyaan
				}

				ArrPertanyaanTugasQuiz[i*10+j] = tempPertanyaan // Masukkan pertanyaan ke array ArrPertanyaanTugasQuiz

				// Input jawaban
				ArrJawabanPertanyaanTugasQuiz[i*10+j] = bacaInput(fmt.Sprintf("Masukkan kunci jawaban pertanyaan ke-%d : ", j+1)) // Input jawaban

				fmt.Printf("Masukkan bobot nilai pertanyaan ke-%d : ", j+1) // Input bobot nilai
				fmt.Scan(&ArrBobotNilaiPertanyaanTugasQuiz[i*10+j])         // Simpan bobot nilai
			}
		}
	}
}

// function untuk menghapus tugas dan quiz
func HapusTugasQuiz(ArrNamaTugasQuiz []string, ArrPertanyaanTugasQuiz []string, ArrJawabanPertanyaanTugasQuiz []string, ArrBobotNilaiPertanyaanTugasQuiz []int) {
	var pilihanHapusTugasQuiz, konfirmasiHapusTugasQuiz string

	pilihanHapusTugasQuiz = bacaInput(fmt.Sprintf("Masukkan nama tugas/quiz yang ingin dihapus : "))
	fmt.Print("Apakah anda yakin ingin mengapus tugas/quiz ", pilihanHapusTugasQuiz, " ? [y/n] : ")
	fmt.Scan(&konfirmasiHapusTugasQuiz)
	if konfirmasiHapusTugasQuiz == "y" || konfirmasiHapusTugasQuiz == "Y" { //kalo masukin konfirmasi y, lakukan penghapusan
		for i := 0; i < SIZE; i++ {
			if ArrNamaTugasQuiz[i] == pilihanHapusTugasQuiz { //kalo nama tugas/quiz yang diinputkan sama dengan nama tugas/quiz yang ada di array ArrNamaTugasQuiz, maka lakukan penghapusan
				ArrNamaTugasQuiz[i] = "" //hapus nama quiz
				for j := 0; j < 10; j++ {
					ArrPertanyaanTugasQuiz[i*10+j] = ""          //hapus pertanyaan tugas/quiz (i = kode tugas/quiz, 10 = maks pertanyaan, j = kode pertanyaan tugas/quiz)
					ArrJawabanPertanyaanTugasQuiz[i*10+j] = ""   //hapus kunci jawaban tugas/quiz (i = kode tugas/quiz, 10 = maks pertanyaan, j = kode pertanyaan tugas/quiz)
					ArrBobotNilaiPertanyaanTugasQuiz[i*10+j] = 0 // Menghapus bobot nilai pertanyaan tugas/quiz (i = kode tugas/quiz, 10 = maks pertanyaan, j = kode pertanyaan tugas/quiz)
				}
				fmt.Println("Tugas/Quiz", pilihanHapusTugasQuiz, "berhasil dihapus.") //tampilkan informasi bahwa tugas/quiz yang dipilih berhasil dihapus
				break                                                                 //keluar dari loop
			}
		}
	} else if konfirmasiHapusTugasQuiz == "n" || konfirmasiHapusTugasQuiz == "N" { //kalo masukin konfirmasi n, batalkan penghapusan
		fmt.Print("Penghapusan tugas/quiz ", pilihanHapusTugasQuiz, " dibatalkan") //tampilkan informasi bahwa tugas/quiz yang dipilih ngga jadi dihapus
	}
}

func EditTugasQuiz(ArrNamaTugasQuiz []string, ArrPertanyaanTugasQuiz []string, ArrJawabanPertanyaanTugasQuiz []string, ArrBobotNilaiPertanyaanTugasQuiz []int) {
	var pilihanEditTugasQuiz string
	var pilihanNomorEditTugasQuiz int

	pilihanEditTugasQuiz = bacaInput(fmt.Sprintf("Masukkan nama tugas/quiz yang ingin diubah : "))
	fmt.Println("List Pertanyaan Tugas/Quiz", pilihanEditTugasQuiz, " :")
	for i := 0; i < SIZE; i++ {
		if ArrNamaTugasQuiz[i] == pilihanEditTugasQuiz {
			for j := 0; j < 10; j++ {
				if ArrPertanyaanTugasQuiz[i*10+j] == "" {
					break
				} else if ArrPertanyaanTugasQuiz[i*10+j] != "" {
					fmt.Println("Pertanyaan ke-", j+1)
					fmt.Println("Pertanyaan : ", ArrPertanyaanTugasQuiz[i*10+j])
					fmt.Println("Kunci Jawaban : ", ArrJawabanPertanyaanTugasQuiz[i*10+j])
					fmt.Println("Bobot Nilai : ", ArrBobotNilaiPertanyaanTugasQuiz[i*10+j])
				}
			}
			fmt.Println()
			fmt.Print("Masukkan nomor pertanyaan yang ingin diubah : ")
			fmt.Scan(&pilihanNomorEditTugasQuiz)
			ArrPertanyaanTugasQuiz[i*10+pilihanNomorEditTugasQuiz-1] = bacaInput(fmt.Sprintf("Masukkan pertanyaan baru : "))
			ArrJawabanPertanyaanTugasQuiz[i*10+pilihanNomorEditTugasQuiz-1] = bacaInput(fmt.Sprintf("Masukkan kunci jawaban baru : "))
			fmt.Print("Masukkan bobot nilai baru : ")
			fmt.Scan(&ArrBobotNilaiPertanyaanTugasQuiz[i*10+pilihanNomorEditTugasQuiz-1])
			fmt.Println("Pertanyaan ke-", pilihanNomorEditTugasQuiz, "pada tugas/quiz", pilihanEditTugasQuiz, "berhasil diubah")
			break
		}
	}
}

func KerjakanTugasQuiz(ArrNamaTugasQuiz []string, ArrPertanyaanTugasQuiz []string, ArrInputJawabanPertanyaanTugasQuiz []string, ArrJawabanPertanyaanTugasQuiz []string, ArrInputNilaiPertanyaanTugasQuiz []int, ArrBobotNilaiPertanyaanTugasQuiz []int, ArrNamaMurid []string, namaMurid string, ArrNilaiTugasQuiz []float64) {
	var namaTugasQuiz string
	var totalPoinSoalTugasQuiz int
	var totalPoinJawabanTugasQuiz int

	namaTugasQuiz = bacaInput(fmt.Sprintf("Masukkan nama tugas/quiz yang ingin dikerjakan : "))
	fmt.Println()
	for i := 0; i < SIZE; i++ {
		if ArrNamaTugasQuiz[i] == namaTugasQuiz { //kalo nama tugas/quiz yang diinputkan sama dengan nama tugas/quiz yang ada di array ArrNamaTugasQuiz, maka lakukan pengerjaan tugas/quiz
			totalPoinSoalTugasQuiz = 0    //ini nilai poin total soal tugas/quiz
			totalPoinJawabanTugasQuiz = 0 //ini nilai poin total jawaban tugas/quiz siswa

			for j := 0; j < 10; j++ {
				if ArrPertanyaanTugasQuiz[i*10+j] != "" {
					fmt.Println(ArrPertanyaanTugasQuiz[i*10+j])                                       //tampilkan pertanyaan tugas/quiz di array ArrPertanyaanTugasQuiz indeks ke i*10+j (i = kode tugas/quiz, 10 = maks pertanyaan, j = kode pertanyaan tugas/quiz)
					ArrInputJawabanPertanyaanTugasQuiz[i*10+j] = bacaInput(fmt.Sprintf("Jawaban : ")) //input jawaban siswa, simpan di array jawabanPertanyaanQuiz indeks ke i*10+j (i = kode tugas/quiz, 10 = maks pertanyaan, j = kode pertanyaan tugas/quiz)
				}

				if ArrInputJawabanPertanyaanTugasQuiz[i*10+j] == ArrJawabanPertanyaanTugasQuiz[i*10+j] { //kalo jawaban siswa sama dengan kunci jawaban yang ada di array ArrJawabanPertanyaanTugasQuiz indeks ke i*10+j (i = kode tugas/quiz, 10 = maks pertanyaan, j = kode pertanyaan tugas/quiz), maka
					ArrInputNilaiPertanyaanTugasQuiz[i*10+j] += ArrBobotNilaiPertanyaanTugasQuiz[i*10+j] //maka nilai jawaban siswa di array ArrInputNilaiPertanyaanTugasQuiz indeks ke i*10+j (i = kode tugas/quiz, 10 = maks pertanyaan, j = kode nilai jawaban pertanyaan tugas/quiz) ditambahkan dengan bobot nilai quiz yang ada di array ArrBobotNilaiPertanyaanTugasQuiz indeks ke i*10+j (i = kode tugas/quiz, 10 = maks pertanyaan, j = kode bobot nilai tugas/quiz)
				} else { //kalo jawabannya ga sama kaya kunci jawaban
					ArrInputNilaiPertanyaanTugasQuiz[i*10+j] += 0 //maka nilai jawaban siswa di array ArrInputNilaiPertanyaanTugasQuiz indeks ke i*10+j (i = kode tugas/quiz, 10 = maks pertanyaan, j = kode nilai jawaban pertanyaan tugas/quiz) ditambahkan dengan 0
				}
				totalPoinSoalTugasQuiz += ArrBobotNilaiPertanyaanTugasQuiz[i*10+j]    //hitung total poin soal quiz
				totalPoinJawabanTugasQuiz += ArrInputNilaiPertanyaanTugasQuiz[i*10+j] //hitung total poin jawaban quiz siswa
			}

			ArrNamaMurid[i] = namaMurid
			ArrNilaiTugasQuiz[i] = float64(totalPoinJawabanTugasQuiz) / float64(totalPoinSoalTugasQuiz) * 100 //perhitungan nilai akhir tugas/quiz siswa, disimpan di array ArrNilaiTugasQuiz indeks ke i (i = kode tugas/quiz)
			fmt.Println("Nilai Tugas/Quiz Anda :", ArrNilaiTugasQuiz[i])                                      //tampilkan nilai tugas/quiz
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

func SelectionSortMurid(arrNilai []float64, arrNamaMurid []string) {
	n := len(arrNilai)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arrNilai[j] < arrNilai[minIdx] {
				minIdx = j
			}
		}
		// Tukar nilai dan nama murid
		arrNilai[i], arrNilai[minIdx] = arrNilai[minIdx], arrNilai[i]
		arrNamaMurid[i], arrNamaMurid[minIdx] = arrNamaMurid[minIdx], arrNamaMurid[i]
	}
}

func SequentialBinarySearchMurid(arrNilai []float64, arrNamaMurid []string, target float64) (int, string) {
	for i := 0; i < len(arrNilai); i++ {
		if arrNilai[i] == target {
			return i, arrNamaMurid[i]
		}
	}
	return -1, ""
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
				if dataMatkul.namaTugas[i] == "" { // Cek apakah ada data kosong pada array namaQuiz, jika indeks i kosong maka lakukan input
					dataMatkul.namaTugas[i] = bacaInput(fmt.Sprintf("Masukkan nama Tugas : ")) //input nama quiz
					tempNamaTugas = dataMatkul.namaTugas[i]
					break
				}
			}

			fmt.Println("!! PERHATIAN !!")
			fmt.Println("Jumlah pertanyaan tugas yang bisa dimasukkan adalah 10 pertanyaan")
			fmt.Println("Masukkan 'stop' untuk berhenti apabila ingin menambahkan pertanyaan tugas kurang dari 10")

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
			if isEmpty(dataMatkul.namaTugas[:]) {
				fmt.Println("Tugas kosong!")
				fmt.Println("Tambah tugas terlebih dahulu!")
			} else {
				fmt.Println("Data Tugas yang tersimpan:")
				tampilkan(dataMatkul.namaTugas[:])
				fmt.Println()

				EditTugasQuiz(dataMatkul.namaTugas[:], dataMatkul.pertanyaanTugas[:], dataMatkul.jawabanPertanyaanTugas[:], dataMatkul.bobotNilaiTugas[:])
			}
			fmt.Println()
			goto MenuGuru

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
			if isEmpty(dataMatkul.namaTugas[:]) {
				fmt.Println("Tugas kosong!")
				fmt.Println("Tambah tugas terlebih dahulu!")
			} else {
				fmt.Println("Data Tugas yang tersimpan:")
				tampilkan(dataMatkul.namaTugas[:])
			}

			HapusTugasQuiz(dataMatkul.namaTugas[:], dataMatkul.pertanyaanTugas[:], dataMatkul.jawabanPertanyaanTugas[:], dataMatkul.bobotNilaiTugas[:])
			fmt.Println()
			goto MenuGuru

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
			if isEmpty(dataMatkul.namaTugas[:]) {
				fmt.Println("Tugas kosong!")
				fmt.Println("Tambah tugas terlebih dahulu!")
			} else {
				fmt.Println("Data Tugas yang tersimpan:")
				tampilkan(dataMatkul.namaTugas[:])
			}
			fmt.Println()
			goto MenuGuru

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
			var pilihan int
			fmt.Println("Pilih jenis nilai:")
			fmt.Println("1. Tugas")
			fmt.Println("2. Quiz")
			fmt.Print("Pilihan Anda: ")
			fmt.Scan(&pilihan)

			var arrNilai []float64
			var arrNamaMurid []string

			// Pilih data nilai dan nama murid
			if pilihan == 1 {
				arrNilai = dataJawabanMatkul.nilaiTugas[:]
				arrNamaMurid = dataJawabanMatkul.namaMuridTugas[:]
			} else if pilihan == 2 {
				arrNilai = dataJawabanMatkul.nilaiQuiz[:]
				arrNamaMurid = dataJawabanMatkul.namaMuridQuiz[:]
			} else {
				fmt.Println("Pilihan tidak valid.")
				goto MenuGuru
			}

			// Cari nilai tertentu terlebih dahulu
			var target float64
			fmt.Print("Masukkan nilai yang ingin dicari: ")
			fmt.Scan(&target)

			index, namaMurid := SequentialBinarySearchMurid(arrNilai, arrNamaMurid, target)
			if index != -1 {
				fmt.Printf("Nilai %.2f ditemukan pada murid: %s di indeks ke-%d (sebelum pengurutan)\n", target, namaMurid, index)
			} else {
				fmt.Println("Nilai tidak ditemukan (sebelum pengurutan).")
			}

			// Urutkan nilai dengan Selection Sort
			SelectionSortMurid(arrNilai, arrNamaMurid)

			// Tampilkan daftar nilai yang sudah diurutkan
			fmt.Println("\n-- Daftar Nilai Murid (Setelah Pengurutan) --")
			for i, nilai := range arrNilai {
				if nilai != 0 { // Tampilkan hanya nilai yang valid
					fmt.Printf("%d. Nama: %s, Nilai: %.2f\n", i+1, arrNamaMurid[i], nilai)
				}
			}

			fmt.Println()
			goto MenuGuru

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
				if isEmpty(dataMatkul.namaTugas[:]) { // Jika array namaTugas kosong
					fmt.Println("Tugas Tidak Tersedia!")
					fmt.Println("Hubungi guru anda untuk menambahkan tugas.")
				} else { // Jika array namaTugas tidak kosong
					fmt.Println("List Tugas Tersedia:")
					tampilkan(dataMatkul.namaTugas[:])

					KerjakanTugasQuiz(
						dataMatkul.namaTugas[:],                     // Nama tugas
						dataMatkul.pertanyaanTugas[:],               // Pertanyaan tugas
						dataJawabanMatkul.jawabanPertanyaanTugas[:], // Jawaban siswa untuk tugas
						dataMatkul.jawabanPertanyaanTugas[:],        // Kunci jawaban tugas
						dataJawabanMatkul.nilaiPertanyaanTugas[:],   // Nilai per pertanyaan tugas
						dataMatkul.bobotNilaiTugas[:],               // Bobot nilai per pertanyaan tugas
						dataJawabanMatkul.namaMuridTugas[:],         // Nama murid untuk tugas
						dataMurid.nama,                              // Nama murid yang sedang login
						dataJawabanMatkul.nilaiTugas[:],             // Nilai total tugas
					)
				}
				fmt.Println()
				goto MenuMatkulMatematika

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
				if isEmpty(dataMatkul.namaTugas[:]) { // Jika array namaTugas kosong
					fmt.Println("Tugas Tidak Tersedia!")
					fmt.Println("Hubungi guru anda untuk menambahkan tugas.")
				} else { // Jika array namaTugas tidak kosong
					fmt.Println("List Tugas Tersedia:")
					tampilkan(dataMatkul.namaTugas[:])

					KerjakanTugasQuiz(
						dataMatkul.namaTugas[:],                     // Nama tugas
						dataMatkul.pertanyaanTugas[:],               // Pertanyaan tugas
						dataJawabanMatkul.jawabanPertanyaanTugas[:], // Jawaban siswa untuk tugas
						dataMatkul.jawabanPertanyaanTugas[:],        // Kunci jawaban tugas
						dataJawabanMatkul.nilaiPertanyaanTugas[:],   // Nilai per pertanyaan tugas
						dataMatkul.bobotNilaiTugas[:],               // Bobot nilai per pertanyaan tugas
						dataJawabanMatkul.namaMuridTugas[:],         // Nama murid untuk tugas
						dataMurid.nama,                              // Nama murid yang sedang login
						dataJawabanMatkul.nilaiTugas[:],             // Nilai total tugas
					)
				}
				fmt.Println()
				goto MenuMatkulIPA

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
				if isEmpty(dataMatkul.namaTugas[:]) { // Jika array namaTugas kosong
					fmt.Println("Tugas Tidak Tersedia!")
					fmt.Println("Hubungi guru anda untuk menambahkan tugas.")
				} else { // Jika array namaTugas tidak kosong
					fmt.Println("List Tugas Tersedia:")
					tampilkan(dataMatkul.namaTugas[:])

					KerjakanTugasQuiz(
						dataMatkul.namaTugas[:],                     // Nama tugas
						dataMatkul.pertanyaanTugas[:],               // Pertanyaan tugas
						dataJawabanMatkul.jawabanPertanyaanTugas[:], // Jawaban siswa untuk tugas
						dataMatkul.jawabanPertanyaanTugas[:],        // Kunci jawaban tugas
						dataJawabanMatkul.nilaiPertanyaanTugas[:],   // Nilai per pertanyaan tugas
						dataMatkul.bobotNilaiTugas[:],               // Bobot nilai per pertanyaan tugas
						dataJawabanMatkul.namaMuridTugas[:],         // Nama murid untuk tugas
						dataMurid.nama,                              // Nama murid yang sedang login
						dataJawabanMatkul.nilaiTugas[:],             // Nilai total tugas
					)
				}
				fmt.Println()
				goto MenuMatkulBIndo

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
