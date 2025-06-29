package controller

// Fungsi untuk menghitung item total
func CalculateItemTotal(harga float64, quantity uint, diskonBarang float64) float64 {
	// Jika ada diskon per item
	if diskonBarang > 0 {
		// Jika diskon adalah persen, kurangi harga sesuai diskon
		return (harga * float64(quantity)) * (1 - diskonBarang/100)
	}
	// Jika tidak ada diskon, hanya menghitung harga * quantity
	return harga * float64(quantity)
}

// CalculateSubtotal calculates the subtotal for the invoice from the items
func CalculateSubtotal(items []struct {
	Kode         string  `json:"kode_barang"`
	Jumlah       uint    `json:"quantity"`
	Harga        float64 `json:"harga"`
	DiskonBarang float64 `json:"diskon_barang"`
}) float64 {
	var subtotal float64
	for _, item := range items {
		itemTotal := float64(item.Jumlah) * item.Harga
		if item.DiskonBarang > 0 {
			itemTotal -= (itemTotal * item.DiskonBarang / 100) // Apply discount if applicable
		}
		subtotal += itemTotal
	}
	return subtotal
}

// CalculateTotal calculates the total for the invoice after applying discount, tax, and shipping cost
func CalculateTotal(subtotal, ppn, diskonTotal, biayaPengiriman float64) float64 {
	// Calculate the total including PPN
	totalWithPPN := subtotal + (subtotal * (ppn / 100)) // PPN is assumed to be a percentage

	// Calculate the total after applying discount
	totalAfterDiscount := totalWithPPN - diskonTotal

	// Add the shipping cost
	total := totalAfterDiscount + biayaPengiriman

	return total
}
