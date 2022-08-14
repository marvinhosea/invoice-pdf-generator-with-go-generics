package internal

import (
	"fmt"
	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/creator"
	"marvinhosea/invoices/internal/config"
)

type Client[T Amount] struct {
	creator *creator.Creator
}

type cellStyle struct {
	ColSpan         int
	HAlignment      creator.CellHorizontalAlignment
	BackgroundColor creator.Color
	BorderSide      creator.CellBorderSide
	BorderStyle     creator.CellBorderStyle
	BorderWidth     float64
	BorderColor     creator.Color
	Indent          float64
}

var cellStyles = map[string]cellStyle{
	"heading-left": {
		BackgroundColor: creator.ColorRGBFromHex("#332f3f"),
		HAlignment:      creator.CellHorizontalAlignmentLeft,
		BorderColor:     creator.ColorWhite,
		BorderSide:      creator.CellBorderSideAll,
		BorderStyle:     creator.CellBorderStyleSingle,
		BorderWidth:     6,
	},
	"heading-centered": {
		BackgroundColor: creator.ColorRGBFromHex("#332f3f"),
		HAlignment:      creator.CellHorizontalAlignmentCenter,
		BorderColor:     creator.ColorWhite,
		BorderSide:      creator.CellBorderSideAll,
		BorderStyle:     creator.CellBorderStyleSingle,
		BorderWidth:     6,
	},
	"left-highlighted": {
		BackgroundColor: creator.ColorRGBFromHex("#dde4e5"),
		HAlignment:      creator.CellHorizontalAlignmentLeft,
		BorderColor:     creator.ColorWhite,
		BorderSide:      creator.CellBorderSideAll,
		BorderStyle:     creator.CellBorderStyleSingle,
		BorderWidth:     6,
	},
	"centered-highlighted": {
		BackgroundColor: creator.ColorRGBFromHex("#dde4e5"),
		HAlignment:      creator.CellHorizontalAlignmentCenter,
		BorderColor:     creator.ColorWhite,
		BorderSide:      creator.CellBorderSideAll,
		BorderStyle:     creator.CellBorderStyleSingle,
		BorderWidth:     6,
	},
	"left": {
		HAlignment: creator.CellHorizontalAlignmentLeft,
	},
	"centered": {
		HAlignment: creator.CellHorizontalAlignmentCenter,
	},
	"gradingsys-head": {
		HAlignment: creator.CellHorizontalAlignmentLeft,
	},
	"gradingsys-row": {
		HAlignment: creator.CellHorizontalAlignmentCenter,
	},
	"conduct-head": {
		HAlignment: creator.CellHorizontalAlignmentLeft,
	},
	"conduct-key": {
		HAlignment: creator.CellHorizontalAlignmentLeft,
	},
	"conduct-val": {
		BackgroundColor: creator.ColorRGBFromHex("#dde4e5"),
		HAlignment:      creator.CellHorizontalAlignmentCenter,
		BorderColor:     creator.ColorWhite,
		BorderSide:      creator.CellBorderSideAll,
		BorderStyle:     creator.CellBorderStyleSingle,
		BorderWidth:     3,
	},
}

func GenerateInvoicePdf[T Amount](invoice Invoice[T]) error {
	conf, err := config.GetUniDocCred()
	if err != nil {
		return err
	}

	err = license.SetMeteredKey(conf.Key)
	if err != nil {
		return err
	}

	c := creator.New()
	c.SetPageMargins(40, 40, 0, 0)

	cr := &Client[T]{creator: c}
	err = cr.GeneratePdf(invoice)
	if err != nil {
		return err
	}
	return nil
}
func (c *Client[T]) GeneratePdf(invoice Invoice[T]) error {
	cr := c.creator
	rect := cr.NewRectangle(0, 0, creator.PageSizeLetter[0], 120)
	rect.SetFillColor(creator.ColorRGBFromHex("#dde4e5"))
	rect.SetBorderWidth(0)
	err := cr.Draw(rect)
	if err != nil {
		return err
	}

	headerStyle := cr.NewTextStyle()
	headerStyle.FontSize = 50

	table := cr.NewTable(1)
	table.SetMargins(0, 0, 20, 0)
	err = drawCell(table, newPara(cr, "Sample Invoice", headerStyle), cellStyles["centered"])
	if err != nil {
		return err
	}
	err = cr.Draw(table)
	if err != nil {
		return err
	}

	err = writeInvoice(cr, invoice)
	if err != nil {
		return err
	}

	err = cr.WriteToFile(invoice.Name + "_invoice.pdf")
	if err != nil {
		return err
	}
	return nil
}

func newPara(c *creator.Creator, text string, tstyle creator.TextStyle) *creator.StyledParagraph {
	p := c.NewStyledParagraph()
	p.Append(text).Style = tstyle
	p.SetEnableWrap(false)
	return p
}

func drawCell(table *creator.Table, content creator.VectorDrawable, cellStyle cellStyle) error {
	var cell *creator.TableCell
	if cellStyle.ColSpan > 1 {
		cell = table.MultiColCell(cellStyle.ColSpan)
	} else {
		cell = table.NewCell()
	}
	err := cell.SetContent(content)
	if err != nil {
		return err
	}
	cell.SetHorizontalAlignment(cellStyle.HAlignment)
	if cellStyle.BackgroundColor != nil {
		cell.SetBackgroundColor(cellStyle.BackgroundColor)
	}
	cell.SetBorder(cellStyle.BorderSide, cellStyle.BorderStyle, cellStyle.BorderWidth)
	if cellStyle.BorderColor != nil {
		cell.SetBorderColor(cellStyle.BorderColor)
	}
	if cellStyle.Indent > 0 {
		cell.SetIndent(cellStyle.Indent)
	}
	return nil
}

func writeInvoice[T Amount](c *creator.Creator, invoice Invoice[T]) error {
	headerStyle := c.NewTextStyle()
	// Invoice Header info table.
	table := c.NewTable(2)
	table.SetMargins(0, 0, 50, 0)
	err := drawCell(table, newPara(c, "Business: "+invoice.Name, headerStyle), cellStyles["left"])
	if err != nil {
		return err
	}
	err = drawCell(table, newPara(c, "Address: "+invoice.Address, headerStyle), cellStyles["left"])
	if err != nil {
		return err
	}
	err = c.Draw(table)
	if err != nil {
		return err
	}

	// Invoice items table.
	table = c.NewTable(4)
	table.SetMargins(0, 0, 20, 0)
	err = table.SetColumnWidths(0.4, 0.2, 0.2, 0.2)
	if err != nil {
		return err
	}
	headingStyle := c.NewTextStyle()
	headingStyle.FontSize = 20
	headingStyle.Color = creator.ColorRGBFromHex("#fdfdfd")
	regularStyle := c.NewTextStyle()

	// Draw table header.
	err = drawCell(table, newPara(c, " Title", headingStyle), cellStyles["heading-left"])
	if err != nil {
		return err
	}
	err = drawCell(table, newPara(c, "Quantity", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, newPara(c, "Price", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	err = drawCell(table, newPara(c, "Total", headingStyle), cellStyles["heading-centered"])
	if err != nil {
		return err
	}
	for _, datum := range invoice.InvoiceItems {
		err = drawCell(table, newPara(c, " "+datum.Title, regularStyle), cellStyles["left-highlighted"])
		if err != nil {
			return err
		}
		err = drawCell(table, newPara(c, fmt.Sprintf("%v", datum.Quantity), regularStyle), cellStyles["centered-highlighted"])
		if err != nil {
			return err
		}
		err = drawCell(table, newPara(c, fmt.Sprintf("%v", datum.Price), regularStyle), cellStyles["centered-highlighted"])
		if err != nil {
			return err
		}
		err = drawCell(table, newPara(c, fmt.Sprintf("%v", datum.CalculateTotalAmount()), regularStyle), cellStyles["centered-highlighted"])
		if err != nil {
			return err
		}
	}
	err = c.Draw(table)
	if err != nil {
		return err
	}

	boldStyle := c.NewTextStyle()
	boldStyle.FontSize = 16
	grid := c.NewTable(12)
	grid.SetMargins(0, 0, 50, 0)
	gradeInfoStyle := c.NewTextStyle()

	table = c.NewTable(2)
	err = table.SetColumnWidths(0.6, 0.4)
	if err != nil {
		return err
	}
	err = drawCell(table, newPara(c, "Total Amount:", gradeInfoStyle), cellStyles["conduct-key"])
	if err != nil {
		return err
	}
	err = drawCell(table, newPara(c, fmt.Sprintf("%v", invoice.CalculateInvoiceTotalAmount()), gradeInfoStyle), cellStyles["conduct-val"])
	if err != nil {
		return err
	}
	err = grid.MultiColCell(5).SetContent(table)
	if err != nil {
		return err
	}
	err = c.Draw(grid)
	if err != nil {
		return err
	}
	return nil
}
