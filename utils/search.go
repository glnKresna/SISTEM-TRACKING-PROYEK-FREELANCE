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

// Binary search untuk mencari proyek berdasarkan judul
func BinSearch(proyekList []models.Proyek, judul string) (models.Proyek, bool) {
	judul = strings.ToLower(judul)
	kiri := 0
	kanan := len(proyekList) - 1

	for kiri <= kanan {
		mid := (kiri + kanan) / 2
		midTitle := strings.ToLower(proyekList[mid].Judul)
		if midTitle == judul {
			return proyekList[mid], true
		}
		if midTitle < judul {
			kiri = mid + 1
		} else {
			kanan = mid - 1
		}
	}
	return models.Proyek{}, false
}
