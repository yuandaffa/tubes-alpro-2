package main
import "fmt"

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
	{keyword: "traktor", industri: "Pertanian", gaji: 4000000, posisi: "Mekanik Alat Pertanian"}
	{keyword: "traktor", industri: "Pertanian", gaji: 5500000, posisi: "instruktur operasi"},
	{keyword: "laut", industri: "perikanan", gaji: 40000, posisi: "nelayan"},
	{keyword: "kandang", industri: "peternakan", gaji: 50000, posisi: "peternak"},
	{keyword: "pemrograman", industri: "TI", gaji: 25000000, posisi: "Programmer"},
	{keyword: "jaringan komputer", industri: "TI", gaji: 15000000, posisi: "Network Engineer"},
}

func main (){
	var cv resume
	var pilih int

	for {
		menuUtama()
		fmt.Scan(&pilih)
		switch pilih{
		case 1:
			tambahResume(&cv)
		case 2:
			ubahResume()
			fmt.Scan(&pilih)
			switch pilih{
			case 1:
				ubahPengalaman(&cv)
			case 2:
				ubahKeahlian(&cv)
			case 3:
				ubahPendidikan(&cv)
			case 4:
				ubahDeskripsi(&cv)
			case 0:
				return
		}
		case 3:
			hapusResume(&cv)
		case 4:
			tampilkanResume(cv)
		case 0:
			return
		}
	}
}

func menuUtama(){
	fmt.Println("		Menu Utama			")
	fmt.Println("1. Tambah Resume")
	fmt.Println("2. Ubah Resume")
	fmt.Println("3. Hapus Resume")
	fmt.Println("4. Tampilkan Resume")
	fmt.Println("5. Tampilkan Pekerjaan yang sesuai")
	fmt.Println("0. Back Menu Utama")
	fmt.Print("Pilih 0-5:")
}

func tambahResume(cv *resume){
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

func ubahResume(){
	fmt.Println(" Ubah Resume ")
	fmt.Println("1. Masukkan Pengalaman kerja yang baru:")
	fmt.Println("2. Masukkan Keahlian yang baru:")
	fmt.Println("3. Masukkan Pendidikan yang baru:")
	fmt.Println("4. Masukkan deskripsi yang baru:")
	fmt.Println("0. back menu utama:")
	fmt.Println("pilih 1-5:")
}

func ubahPengalaman(cv *resume){
	fmt.Print("Pengalaman Kerja:")
	cv.pengalamanKerja = " "
 	scanWord(&cv.pengalamanKerja)
}

func ubahKeahlian(cv *resume){
	fmt.Print("Keterampilan:")
	cv.keterampilan = " "
	scanWord(&cv.keterampilan)
}

func ubahPendidikan(cv *resume){
	fmt.Print("Pendidikan:")
	cv.pendidikan = " "
	scanWord(&cv.pendidikan)
}

func ubahDeskripsi(cv *resume){
	fmt.Print("Deskripsi:")
	cv.deskripsi = " "
	scanWord(&cv.deskripsi)
}

func hapusResume(cv *resume){
	*cv = resume{}
}

func tampilkanResume(cv resume){
	fmt.Println("============================================")
	fmt.Println("		Resume Anda			")
	fmt.Println("Pengalaman Kerja Anda:", cv.pengalamanKerja)
	fmt.Println("Keterampilan Anda:", cv.keterampilan)
	fmt.Println("Pendidikan Anda:", cv.pendidikan)
	fmt.Println("Deskripsi Anda:", cv.deskripsi)
	fmt.Println("=============================================")
}

func urutanGaji(A *kerja){
	var i, idx, pass int
	var temp float64

	pass = 1
	for pass <= len(daftarPekerjaanpe){
		i = pass
		temp = A[pass]
		for i > 0 && temp < A[i].gaji{
			A[i] = A[i].gaji
			i++
		}
		A[i] = temp
		pass++
	}
}

func cetakGaji(A *kerja){
	var i int

	for i = 0 < len(daftarPekerjaan){
		fmt.Println(A[i].gaji)
	}
}