package internal

import "errors"

type InvoiceData struct {
	Title       string
	Quantity    float32
	Price       float32
	TotalAmount float32
}

func (d *InvoiceData) CalculateTotalAmount() float32 {
	totalAmount := d.Quantity * d.Price
	return totalAmount
}

func NewInvoiceData(title string, qty float32, price interface{}) (*InvoiceData, error) {
	var convertedPrice float32

	switch priceValue := price.(type) {
	case int:
		convertedPrice = float32(priceValue)
	case float32:
		convertedPrice = priceValue
	case float64:
		convertedPrice = float32(priceValue)
	default:
		return nil, errors.New("type not accepted")
	}

	return &InvoiceData{
		Title:    title,
		Quantity: qty,
		Price:    convertedPrice,
	}, nil
}
