package main

import (
	"fmt"
	internal "marvinhosea/invoices/internal"
)

func main() {
	// Generate sample invoice data
	ecommerceInvoiceData := internal.NewInvoiceData("Ecommerce application", 1, 3000)
	laptopInvoiceData := internal.NewInvoiceData("Macbook Pro", 2, 200)
	// Invoice Items collection
	invoiceItems := []*internal.InvoiceData{ecommerceInvoiceData, laptopInvoiceData}

	// Create single invoice
	invoice := internal.CreateInvoice("Example Shop", "Example address", invoiceItems)
	fmt.Printf("The Total Invoice Amount is: %f", invoice.CalculateInvoiceTotalAmount())

}
