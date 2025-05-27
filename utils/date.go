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

// InputDate prompts for and validates a date input
func InputDate(reader *bufio.Reader, prompt string) string {
	for {
		fmt.Print(prompt)
		tanggal, _ := reader.ReadString('\n')
		tanggal = strings.TrimSpace(tanggal)

		formattedDate := FormatDate(tanggal)
		if formattedDate != "" {
			return formattedDate
		}
		fmt.Println("Format tanggal tidak valid! Gunakan format DD MM YYYY")
	}
}
