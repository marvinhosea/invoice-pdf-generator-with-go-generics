package internal

type InvoiceData struct {
	Title       string
	Quantity    int
	Price       int
	TotalAmount int
}

func (d *InvoiceData) CalculateTotalAmount() int {
	totalAmount := d.Quantity * d.Price
	return totalAmount
}

func CreateInvoiceItems(invoiceData ...*InvoiceData) []*InvoiceData {
	var items []*InvoiceData
	for _, datum := range invoiceData {
		items = append(items, datum)
	}
	return items
}

func NewInvoiceData(title string, qty, price int) *InvoiceData {
	return &InvoiceData{
		Title:    title,
		Quantity: qty,
		Price:    price,
	}
}
