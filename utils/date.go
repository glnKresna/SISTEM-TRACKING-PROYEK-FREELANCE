package utils

import (
	"bufio"
	"fmt"
	"strings"
)

// Format penanggalan
func FormatDate(input string) string {
	input = strings.ReplaceAll(input, " ", "")

	if len(input) != 8 {
		return ""
	}

	for _, c := range input {
		if c < '0' || c > '9' {
			return ""
		}
	}

	hari := input[0:2]
	bulan := input[2:4]
	tahun := input[4:8]

	return hari + "/" + bulan + "/" + tahun
}

// Fungsi input & validasi tanggal
func InputTgl(reader *bufio.Reader, inputTanggal string) string {
	for {
		fmt.Print(inputTanggal)
		tanggal, _ := reader.ReadString('\n')
		tanggal = strings.TrimSpace(tanggal)

		tanggalTerformat := FormatDate(tanggal)
		if tanggalTerformat != "" {
			return tanggalTerformat
		}
		fmt.Println("Format tanggal tidak valid! Gunakan format DD MM YYYY")
	}
}
