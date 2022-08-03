package internal

type InvoiceData struct {
	Title       string
	Quantity    int
	Price       float32
	TotalAmount float32
}

func (d *InvoiceData) CalculateTotalAmount() float32 {
	totalAmount := float32(d.Quantity) * d.Price
	d.TotalAmount = totalAmount
	return totalAmount
}

func NewInvoiceData(tittle string, qty int, price interface{}) *InvoiceData {
	var convertedPrice float32

	switch priceValue := price.(type) {
	case int:
		convertedPrice = float32(qty * priceValue)
	case float32:
		convertedPrice = float32(qty) * priceValue
	case float64:
		convertedPrice = float32(qty) * float32(priceValue)
	default:
		panic("invalid data type")
	}

	return &InvoiceData{
		Title:    tittle,
		Quantity: qty,
		Price:    convertedPrice,
	}
}
