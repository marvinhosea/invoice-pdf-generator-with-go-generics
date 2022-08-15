package internal

type Invoice struct {
	Name         string
	Address      string
	InvoiceItems []*InvoiceData
}

func CreateInvoice(name string, address string, invoiceItems []*InvoiceData) *Invoice {
	return &Invoice{
		Name:         name,
		Address:      address,
		InvoiceItems: invoiceItems,
	}
}

func (i *Invoice) CalculateInvoiceTotalAmount() float32 {
	var invoiceTotalAmount float32 = 0
	for _, data := range i.InvoiceItems {
		amount := data.CalculateTotalAmount()
		invoiceTotalAmount += amount
	}

	return invoiceTotalAmount
}
