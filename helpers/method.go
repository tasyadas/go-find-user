package helpers

import (
	"fmt"
	"strings"
)

func (p Peserta) FindPeserta() {
	var pesertas = dataPeserta()

	for key := range pesertas {
		if strings.ToLower(pesertas[key].Nama) == strings.ToLower(p.Nama) {
			fmt.Println("ID:", key)
			fmt.Println("Nama:", pesertas[key].Nama)
			fmt.Println("Alamat:", pesertas[key].Alamat)
			fmt.Println("Pekerjaan:", pesertas[key].Pekerjaan)
			fmt.Println("Alasan:", pesertas[key].Alasan)
		}
	}

}
