package main

import (
	"h8-first-assignment/helpers"
	"os"
)

func main() {
	nama := os.Args
	peserta := helpers.Peserta{Nama: nama[1]}

	peserta.FindPeserta()
}
