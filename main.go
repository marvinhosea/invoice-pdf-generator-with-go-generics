package main

import (
	"fmt"
	"marvinhosea/invoices/internal"
)

func main() {
	// Generate sample invoice data
	ecommerceInvoiceData := internal.NewInvoiceData("Ecommerce application", 1, 3000.50)
	laptopInvoiceData := internal.NewInvoiceData("Macbook Pro", 2, 1999.70)
	// Invoice Items collection
	invoiceItems := []*internal.InvoiceData{ecommerceInvoiceData, laptopInvoiceData}

	// Create single invoice
	invoice := internal.CreateInvoice("Example Shop", "Example address", invoiceItems)
	fmt.Printf("The Total Invoice Amount is: %d", invoice.CalculateInvoiceTotalAmount())
}
