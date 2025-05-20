package main

import (
	"fmt"
)

const NMAX int = 100

type resume struct {
	pengalamanKerja string
	keterampilan    string
	pendidikan      string
	deskripsi       string
}

type pekerjaan struct {
	industri string
	gaji     float64
	posisi   string
}
type kerja [NMAX]pekerjaan
// type exp [NMAX]expertise

func menuUtama() {
	fmt.Println("			Menu Resume				")
	fmt.Println("1. Tambah Resume")
	fmt.Println("2. Ubah Resume")
	fmt.Println("3. Hapus Resume")
	fmt.Println("0. Back Menu")
	fmt.Print("Pilih (1/2/3/0):")
} 
func tambahResume(cv *resume) {
	fmt.Println("		Masukan Resume Anda			")
	fmt.Print("Pengalaman Kerja:")
	fmt.Scanf("%[^\n]", &cv.pengalamanKerja)
	fmt.Scanln()
	fmt.Println("Keterampilan:")
	fmt.Scanf("%[^\n]", &cv.keterampilan)
	fmt.Scanln()
	fmt.Println("Pendidikan Terakhir:")
	fmt.Scanf("%[^\n]", &cv.pendidikan)
	fmt.Scanln()
	fmt.Println("Deskripsi diri anda:")
	fmt.Scanf("%[^\n]", &cv.deskripsi)
	fmt.Scanln()
}

func ubahResume(cv *resume) {
	fmt.Println("		Ubah Resume		")
	tambahResume(cv)
}

func hapusResume(cv *resume) {

}


func main() {
	var pilih int
	var cv resume
	var job kerja
	
	// var pekerjaan [5]string{"Pendidikan", "Pertanian", "Perikanan",  "Teknologi", "Seni"}
	// {tani, teknologi, belajar, mengajar}
	// industri[10] {pendidikan, pertanian}
	// var pekerjaan.industri [5]string {pendidikan, pertanian,perikanan, }
	job[0] = pekerjaan{"Pendidikan", 1000000, "Kepsek"}
	job[1] = pekerjaan{"Kesehatan", 1000000, "Dokter"}
	job[2] = pekerjaan{"Pertanian", 1000000, "Petani"}
	job[3] = pekerjaan{"Perikanan", 1000000, "Nelayan"}
	job[4] = pekerjaan{"Seni", 1000000, "Seniman"}
	job[5] = pekerjaan{"Teknologi", 1000000, "Back-End dev"}
	
	for {
		menuUtama()
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			tambahResume(&cv)
		case 2:
			ubahResume(&cv)
		case 3:
			hapusResume(&cv)
		case 0:
			return
		}
	}

}
