package utils

import (
	"freelancing/models"
	"strings"
)

// Urutkan berdasarkan judul
func BubbleSortByJudul(proyekList *[]models.Proyek) {
	n := len(*proyekList)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if strings.ToLower((*proyekList)[j].Judul) > strings.ToLower((*proyekList)[j+1].Judul) {
				(*proyekList)[j], (*proyekList)[j+1] = (*proyekList)[j+1], (*proyekList)[j]
			}
		}
	}
}

// Urutkan berdasarkan ID
func InsertionSortByID(proyekList *[]models.Proyek) {
	for i := 1; i < len(*proyekList); i++ {
		key := (*proyekList)[i]
		j := i - 1
		for j >= 0 && (*proyekList)[j].ID > key.ID {
			(*proyekList)[j+1] = (*proyekList)[j]
			j--
		}
		(*proyekList)[j+1] = key
	}
}

// Urutkan berdasarkan status (Pending > Ongoing > Selesai)
func urutanStatus(status string) int {
	switch strings.ToLower(status) {
	case "pending":
		return 1
	case "ongoing":
		return 2
	case "selesai":
		return 3
	default:
		return 4
	}
}

// Mengurutkan berdasarkan status dengan selection sort
func SelectionSortByStatus(proyekList *[]models.Proyek) {
	n := len(*proyekList)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if urutanStatus((*proyekList)[j].Status) < urutanStatus((*proyekList)[minIdx].Status) {
				minIdx = j
			}
		}
		(*proyekList)[i], (*proyekList)[minIdx] = (*proyekList)[minIdx], (*proyekList)[i]
	}
}
