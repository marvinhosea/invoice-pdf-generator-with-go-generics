package main

import (
	internal "marvinhosea/invoices/internal"
)

func main() {
	// Generate sample invoice data
	ecommerceInvoiceData := internal.NewInvoiceData("Ecommerce application", 1, 3000)
	laptopInvoiceData := internal.NewInvoiceData("Macbook Pro", 2, 1999)
	// Invoice Items collection
	invoiceItems := internal.CreateInvoiceItems(ecommerceInvoiceData, laptopInvoiceData)

	// Create single invoice
	invoice := internal.CreateInvoice("Example Shop", "Example address", invoiceItems)
	err := internal.GenerateInvoicePdf(*invoice)
	if err != nil {
		panic(err)
	}
}
