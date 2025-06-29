package generator

import "time"

// Fungsi untuk memformat tanggal ke format yang diinginkan "13-Jun-2025"
func FormatTanggalBeli(tanggal time.Time) string {
	// Format menjadi "13-Jun-2025"
	return tanggal.Format("02-Jan-2006")
}
