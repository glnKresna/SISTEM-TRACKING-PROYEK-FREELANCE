package services

import (
	"bufio"
	"fmt"
	"freelancing/models"
	"freelancing/utils"
	"os"
	"strings"
)

// Fungsi tambah proyek
func TambahProyek(proyekList *[]models.Proyek) {
	var judul, klien, status, catatan string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Judul Proyek: ")
	judul, _ = reader.ReadString('\n')
	judul = strings.TrimSpace(judul)

	fmt.Print("Nama Klien: ")
	klien, _ = reader.ReadString('\n')
	klien = strings.TrimSpace(klien)

	tanggalTerima := utils.InputTgl(reader, "Tanggal Terima (DD MM YYYY): ")
	deadline := utils.InputTgl(reader, "Deadline (DD MM YYYY): ")

	fmt.Print("Status Pengerjaan (1: Pending, 2: Ongoing, 3: Selesai): ")
	status, _ = reader.ReadString('\n')
	status = strings.TrimSpace(status)

	switch status {
	case "1":
		status = "Pending"
	case "2":
		status = "Ongoing"
	case "3":
		status = "Selesai"
	default:
		status = "Pending"
	}

	fmt.Print("Tambah Catatan: ")
	catatan, _ = reader.ReadString('\n')
	catatan = strings.TrimSpace(catatan)

	proyekBaru := models.Proyek{
		ID:            utils.GenerateID(len(*proyekList)),
		Judul:         judul,
		Klien:         klien,
		Status:        status,
		TanggalTerima: tanggalTerima,
		Deadline:      deadline,
		Catatan:       catatan,
	}

	*proyekList = append(*proyekList, proyekBaru)
	fmt.Println("Proyek baru berhasil ditambahkan.")
}

// Fungsi menampilkan tabel daftar proyek
func LihatProyek(proyekList []models.Proyek) {
	if len(proyekList) == 0 {
		fmt.Println("Belum ada proyek, tambahkan untuk melihat.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nPilih urutan tampilan:")
	fmt.Println("1. Berdasarkan Judul Proyek")
	fmt.Println("2. Berdasarkan ID Proyek")
	fmt.Println("3. Berdasarkan Status")
	fmt.Print("Pilihan (1-3): ")

	pilihan, _ := reader.ReadString('\n')
	pilihan = strings.TrimSpace(pilihan)

	proyekTerurut := make([]models.Proyek, len(proyekList))
	copy(proyekTerurut, proyekList)

	switch pilihan {
	case "1":
		utils.BubbleSortByJudul(&proyekTerurut)
		fmt.Println("\nDaftar Proyek (Diurutkan berdasarkan Judul)")
	case "2":
		utils.InsertionSortByID(&proyekTerurut)
		fmt.Println("\nDaftar Proyek (Diurutkan berdasarkan ID)")
	case "3":
		utils.SelectionSortByStatus(&proyekTerurut)
		fmt.Println("\nDaftar Proyek (Diurutkan berdasarkan status pengerjaan)")
	default:
		fmt.Println("Pilihan tidak valid! Menampilkan data tanpa pengurutan.")
		fmt.Println("\nDaftar Proyek")
		proyekTerurut = proyekList
	}

	fmt.Println("|======================================================================================================================|")
	fmt.Println("|   ID   |   Judul Proyek   |     Klien     |     Status     |   Tanggal Terima   |    Deadline    |      Catatan      |")
	fmt.Println("|======================================================================================================================|")
	for _, p := range proyekTerurut {
		fmt.Printf("| %-6s | %-16s | %-13s | %-14s | %-18s | %-14s | %-17s |\n",
			p.ID,
			utils.TruncateStr(p.Judul, 16),
			utils.TruncateStr(p.Klien, 13),
			utils.TruncateStr(p.Status, 14),
			utils.TruncateStr(p.TanggalTerima, 18),
			utils.TruncateStr(p.Deadline, 14),
			utils.TruncateStr(p.Catatan, 17))
	}
	fmt.Println("|======================================================================================================================|")
}

// Fungsi edit proyek
func EditProyek(proyekList *[]models.Proyek) {
	if len(*proyekList) == 0 {
		fmt.Println("Belum ada proyek yang bisa diedit.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var searchInput string

	fmt.Print("Masukkan ID proyek yang ingin diedit: ")
	searchInput, _ = reader.ReadString('\n')
	searchInput = strings.TrimSpace(searchInput)

	proyek, found := utils.SeqSearch(*proyekList, searchInput)
	if !found {
		fmt.Println("Proyek tidak ditemukan.")
		return
	}

	targetIndex := -1
	for i, p := range *proyekList {
		if p.ID == proyek.ID {
			targetIndex = i
			break
		}
	}

	target := &(*proyekList)[targetIndex]
	fmt.Printf("\nData proyek yang akan diedit:")
	fmt.Printf("\nID: %s", target.ID)
	fmt.Printf("\nJudul: %s", target.Judul)
	fmt.Printf("\nKlien: %s", target.Klien)
	fmt.Printf("\nStatus: %s", target.Status)
	fmt.Printf("\nTanggal Terima: %s", target.TanggalTerima)
	fmt.Printf("\nDeadline: %s", target.Deadline)
	fmt.Printf("\nCatatan: %s\n", target.Catatan)

	fmt.Printf("\nJudul Baru (kosongkan jika tidak ingin mengubah): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		target.Judul = input
	}

	fmt.Printf("Klien Baru (kosongkan jika tidak ingin mengubah): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		target.Klien = input
	}

	fmt.Printf("Status Baru (1: Pending, 2: Ongoing, 3: Selesai) (kosongkan jika tidak ingin mengubah): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		switch input {
		case "1":
			target.Status = "Pending"
		case "2":
			target.Status = "Ongoing"
		case "3":
			target.Status = "Selesai"
		}
	}

	fmt.Printf("Tanggal Terima Baru (kosongkan jika tidak ingin mengubah): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		if tanggalTerformat := utils.FormatDate(input); tanggalTerformat != "" {
			target.TanggalTerima = tanggalTerformat
		} else {
			fmt.Println("Format tanggal tidak valid! Tanggal tidak diubah.")
		}
	}

	fmt.Printf("Deadline Baru (kosongkan jika tidak ingin mengubah): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		if tanggalTerformat := utils.FormatDate(input); tanggalTerformat != "" {
			target.Deadline = tanggalTerformat
		} else {
			fmt.Println("Format tanggal tidak valid! Deadline tidak diubah.")
		}
	}

	fmt.Printf("Catatan Baru (kosongkan jika tidak ingin mengubah): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		target.Catatan = input
	}

	fmt.Println("Proyek berhasil diperbarui.")
}

// Fungsi hapus proyek
func HapusProyek(proyekList *[]models.Proyek) {
	if len(*proyekList) == 0 {
		fmt.Println("Belum ada proyek yang bisa dihapus.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var searchInput string

	fmt.Print("Masukkan ID proyek yang ingin dihapus: ")
	searchInput, _ = reader.ReadString('\n')
	searchInput = strings.TrimSpace(searchInput)

	proyek, found := utils.SeqSearch(*proyekList, searchInput)
	if !found {
		fmt.Println("Proyek tidak ditemukan.")
		return
	}

	targetIndex := -1
	for i, p := range *proyekList {
		if p.ID == proyek.ID {
			targetIndex = i
			break
		}
	}

	fmt.Printf("\nData proyek yang akan dihapus:")
	fmt.Printf("\nID: %s", proyek.ID)
	fmt.Printf("\nJudul: %s", proyek.Judul)
	fmt.Printf("\nKlien: %s", proyek.Klien)
	fmt.Printf("\nStatus: %s", proyek.Status)
	fmt.Printf("\nTanggal Terima: %s", proyek.TanggalTerima)
	fmt.Printf("\nDeadline: %s", proyek.Deadline)
	fmt.Printf("\nCatatan: %s\n", proyek.Catatan)

	fmt.Print("\nApakah Anda yakin ingin menghapus proyek ini? (y/n): ")
	konfirmasi, _ := reader.ReadString('\n')
	konfirmasi = strings.TrimSpace(strings.ToLower(konfirmasi))

	if konfirmasi == "y" {
		(*proyekList)[targetIndex] = (*proyekList)[len(*proyekList)-1]
		*proyekList = (*proyekList)[:len(*proyekList)-1]
		fmt.Println("Proyek berhasil dihapus.")
	} else {
		fmt.Println("Penghapusan proyek dibatalkan.")
	}
}

// Fungsi cari proyek
func CariProyek(proyekList []models.Proyek) {
	if len(proyekList) == 0 {
		fmt.Println("Belum ada proyek yang bisa dicari.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nPilih metode pencarian:")
	fmt.Print("\n1. Cari berdasarkan ID")
	fmt.Print("\n2. Cari berdasarkan nama proyek")
	fmt.Print("\nPilihan (1/2): ")

	pilihan, _ := reader.ReadString('\n')
	pilihan = strings.TrimSpace(pilihan)

	var hasil models.Proyek
	var found bool
	var searchInput string

	switch pilihan {
	case "1":
		fmt.Print("Masukkan ID proyek yang dicari: ")
		searchInput, _ = reader.ReadString('\n')
		searchInput = strings.TrimSpace(searchInput)
		hasil, found = utils.SeqSearch(proyekList, searchInput)
	case "2":
		fmt.Print("Masukkan nama proyek yang dicari: ")
		searchInput, _ = reader.ReadString('\n')
		searchInput = strings.TrimSpace(searchInput)
		proyekTerurut := make([]models.Proyek, len(proyekList))
		copy(proyekTerurut, proyekList)
		utils.BubbleSortByJudul(&proyekTerurut)
		hasil, found = utils.BinSearch(proyekTerurut, searchInput)
	default:
		fmt.Println("Pilihan tidak valid!")
		return
	}

	if found {
		fmt.Println("\nProyek ditemukan:")
		fmt.Printf("ID: %s\n", hasil.ID)
		fmt.Printf("Judul: %s\n", hasil.Judul)
		fmt.Printf("Klien: %s\n", hasil.Klien)
		fmt.Printf("Status: %s\n", hasil.Status)
		fmt.Printf("Tanggal Terima: %s\n", hasil.TanggalTerima)
		fmt.Printf("Deadline: %s\n", hasil.Deadline)
		fmt.Printf("Catatan: %s\n", hasil.Catatan)
	} else {
		fmt.Println("Proyek tidak ditemukan.")
	}
}
