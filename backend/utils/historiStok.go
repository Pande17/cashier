package utils

import (
	"cashier-machine/model"
	repository "cashier-machine/repository/config"
	"cashier-machine/repository/modelfunc"

	"github.com/siruspen/logrus"
)

// Function to create a new history record for an item
func CreateHistoriBarang(p *model.Details, keterangan string, amount int, status string) (model.Histori, error) {
	// Create a new history record
	histori := model.Histori{
		ID_barang:  uint(p.ID), // Set ID_barang from the parameter p
		Amount:     amount,     // Set amount from the parameter
		Status:     status,     // Set status from the parameter
		Keterangan: keterangan, // Set description from the parameter
	}

	// Save the history record to the database
	err := repository.Mysql.DB.Create(&histori).Error // Save history to the database
	if err != nil {
		return model.Histori{}, err // Return error if saving fails
	}

	return histori, nil // Return the newly created history record
}

// Function to create a new sales history record
func CreateHistoriInvoice(p *model.CreateP, keterangan string, amount int, status string) (model.Histori, error) {
	histori := modelfunc.Histori{
		Histori: model.Histori{
			ID_barang:  uint(p.ID), // Set ID_barang from the parameter p
			Amount:     amount,     // Set amount from the parameter
			Status:     status,     // Set status from the parameter
			Keterangan: keterangan, // Set description from the parameter
		},
	}

	err := histori.Create(repository.Mysql.DB) // Save the sales history record to the database
	if err != nil {
		return model.Histori{}, err // Return error if saving fails
	}

	return histori.Histori, nil // Return the newly created history record
}

// Function to get history records by item ID
func GetASKMByIDBarang(idb uint64) ([]model.HistoriASKM, error) {
	histori := modelfunc.Histori{
		Histori: model.Histori{
			ID_barang: uint(idb), // Set ID_barang from the parameter idb
		},
	}

	logrus.Println(idb) // Log ID_barang for debugging

	newHistory, err := histori.GetIDBarang(repository.Mysql.DB) // Retrieve history records by ID_barang
	if err != nil {
		return nil, err // Return error if retrieval fails
	}

	var haskm []model.HistoriASKM  // Initialize a slice for the retrieved history records
	for _, h := range newHistory { // Iterate through the retrieved history records
		haskm = append(haskm, model.HistoriASKM{ // Add each history record to the slice
			Amount:     h.Amount,     // Set Amount from the history record
			Status:     h.Status,     // Set Status from the history record
			Keterangan: h.Keterangan, // Set Description from the history record
			Model:      h.Model,      // Set Model from the history record
		})
	}

	return haskm, nil // Return the slice of retrieved history records
}

// Function to get history records by item ID
func GetASK(idb uint64) ([]model.HistoriASK, error) {
	histori := modelfunc.Histori{
		Histori: model.Histori{
			ID_barang: uint(idb), // Set ID_barang from the parameter idb
		},
	}

	newHistory, err := histori.GetIDBarang(repository.Mysql.DB) // Retrieve history records by ID_barang
	if err != nil {
		return nil, err // Return error if retrieval fails
	}

	var hask []model.HistoriASK    // Initialize a slice for the retrieved history records
	for _, h := range newHistory { // Iterate through the retrieved history records
		hask = append(hask, model.HistoriASK{ // Add each history record to the slice
			Amount:     h.Amount,     // Set Amount from the history record
			Status:     h.Status,     // Set Status from the history record
			Keterangan: h.Keterangan, // Set Description from the history record
		})
	}

	return hask, nil // Return the slice of retrieved history records
}
