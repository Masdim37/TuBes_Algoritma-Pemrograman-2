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
	jawabanPertanyaanTugas [SIZE]string
	jawabanPertanyaanQuiz  [SIZE]string
	nilaiPertanyaanTugas   [SIZE]int
	nilaiPertanyaanQuiz    [SIZE]int
	nilaiTugas             [SIZE]float64
	nilaiQuiz              [SIZE]float64
	nilaiMatkul            float64
}

type ArrMurid [SIZE]Murid //membuat tipe data ArrMurid yang merupakan array struct
type ArrGuru [SIZE]Guru   //membuat tipe data ArrGuru yang merupakan array struct

var murid ArrMurid //membuat array bernama murid dengan tipe data ArrMurid (array struct)

var guru ArrGuru //membuat array bernama guru dengan tipe data ArrGuru (array struct)

var matkulMatematika Matkul
var matkulIPA Matkul
var matkulBIndo Matkul

var jawabanMatkulMatematika jawabanMatkul //membuat array bernama matkulMatematika dengan tipe data ArrMatkul (array struct)
var jawabanMatkulIPA jawabanMatkul        //membuat array bernama matkulIPA dengan tipe data ArrMatkul (array struct)
var jawabanMatkulBIndo jawabanMatkul      //membuat array bernama matkulBIndo dengan tipe data ArrMatkul (array struct)

func isEmpty(arrData []string) bool { //function untuk cek apakah array kosong atau tidak
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

func tampilkan(arrData []string) { //prosedur untuk menampilkan isi array
	for i := 0; i < len(arrData); i++ {
		fmt.Printf("%d. %s\n", i+1, arrData[i])
		if arrData[i+1] == "" {
			break
		}
	}
}

func main() {
	var dataMurid *Murid
	var dataGuru *Guru
	var dataMatkul *Matkul
	var dataJawabanMatkul *jawabanMatkul
	var pilihanLogin int
	var username, password string
	var pilihanMenu, pilihanMenuMatkul int

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
			} else if dataGuru.nama == "Satrio Wibowo" {
				dataMatkul = &matkulIPA
			} else if dataGuru.nama == "Azaria Nanda" {
				dataMatkul = &matkulBIndo
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

		case 2: //tampilan jika memilih menu 2. Tambah Quiz
			var tempNamaQuiz string
			var tempPertanyaan string
			for i := 0; i < SIZE; i++ {
				if dataMatkul.namaQuiz[i] == "" { // Cek apakah ada data kosong pada array namaQuiz, jika indeks i kosong maka lakukan input
					fmt.Print("Masukkan nama Quiz : ")
					fmt.Scan(&dataMatkul.namaQuiz[i])
					tempNamaQuiz = dataMatkul.namaQuiz[i]
					break
				}
			}

			fmt.Println("!! PERHATIAN !!")
			fmt.Println("Jumlah pertanyaan quiz yang bisa dimasukkan adalah 10 pertanyaan")
			fmt.Println("Masukkan 'stop' untuk berhenti apabila ingin menambahkan pertanyaan quiz kurang dari 10")

			for i := 0; i < SIZE; i++ {
				if dataMatkul.namaQuiz[i] == tempNamaQuiz {
					for j := 0; j < 10; j++ {
						fmt.Print("Masukkan pertanyaan ke-", j+1, " : ")
						fmt.Scan(&tempPertanyaan)

						if tempPertanyaan == "stop" { //kalo inputan 'stop', maka berhenti
							break
						}

						dataMatkul.pertanyaanQuiz[i*10+j] = tempPertanyaan //kalo inputan bukan 'stop', maka masukkan ke array pertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode pertanyaan quiz)

						fmt.Print("Masukkan kunci jawaban pertanyaan ke-", j+1, " : ") //Input jawaban pertanyaan quiz
						fmt.Scan(&dataMatkul.jawabanPertanyaanQuiz[i*10+j])            //Kunci jawaban disimpan ke array jawabanPertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode kunci jawaban pertanyaan quiz)

						fmt.Print("Masukkan bobot nilai pertanyaan ke-", j+1, " : ") //input bobot nilai pertanyaan quiz
						fmt.Scan(&dataMatkul.bobotNilaiQuiz[i*10+j])                 //Bobot nilai pertanyaan disimpan ke array bobotNilaiQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode bobot nilai pertanyaan quiz)
					}
				}
			}
			fmt.Println()
			goto MenuGuru

		case 3: //tampilan jika memilih menu 3. Edit Tugas

		case 4: //tampilan jika memilih menu 4. Edit Quiz

		case 5: //tampilan jika memilih menu 5. Hapus Tugas

		case 6: //tampilan jika memilih menu 6. Hapus Quiz
			var pilihanHapusQuiz, konfirmasiHapusQuiz string

			if isEmpty(dataMatkul.namaQuiz[:]) { //cek array namaQuiz kosong atau tidak
				fmt.Println("Quiz kosong!")
				fmt.Println("Tambah quiz terlebih dahulu!")
			} else { //kalo ga kosong, tampilkan isi array namaQuiz dan lakukan penghapusan
				fmt.Println("Data Quiz yang tersimpan :")
				tampilkan(dataMatkul.namaQuiz[:])
				fmt.Println()

				fmt.Print("Masukkan nama quiz yang ingin dihapus : ")
				fmt.Scan(&pilihanHapusQuiz)
				fmt.Print("Apakah anda yakin ingin mengapus quiz ", pilihanHapusQuiz, " ? [y/n] : ")
				fmt.Scan(&konfirmasiHapusQuiz)
				if konfirmasiHapusQuiz == "y" || konfirmasiHapusQuiz == "Y" { //kalo masukin konfirmasi y, lakukan penghapusan
					for i := 0; i < SIZE; i++ {
						if dataMatkul.namaQuiz[i] == pilihanHapusQuiz { //kalo nama quiz yang diinputkan sama dengan nama quiz yang ada di array namaQuiz, maka lakukan penghapusan
							dataMatkul.namaQuiz[i] = "" //hapus nama quiz
							for j := 0; j < 10; j++ {
								dataMatkul.pertanyaanQuiz[i*10+j] = ""        //hapus pertanyaan quiz (i = kode quiz, 10 = maks pertanyaan, j = kode pertanyaan quiz)
								dataMatkul.jawabanPertanyaanQuiz[i*10+j] = "" //hapus kunci jawaban quiz (i = kode quiz, 10 = maks pertanyaan, j = kode kunci jawaban pertanyaan quiz)
								dataMatkul.bobotNilaiQuiz[i*10+j] = 0         // Menghapus bobot nilai pertanyaan quiz (i = kode quiz, 10 = maks pertanyaan, j = kode bobot nilai pertanyaan quiz)
							}
							fmt.Println("Quiz", pilihanHapusQuiz, "berhasil dihapus.") //tampilkan informasi bahwa quiz yang dipilih berhasil dihapus
							break                                                      //keluar dari loop
						}
					}
				} else if konfirmasiHapusQuiz == "n" || konfirmasiHapusQuiz == "N" { //kalo masukin konfirmasi n, batalkan penghapusan
					fmt.Print("Penghapusan quiz ", pilihanHapusQuiz, " dibatalkan") //tampilkan informasi bahwa quiz yang dipilih ngga jadi dihapus
				}
				fmt.Println()
				goto MenuGuru
			}

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
				var namaQuiz string
				var totalPoinSoalQuiz int
				var totalPoinJawabanQuiz int
				if isEmpty(dataMatkul.namaQuiz[:]) { //kalo array namaQuiz kosong, tampilkan informasi bahwa quiz tidak tersedia
					fmt.Println("Quiz Tidak Tersedia!")
					fmt.Println("Hubungi guru anda untuk menambahkan quiz")
					fmt.Println()
					goto MenuMatkulMatematika
				} else { //kalo array namaQuiz tidak kosong, tampilkan list quiz yang tersedia dan lakukan pengerjaan quiz
					fmt.Println("List Quiz Tersededia :")
					tampilkan(dataMatkul.namaQuiz[:])

					fmt.Print("Masukkan nama quiz yang ingin dikerjakan : ")
					fmt.Scan(&namaQuiz)
					fmt.Println()
					for i := 0; i < SIZE; i++ {
						if dataMatkul.namaQuiz[i] == namaQuiz { //kalo nama quiz yang diinputkan sama dengan nama quiz yang ada di array namaQuiz, maka lakukan pengerjaan quiz
							totalPoinSoalQuiz = 0    //ini nilai poin total soal quiz
							totalPoinJawabanQuiz = 0 //ini nilai poin total jawaban quiz siswa

							for j := 0; j < 10; j++ {
								fmt.Println(dataMatkul.pertanyaanQuiz[i*10+j])                                                   //tampilkan pertanyaan quiz di array pertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode pertanyaan quiz)
								fmt.Scan(&dataJawabanMatkul.jawabanPertanyaanQuiz[i*10+j])                                       //input jawaban siswa, simpan di array jawabanPertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode jawaban pertanyaan quiz)
								if dataJawabanMatkul.jawabanPertanyaanQuiz[i*10+j] == dataMatkul.jawabanPertanyaanQuiz[i*10+j] { //kalo jawaban siswa sama dengan kunci jawaban yang ada di array jawabanPertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode kunci jawaban pertanyaan quiz)
									dataJawabanMatkul.nilaiPertanyaanQuiz[i*10+j] += dataMatkul.bobotNilaiQuiz[i*10+j] //maka nilai jawaban siswa di array nilaiPertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode nilai jawaban pertanyaan quiz) ditambahkan dengan bobot nilai quiz yang ada di array bobotNilaiQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode bobot nilai quiz)
								} else { //kalo jawabannya ga sama kaya kunci jawaban
									dataJawabanMatkul.nilaiPertanyaanQuiz[i*10+j] += 0 //maka nilai jawaban siswa di array nilaiPertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode nilai jawaban pertanyaan quiz) ditambahkan dengan 0
								}
								totalPoinSoalQuiz += dataMatkul.bobotNilaiQuiz[i*10+j]                //hitung total poin soal quiz
								totalPoinJawabanQuiz += dataJawabanMatkul.nilaiPertanyaanQuiz[i*10+j] //hitung total poin jawaban quiz siswa
							}

							dataJawabanMatkul.nilaiQuiz[i] = float64(totalPoinJawabanQuiz) / float64(totalPoinSoalQuiz) * 100 //perhitungan nilai akhir quiz siswa, disimpan di array nilaiQuiz indeks ke i (i = kode quiz)
							fmt.Println("Nilai Quiz Anda :", dataJawabanMatkul.nilaiQuiz[i])                                  //tampilkan nilai quiz
							fmt.Println()
							goto MenuMatkulMatematika
						}
					}
				}

			case 3: //tampilan jika memilih menu 3. Forum matkul

			case 4: //tampilan jika memilih menu 4. Tampilkan nilai rata-rata
				fmt.Println("Nilai Rata-Rata Tugas :") //hitungnya penjumlahan semua nilai tugas dibagi jumlah tugas yang ada
				//data semua nilai tugas ada di dataJawabanMatkul.nilaiTugas, jumlah tugas dapetnya dari len(dataMatkul.namaTugas)
				fmt.Println("Nilai Rata-Rata Quiz :") //hitungnya penjumlahan semua nilai quiz dibagi jumlah quiz yang ada
				//data semua nilai quiz ada di dataJawabanMatkul.nilaiQuiz, jumlah quiz dapetnya dari len(dataMatkul.namaQuiz)
				fmt.Println("Nilai Rata-Rata Mata Kuliah :") //dapetnya dari (nilai rata2 tugas + nilai rata2 quiz) / 2

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
				var namaQuiz string
				var totalPoinSoalQuiz int
				var totalPoinJawabanQuiz int
				if isEmpty(dataMatkul.namaQuiz[:]) {
					fmt.Println("Quiz Tidak Tersedia!")
					fmt.Println("Hubungi guru anda untuk menambahkan quiz")
					fmt.Println()
					goto MenuMatkulIPA
				} else {
					fmt.Println("List Quiz Tersededia :")
					tampilkan(dataMatkul.namaQuiz[:])

					fmt.Print("Masukkan nama quiz yang ingin dikerjakan : ")
					fmt.Scan(&namaQuiz)
					fmt.Println()
					for i := 0; i < SIZE; i++ {
						if dataMatkul.namaQuiz[i] == namaQuiz { //kalo nama quiz yang diinputkan sama dengan nama quiz yang ada di array namaQuiz, maka lakukan pengerjaan quiz
							totalPoinSoalQuiz = 0    //ini nilai poin total soal quiz
							totalPoinJawabanQuiz = 0 //ini nilai poin total jawaban quiz siswa

							for j := 0; j < 10; j++ {
								fmt.Println(dataMatkul.pertanyaanQuiz[i*10+j])                                                   //tampilkan pertanyaan quiz di array pertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode pertanyaan quiz)
								fmt.Scan(&dataJawabanMatkul.jawabanPertanyaanQuiz[i*10+j])                                       //input jawaban siswa, simpan di array jawabanPertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode jawaban pertanyaan quiz)
								if dataJawabanMatkul.jawabanPertanyaanQuiz[i*10+j] == dataMatkul.jawabanPertanyaanQuiz[i*10+j] { //kalo jawaban siswa sama dengan kunci jawaban yang ada di array jawabanPertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode kunci jawaban pertanyaan quiz)
									dataJawabanMatkul.nilaiPertanyaanQuiz[i*10+j] += dataMatkul.bobotNilaiQuiz[i*10+j] //maka nilai jawaban siswa di array nilaiPertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode nilai jawaban pertanyaan quiz) ditambahkan dengan bobot nilai quiz yang ada di array bobotNilaiQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode bobot nilai quiz)
								} else { //kalo jawabannya ga sama kaya kunci jawaban
									dataJawabanMatkul.nilaiPertanyaanQuiz[i*10+j] += 0 //maka nilai jawaban siswa di array nilaiPertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode nilai jawaban pertanyaan quiz) ditambahkan dengan 0
								}
								totalPoinSoalQuiz += dataMatkul.bobotNilaiQuiz[i*10+j]                //hitung total poin soal quiz
								totalPoinJawabanQuiz += dataJawabanMatkul.nilaiPertanyaanQuiz[i*10+j] //hitung total poin jawaban quiz siswa
							}

							dataJawabanMatkul.nilaiQuiz[i] = float64(totalPoinJawabanQuiz) / float64(totalPoinSoalQuiz) * 100 //perhitungan nilai akhir quiz siswa, disimpan di array nilaiQuiz indeks ke i (i = kode quiz)
							fmt.Println("Nilai Quiz Anda :", dataJawabanMatkul.nilaiQuiz[i])                                  //tampilkan nilai quiz
							fmt.Println()
							goto MenuMatkulIPA
						}
					}
				}

			case 3: //tampilan jika memilih menu 3. Forum matkul

			case 4: //tampilan jika memilih menu 4. Tampilkan nilai rata-rata
				fmt.Println("Nilai Rata-Rata Tugas :") //hitungnya penjumlahan semua nilai tugas dibagi jumlah tugas yang ada
				//data semua nilai tugas ada di dataJawabanMatkul.nilaiTugas, jumlah tugas dapetnya dari len(dataMatkul.namaTugas)
				fmt.Println("Nilai Rata-Rata Quiz :") //hitungnya penjumlahan semua nilai quiz dibagi jumlah quiz yang ada
				//data semua nilai quiz ada di dataJawabanMatkul.nilaiQuiz, jumlah quiz dapetnya dari len(dataMatkul.namaQuiz)
				fmt.Println("Nilai Rata-Rata Mata Kuliah :") //dapetnya dari (nilai rata2 tugas + nilai rata2 quiz) / 2

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
				var namaQuiz string
				var totalPoinSoalQuiz int
				var totalPoinJawabanQuiz int
				if isEmpty(dataMatkul.namaQuiz[:]) {
					fmt.Println("Quiz Tidak Tersedia!")
					fmt.Println("Hubungi guru anda untuk menambahkan quiz")
					fmt.Println()
					goto MenuMatkulBIndo
				} else {
					fmt.Println("List Quiz Tersedia :")
					tampilkan(dataMatkul.namaQuiz[:])

					fmt.Print("Masukkan nama quiz yang ingin dikerjakan : ")
					fmt.Scan(&namaQuiz)
					fmt.Println()
					for i := 0; i < SIZE; i++ {
						if dataMatkul.namaQuiz[i] == namaQuiz { //kalo nama quiz yang diinputkan sama dengan nama quiz yang ada di array namaQuiz, maka lakukan pengerjaan quiz
							totalPoinSoalQuiz = 0    //ini nilai poin total soal quiz
							totalPoinJawabanQuiz = 0 //ini nilai poin total jawaban quiz siswa

							for j := 0; j < 10; j++ {
								fmt.Println(dataMatkul.pertanyaanQuiz[i*10+j])                                                   //tampilkan pertanyaan quiz di array pertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode pertanyaan quiz)
								fmt.Scan(&dataJawabanMatkul.jawabanPertanyaanQuiz[i*10+j])                                       //input jawaban siswa, simpan di array jawabanPertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode jawaban pertanyaan quiz)
								if dataJawabanMatkul.jawabanPertanyaanQuiz[i*10+j] == dataMatkul.jawabanPertanyaanQuiz[i*10+j] { //kalo jawaban siswa sama dengan kunci jawaban yang ada di array jawabanPertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode kunci jawaban pertanyaan quiz)
									dataJawabanMatkul.nilaiPertanyaanQuiz[i*10+j] += dataMatkul.bobotNilaiQuiz[i*10+j] //maka nilai jawaban siswa di array nilaiPertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode nilai jawaban pertanyaan quiz) ditambahkan dengan bobot nilai quiz yang ada di array bobotNilaiQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode bobot nilai quiz)
								} else { //kalo jawabannya ga sama kaya kunci jawaban
									dataJawabanMatkul.nilaiPertanyaanQuiz[i*10+j] += 0 //maka nilai jawaban siswa di array nilaiPertanyaanQuiz indeks ke i*10+j (i = kode quiz, 10 = maks pertanyaan, j = kode nilai jawaban pertanyaan quiz) ditambahkan dengan 0
								}
								totalPoinSoalQuiz += dataMatkul.bobotNilaiQuiz[i*10+j]                //hitung total poin soal quiz
								totalPoinJawabanQuiz += dataJawabanMatkul.nilaiPertanyaanQuiz[i*10+j] //hitung total poin jawaban quiz siswa
							}

							dataJawabanMatkul.nilaiQuiz[i] = float64(totalPoinJawabanQuiz) / float64(totalPoinSoalQuiz) * 100 //perhitungan nilai akhir quiz siswa, disimpan di array nilaiQuiz indeks ke i (i = kode quiz)
							fmt.Println("Nilai Quiz Anda :", dataJawabanMatkul.nilaiQuiz[i])                                  //tampilkan nilai quiz
							fmt.Println()
							goto MenuMatkulBIndo
						}
					}
				}

			case 3: //tampilan jika memilih menu 3. Forum matkul

			case 4: //tampilan jika memilih menu 4. Tampilkan nilai rata-rata
				fmt.Println("Nilai Rata-Rata Tugas :") //hitungnya penjumlahan semua nilai tugas dibagi jumlah tugas yang ada
				//data semua nilai tugas ada di dataJawabanMatkul.nilaiTugas, jumlah tugas dapetnya dari len(dataMatkul.namaTugas)
				fmt.Println("Nilai Rata-Rata Quiz :") //hitungnya penjumlahan semua nilai quiz dibagi jumlah quiz yang ada
				//data semua nilai quiz ada di dataJawabanMatkul.nilaiQuiz, jumlah quiz dapetnya dari len(dataMatkul.namaQuiz)
				fmt.Println("Nilai Rata-Rata Mata Kuliah :") //dapetnya dari (nilai rata2 tugas + nilai rata2 quiz) / 2

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
