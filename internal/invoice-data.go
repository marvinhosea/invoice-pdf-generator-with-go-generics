package internal

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

func NewInvoiceData(title string, qty float32, price interface{}) *InvoiceData {
	var convertedPrice float32

	switch priceValue := price.(type) {
	case int:
		convertedPrice = qty * float32(priceValue)
	case float32:
		convertedPrice = qty * priceValue
	case float64:
		convertedPrice = qty * float32(priceValue)
	default:
		panic("type not accepted")
	}

	return &InvoiceData{
		Title:    title,
		Quantity: qty,
		Price:    convertedPrice,
	}
}
