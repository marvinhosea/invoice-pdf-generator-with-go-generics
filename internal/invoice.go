package internal

type Invoice[T Amount] struct {
	Name         string
	Address      string
	InvoiceItems []*InvoiceData[T]
}

func CreateInvoice[T Amount](name string, address string, invoiceItems []*InvoiceData[T]) *Invoice[T] {
	return &Invoice[T]{
		Name:         name,
		Address:      address,
		InvoiceItems: invoiceItems,
	}
}

func (i *Invoice[T]) CalculateInvoiceTotalAmount() T {
	var invoiceTotalAmount T = 0
	for _, data := range i.InvoiceItems {
		amount := data.CalculateTotalAmount()
		invoiceTotalAmount += amount
	}

	return invoiceTotalAmount
}
