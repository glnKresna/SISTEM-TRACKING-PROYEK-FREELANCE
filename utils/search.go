package utils

import (
	"freelancing/models"
	"strings"
)

// Sequential search untuk mencari proyek berdasarkan ID
func SeqSearch(proyekList []models.Proyek, id string) (models.Proyek, bool) {
	for _, proyek := range proyekList {
		if proyek.ID == id {
			return proyek, true
		}
	}
	return models.Proyek{}, false
}

// Binary search untuk mencari proyek berdasarkan nama proyek
func BinSearch(proyekList []models.Proyek, nama string) (models.Proyek, bool) {
	nama = strings.ToLower(nama)
	left := 0
	right := len(proyekList) - 1

	for left <= right {
		mid := (left + right) / 2
		midTitle := strings.ToLower(proyekList[mid].Judul)
		if midTitle == nama {
			return proyekList[mid], true
		}
		if midTitle < nama {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return models.Proyek{}, false
}
