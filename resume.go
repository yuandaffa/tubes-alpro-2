package main

import "fmt"

const NMAX int = 100

type resume struct {
	pengalamanKerja string
	keterampilan    string
	pendidikan      string
	deskripsi       string
}

type pekerjaan struct {
	keyword  string
	industri string
	posisi   string
	gaji     float64
}

type kerja [NMAX]pekerjaan

var kataKunci = []string{"sawah", "traktor", "pemrograman", "laut", "kandang", "jaringan komputer"}

var daftarPekerjaan = []pekerjaan{
	{keyword: "sawah", industri: "Pertanian", gaji: 30000, posisi: "petani"},
	{keyword: "traktor", industri: "Pertanian", gaji: 6000000, posisi: "Operator Traktor"},
	{keyword: "traktor", industri: "Pertanian", gaji: 4000000, posisi: "Mekanik Alat Pertanian"},
	{keyword: "traktor", industri: "Pertanian", gaji: 5500000, posisi: "instruktur operasi"},
	{keyword: "laut", industri: "perikanan", gaji: 40000, posisi: "nelayan"},
	{keyword: "kandang", industri: "peternakan", gaji: 50000, posisi: "peternak"},
	{keyword: "pemrograman", industri: "TI", gaji: 25000000, posisi: "Programmer"},
	{keyword: "jaringan komputer", industri: "TI", gaji: 15000000, posisi: "Network Engineer"},
}

func main() {
	var pilih int
	var job kerja
	var cv resume
	var jumlah int

	for {
		menuUtama()
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			tambahResume(&cv)
		case 2:
			ubahResume(&cv)
			fmt.Scan(&pilih)
			switch pilih {
			case 1:
				ubahPengalamanKerjaBaru(&cv)
			case 2:
				ubahKeterampilanBaru(&cv)
			case 3:
				ubahPendidikanBaru(&cv)
			case 4:
				ubahDeskripsiBaru(&cv)
			default:
				return
			}
		case 3:
			hapusResume(&cv)
		case 4:
			tampilkanResume(cv)
		case 5:
			cetak(job, cariKecocokan(&cv, &job))
		case 6:
			jumlah = cariKecocokan(&cv, &job)
			urutanGaji(&job, jumlah)
			cetakGaji(job, jumlah)
		case 0:
			return
		}
	}
}

func menuUtama() {
	fmt.Println("		Menu Utama			")
	fmt.Println("1. Tambah Resume")
	fmt.Println("2. Ubah Resume")
	fmt.Println("3. Hapus Resume")
	fmt.Println("4. Tampilkan Resume")
	fmt.Println("5. Tampilkan Pekerjaan yang sesuai")
	fmt.Println("6. Tampilkan Urutan Gaji")
	fmt.Println("0. Back Menu Utama")
	fmt.Print("Pilih 0-5:")
}
func scanWord(str *string) {
	var temp string
	for {
		fmt.Scan(&temp)
		*str = *str + temp + " "
		if temp[len(temp)-1] == '.' {
			break
		}
	}
}
func tambahResume(cv *resume) {
	fmt.Println("		Masukkan Resume Anda		")
	fmt.Print("Pengalaman Kerja:")
	scanWord(&cv.pengalamanKerja)
	fmt.Print("Keterampilan:")
	scanWord(&cv.keterampilan)
	fmt.Print("Pendidikan:")
	scanWord(&cv.pendidikan)
	fmt.Print("Deskripsi:")
	scanWord(&cv.deskripsi)
}

func ubahResume(cv *resume) {
	fmt.Println("			Ubah Resume Anda			")
	fmt.Println("1. Masukkan Pengalaman kerja yang baru:")
	fmt.Println("2. Masukkan Keterampilan yang baru:")
	fmt.Println("3. Masukkan Pendidikan yang baru:")
	fmt.Println("4. Masukkan Deskripsi yang baru:")
	fmt.Print("Pilih antara (1/2/3/4):")
}

func ubahPengalamanKerjaBaru(cv *resume) {
	fmt.Print("Pengalaman Kerja Baru:")
	cv.pengalamanKerja = " "
	scanWord(&cv.pengalamanKerja)
}
func ubahKeterampilanBaru(cv *resume) {
	fmt.Print("Keterampilan Baru:")
	cv.keterampilan = " "
	scanWord(&cv.keterampilan)
}
func ubahPendidikanBaru(cv *resume) {
	fmt.Print("Pendidikan Baru:")
	cv.pendidikan = " "
	scanWord(&cv.pendidikan)
}
func ubahDeskripsiBaru(cv *resume) {
	fmt.Print("Deskripsi Baru:")
	cv.deskripsi = " "
	scanWord(&cv.deskripsi)
}

func hapusResume(cv *resume) {
	*cv = resume{}
}

func tampilkanResume(cv resume) {
	fmt.Println("============================================")
	fmt.Println("		Resume Anda			")
	fmt.Println("Pengalaman Kerja Anda:", cv.pengalamanKerja)
	fmt.Println("Keterampilan Anda:", cv.keterampilan)
	fmt.Println("Pendidikan Anda:", cv.pendidikan)
	fmt.Println("Deskripsi Anda:", cv.deskripsi)
	fmt.Println("=============================================")
}

func cariKecocokan(cv *resume, job *kerja) int {
	var i, j int
	var text string
	var jumlahHasil int

	jumlahHasil = 0
	text = cv.pengalamanKerja + " " + cv.keterampilan + " " + cv.pendidikan + " " + cv.deskripsi
	for i = 0; i < len(kataKunci); i++ {
		if seqSearch(kataKunci[i], text) {
			for j = 0; j < len(daftarPekerjaan); j++ {
				if daftarPekerjaan[j].keyword == kataKunci[i] {
					if jumlahHasil < NMAX {
						job[jumlahHasil] = daftarPekerjaan[j]
						jumlahHasil++
					}
				}
			}
		}
	}
	return jumlahHasil
}

func seqSearch(kata string, text string) bool {
	var i, j int

	for i = 0; i <= len(text)-len(kata); i++ {
		j = 0
		for j < len(kata) && text[i+j] == kata[j] {
			j++
		}
		if j == len(kata) {
			return true
		}
	}
	return false
}

func cetak(job kerja, jumlahHasil int) {
	var i int

	if jumlahHasil > 0 {
		fmt.Println("Pekerjaan yang sesuai dengan resume mu adalah:")
		for i = 0; i < jumlahHasil; i++ {
			fmt.Println("-", job[i].posisi, "Bidang", job[i].industri)
		}
	} else {
		fmt.Println("Tidak ada pekerjaan yang sesuai")
	}
}

func urutanGaji(A *kerja, jumlah int) {
	var i, pass int
	var temp pekerjaan

	for pass = 1; pass < jumlah; pass++ {
		temp = A[pass]
		i = pass

		for i > 0 && temp.gaji < A[i-1].gaji {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
	}

}

func cetakGaji(A kerja, jumlah int) {
	var i int

	for i = 0; i < jumlah; i++ {
		fmt.Printf("Rp %.0f - %s (%s)\n", A[i].gaji, A[i].posisi, A[i].industri)
	}
}
