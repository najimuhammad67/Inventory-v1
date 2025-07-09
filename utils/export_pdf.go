package utils

import (
	"github.com/jung-kurt/gofpdf"
	"inventory-app/models"
	"net/http"
	"strconv"
)

func ExportPDF(w http.ResponseWriter, items []models.Item) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Laporan Inventaris")
	pdf.Ln(12)

	// Header tabel
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(10, 10, "No", "1", 0, "C", false, 0, "")
	pdf.CellFormat(60, 10, "Nama Barang", "1", 0, "C", false, 0, "")
	pdf.CellFormat(50, 10, "Kategori", "1", 0, "C", false, 0, "")
	pdf.CellFormat(20, 10, "Jumlah", "1", 1, "C", false, 0, "")

	// Isi tabel
	pdf.SetFont("Arial", "", 12)
	for i, item := range items {
		pdf.CellFormat(10, 10, strconv.Itoa(i+1), "1", 0, "C", false, 0, "")
		pdf.CellFormat(60, 10, item.Name, "1", 0, "", false, 0, "")
		pdf.CellFormat(50, 10, item.Category, "1", 0, "", false, 0, "")
		pdf.CellFormat(20, 10, strconv.Itoa(item.Quantity), "1", 1, "C", false, 0, "")
	}

	err := pdf.Output(w)
	if err != nil {
		http.Error(w, "Gagal membuat PDF", http.StatusInternalServerError)
	}
}
