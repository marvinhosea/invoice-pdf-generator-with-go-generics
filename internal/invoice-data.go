package internal

type Amount interface {
	~int | ~float32 | ~float64
}

type InvoiceData[T Amount] struct {
	Title       string
	Quantity    T
	Price       T
	TotalAmount T
}

func (d *InvoiceData[T]) CalculateTotalAmount() T {
	totalAmount := d.Quantity * d.Price
	return totalAmount
}

func CreateInvoiceItems[T Amount](invoiceData ...*InvoiceData[T]) []*InvoiceData[T] {
	var items []*InvoiceData[T]
	for _, datum := range invoiceData {
		items = append(items, datum)
	}
	return items
}

func NewInvoiceData[T Amount](title string, qty, price T) *InvoiceData[T] {
	return &InvoiceData[T]{
		Title:    title,
		Quantity: qty,
		Price:    price,
	}
}
