package main

import (
	"bufio"
	"fmt"
	"freelancing/models"
	"freelancing/services"
	"os"
	"strings"
)

func main() {
	var proyekList []models.Proyek
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== SISTEM MANAJEMEN PROYEK FREELANCING ===")
		fmt.Println("1. Tambah Proyek")
		fmt.Println("2. Lihat Proyek")
		fmt.Println("3. Edit Proyek")
		fmt.Println("4. Hapus Proyek")
		fmt.Println("5. Cari Proyek")
		fmt.Println("6. Keluar")
		fmt.Print("\nPilih menu (1-6): ")

		pilihan, _ := reader.ReadString('\n')
		pilihan = strings.TrimSpace(pilihan)

		switch pilihan {
		case "1":
			services.TambahProyek(&proyekList)
		case "2":
			services.LihatProyek(proyekList)
		case "3":
			services.EditProyek(&proyekList)
		case "4":
			services.HapusProyek(&proyekList)
		case "5":
			services.CariProyek(proyekList)
		case "6":
			fmt.Println("Terima kasih telah menggunakan sistem manajemen proyek.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih menu 1-6.")
		}
	}
}

