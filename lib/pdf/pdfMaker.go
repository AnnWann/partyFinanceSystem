package pdfMaker

import (
	"fmt"

	"github.com/AnnWann/pstu_finance_system/lib/models"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

func MakePDFReport(r models.MonthReport) error {

	c := creator.New()
	normalFont, err := model.NewStandard14Font("Helvetica")
	if err != nil {
		return err
	}
	boldFont, err := model.NewStandard14Font("Helvetica-Bold")
	if err != nil {
		return err
	}

	titulo := makeChapter(c, "Relatório de Finanças de "+r.Month+"/"+r.Year+" - PSTU - UFBA", boldFont)

	// Receitas

	receitas := makeSubchapter(titulo, "Receitas", normalFont)
	titulo.Add(receitas)

	// Membros

	receitas_membros := makeSubSubChapter(receitas, "Pagamentos de Membros", normalFont)
	receitas.Add(receitas_membros)

	receitas_membros_table := makeTable(c, []string{"Data", "Nome", "Designação", "Valor"}, normalFont)

	for _, mr := range r.MembersPayments.Registers {
		date := mr.Day + "/" + r.Month + "/" + r.Year
		person := r.Members[mr.Giver]
		drawCell(c, receitas_membros_table, date, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, receitas_membros_table, mr.Giver, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, receitas_membros_table, person.Name, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, receitas_membros_table, fmt.Sprintf("%f", mr.Value), normalFont, creator.CellHorizontalAlignmentCenter)
	}

	drawTotalRow(c, receitas_membros_table, r.MembersPayments.Total, boldFont)

	receitas_membros.Add(receitas_membros_table)

	// Vendas

	receitas_vendas := makeSubSubChapter(receitas, "Vendas", normalFont)
	receitas.Add(receitas_vendas)

	// Jornal
	receitas_vendas_jornal := makeSubSubChapter(receitas_vendas, "Jornal", normalFont)
	receitas_vendas.Add(receitas_vendas_jornal)

	receitas_vendas_jornal_table := makeTable(c, []string{"Data", "Vendedor", "Quantidade", "Valor-partido", "Valor"}, normalFont)

	for _, jr := range r.Sales.Jornals.Registers {
		date := jr.Day + "/" + r.Month + "/" + r.Year
		person := r.Members[jr.Receiver]
		drawCell(c, receitas_vendas_jornal_table, date, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, receitas_vendas_jornal_table, person.Name, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, receitas_vendas_jornal_table, person.Name, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, receitas_vendas_jornal_table, fmt.Sprintf("%d", jr.Amount), normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, receitas_vendas_jornal_table, fmt.Sprintf("%f", jr.PartyShare), normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, receitas_vendas_jornal_table, fmt.Sprintf("%f", jr.Value), normalFont, creator.CellHorizontalAlignmentCenter)
	}

	drawTotalRow(c, receitas_vendas_jornal_table, r.Sales.Jornals.Total, boldFont)

	receitas_vendas_jornal.Add(receitas_vendas_jornal_table)

	// Outras vendas

	receitas_vendas_outros := makeSubSubChapter(receitas_vendas, "Outros", normalFont)
	receitas_vendas.Add(receitas_vendas_outros)

	receitas_vendas_outros_table := makeTable(c, []string{"Data", "Descrição", "Vendedor", "Quantidade", "Valor-Partido", "Valor"}, normalFont)

	for _, or := range r.Sales.Others.Registers {
		date := or.Day + "/" + r.Month + "/" + r.Year
		person := r.Members[or.Receiver]
		drawCell(c, receitas_vendas_outros_table, date, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, receitas_vendas_outros_table, or.Description, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, receitas_vendas_outros_table, person.Name, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, receitas_vendas_outros_table, fmt.Sprintf("%d", or.Amount), normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, receitas_vendas_outros_table, fmt.Sprintf("%f", or.PartyShare), normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, receitas_vendas_outros_table, fmt.Sprintf("%f", or.Value), normalFont, creator.CellHorizontalAlignmentCenter)
	}

	drawTotalRow(c, receitas_vendas_outros_table, r.Sales.Others.Total, boldFont)

	receitas_vendas_outros.Add(receitas_vendas_outros_table)

	total_vendas_row := makeTable(c, []string{"Subtotal de vendas", "", "", fmt.Sprintf("%f", r.Sales.TotalSales)}, boldFont)
	receitas_vendas.Add(total_vendas_row)

	total_receitas_row := makeTable(c, []string{"Total de Receitas", "", "", fmt.Sprintf("%f", r.TotalEarned)}, boldFont)
	receitas.Add(total_receitas_row)

	// Despesas

	despesas := makeSubchapter(titulo, "Despesas", normalFont)
	titulo.Add(despesas)

	despesas_table := makeTable(c, []string{"Data", "Descrição", "Gasto por", "Quantidade", "Valor"}, normalFont)

	for _, er := range r.Expenses.Registers {
		date := er.Day + "/" + r.Month + "/" + r.Year
		person := r.Members[er.Giver]
		drawCell(c, despesas_table, date, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, despesas_table, er.Description, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, despesas_table, person.Name, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, despesas_table, fmt.Sprintf("%d", er.Amount), normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, despesas_table, fmt.Sprintf("%f", er.Value), normalFont, creator.CellHorizontalAlignmentCenter)
	}

	drawTotalRow(c, despesas_table, r.Expenses.Total, boldFont)

	despesas.Add(despesas_table)

	// Valores finais

	valores_finais := makeSubchapter(titulo, "Valores Finais", normalFont)
	titulo.Add(valores_finais)

	total_liquido_row := makeTable(c, []string{"Total Líquido", "", "", fmt.Sprintf("%f", r.TotalLiquid)}, boldFont)
	valores_finais.Add(total_liquido_row)

	divida_partido_row := makeTable(c, []string{"Parte do Partido", "", "", fmt.Sprintf("%f", r.PartyDebts)}, boldFont)
	valores_finais.Add(divida_partido_row)

	superavit_core_row := makeTable(c, []string{"Superávit do Núcleo", "", "", fmt.Sprintf("%f", r.CoreSurplus)}, boldFont)
	valores_finais.Add(superavit_core_row)

	err = c.WriteToFile("report.pdf")
	return err
}

func makeChapter(c *creator.Creator, title string, font *model.PdfFont) *creator.Chapter {
	chapter := c.NewChapter(title)
	heading := chapter.GetHeading()
	heading.SetFontSize(24)
	heading.SetFont(font)
	heading.SetColor(creator.ColorBlack)
	return chapter
}

func makeSubchapter(c *creator.Chapter, title string, font *model.PdfFont) *creator.Chapter {
	subchapter := c.NewSubchapter(title)
	subheading := subchapter.GetHeading()
	subheading.SetFontSize(18)
	subheading.SetFont(font)
	subheading.SetColor(creator.ColorBlack)
	return subchapter
}

func makeSubSubChapter(c *creator.Chapter, title string, font *model.PdfFont) *creator.Chapter {
	subsubchapter := c.NewSubchapter(title)
	subsubheading := subsubchapter.GetHeading()
	subsubheading.SetFont(font)
	subsubheading.SetFontSize(14)
	subsubheading.SetColor(creator.ColorBlack)
	return subsubchapter
}

func makeStyledParagraph(c *creator.Creator, text string, font *model.PdfFont) *creator.StyledParagraph {
	paragraph := c.NewStyledParagraph()
	paragraph.SetMargins(0, 0, 15, 15)
	chunk := paragraph.Append(text)
	chunk.Style.Font = font
	chunk.Style.FontSize = 12
	chunk.Style.Color = creator.ColorBlack

	return paragraph
}

func makeTable(c *creator.Creator, fields []string, font *model.PdfFont) *creator.Table {
	table := c.NewTable(len(fields))
	table.SetMargins(0, 0, 15, 15)

	for _, name := range fields {
		drawCell(c, table, name, font, creator.CellHorizontalAlignmentCenter)
	}

	return table
}

func drawCell(c *creator.Creator, table *creator.Table, text string, font *model.PdfFont, align creator.CellHorizontalAlignment) {
	p := makeStyledParagraph(c, "", font)
	p.Append(text).Style.Font = font

	cell := table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(align)
	cell.SetContent(p)
}

func drawTotalRow(c *creator.Creator, table *creator.Table, total float32, font *model.PdfFont) {
	drawCell(c, table, "Total", font, creator.CellHorizontalAlignmentCenter)
	drawCell(c, table, "", font, creator.CellHorizontalAlignmentCenter)
	drawCell(c, table, "", font, creator.CellHorizontalAlignmentCenter)
	drawCell(c, table, fmt.Sprintf("%f", total), font, creator.CellHorizontalAlignmentCenter)
}
