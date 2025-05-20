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
	industri     string
	gaji         float64
	posisi       string
	industriAsli string
}

// deklarasi global
var kataKata = []string{"sawah", "laut", "kandang", "komputer", "kelas", "singkong",
 "kedelai", "ubi"}

var daftarPekerjaan = []pekerjaan{
	{industri: "sawah", gaji: 3000, posisi: "Petani", industriAsli: "Pertanian"},
	{industri: "laut", gaji: 3500, posisi: "Nelayan", industriAsli: "Perikanan"},
	{industri: "kandang", gaji: 3200, posisi: "Peternak", industriAsli: "Perternakan"},
	{industri: "komputer", gaji: 7000, posisi: "Programmer", industriAsli: "Teknologi"},
	{industri: "kelas", gaji: 4000, posisi: "Guru", industriAsli: "Pertanian"},
	{industri: "singkong", gaji: 4000, posisi: "Farmer", industriAsli: "Pertanian"},
	{industri: "ubi", gaji: 4000, posisi: "Tukang kebun", industriAsli: "Perkebunan"},	
}
	
func main() {
	var pengalaman resume
	var hasilCocok [NMAX]pekerjaan
	var jumlahHasil int = 0

	pengalaman = resume{
		pengalamanKerja: "saya bekerja di kebun",
		keterampilan:    "menanam ubi",
		pendidikan:      "SMK Pertanian",
		deskripsi:       "gua pernah memenangkan lomba alat inovatif",
	}

	text := pengalaman.pengalamanKerja + " " + pengalaman.keterampilan + " " + pengalaman.pendidikan + " " + pengalaman.deskripsi

	for i := 0; i < len(kataKata); i++ {
		if seqSearch(kataKata[i], text) {
			for j := 0; j < len(daftarPekerjaan); j++ {
				if daftarPekerjaan[j].industri == kataKata[i] {
					if jumlahHasil < NMAX {
						hasilCocok[jumlahHasil] = daftarPekerjaan[j]
						jumlahHasil++
					}
					break
				}
			}
		}
	}

	if jumlahHasil > 0 {
		fmt.Println("Pekerjaan yang sesuai dengan resumemu adalah:")
		for i := 0; i < jumlahHasil; i++ {
			fmt.Println("-", hasilCocok[i].posisi)
		}
	} else {
		fmt.Println("Tidak ada pekerjaan yang sesuai")
	}
}

func seqSearch(kata string, text string) bool {
	for i := 0; i <= len(text)-len(kata); i++ {
		j := 0
		for j < len(kata) && text[i+j] == kata[j] {
			j++
		}
		if j == len(kata) {
			return true
		}
	}
	return false
}