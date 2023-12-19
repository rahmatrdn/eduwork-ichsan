package main

import "fmt"

type Mahasiswa struct {
	nim     uint
	nama    string
	jurusan string
}

var mahasiswa []Mahasiswa
var option uint

func TambahMahasiswa() {
	var request Mahasiswa

	fmt.Print("Input NIM: ")
	fmt.Scan(&request.nim)

	fmt.Print("Input Nama: ")
	fmt.Scan(&request.nama)

	fmt.Print("Input Jurusan: ")
	fmt.Scan(&request.jurusan)

	mahasiswa = append(mahasiswa, request)
	fmt.Println("created data: ", mahasiswa)
}

func HapusMahasiswa() {
	var request uint
	var index int

	fmt.Print("Input NIM to delete: ")
	fmt.Scan(&request)

	for i := 0; i < len(mahasiswa); i++ {
		el := mahasiswa[i]
		if el.nim == request {
			index = i
		}
	}

	if index != -1 {
		mahasiswa = append(mahasiswa[:index], mahasiswa[index+1:]...)
	}
	fmt.Println("success delete")
}

func TampilkanData() {
	for i := 0; i < len(mahasiswa); i++ {
		number := i + 1
		el := mahasiswa[i]
		fmt.Println("mahasiswa", number, ":", el)
	}
}

func main() {
	for {
		fmt.Println("select the command to use")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Hapus Mahasiswa")
		fmt.Println("3. Tampilkan Mahasiswa")
		fmt.Println("4. Exit")
		fmt.Print("select option: ")
		fmt.Scan(&option)

		if option == 1 {
			TambahMahasiswa()
		} else if option == 2 {
			HapusMahasiswa()
		} else if option == 3 {
			TampilkanData()
		} else if option == 4 {
			fmt.Println("--exit program--")
			break
		}
	}
}
