package main

import (
	"fmt"
	"os"
)

type Teman struct {
	Nama     string
	Alamat   string
	Pekerjaan string
	Alasan   string
}

var dataTeman = map[int]Teman{
	1: {"Jasmine", "Jakarta", "Pramugari", "Ingin belajar pemrograman lebih lanjut"},
	2: {"Belle", "Prancis", "Bioinformasi", "Mengembangkan keterampilan pemrograman back-end"},
	3: {"Mulan", "Hangzhou", "Developer", "Tertarik dengan bahasa Go"},
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run biodata.go [nomor_absen]")
		os.Exit(1)
	}

	nomorAbsen := args[1]

	display(nomorAbsen)
}

func display(nomorAbsen string) {
	var nomor int
	_, err := fmt.Sscanf(nomorAbsen, "%d", &nomor)
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka")
		os.Exit(1)
	}

	teman, found := dataTeman[nomor]
	if !found {
		fmt.Println("Teman dengan nomor absen", nomor, "tidak ditemukan")
		os.Exit(1)
	}

	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
}