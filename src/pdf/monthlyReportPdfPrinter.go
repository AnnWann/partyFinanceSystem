package pdfMaker

import (
	"fmt"

	"github.com/AnnWann/pstu_finance_system/src/models"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

func PrintPDFMonthlyReport(r models.MonthReport, path string) error {

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

	var receitas_vendas_each_type []*creator.Chapter
	var receitas_vendas_each_type_table []*creator.Table

	for _, st := range r.Sales.EachType {
		receitas_vendas_each_type = append(receitas_vendas_each_type, makeSubSubChapter(receitas_vendas, st.Type, normalFont))
		for _, sr := range st.Registers {
			receitas_vendas_each_type_table = append(receitas_vendas_each_type_table, makeTable(c, []string{"Data", "Descrição", "Quantidade", "Valor"}, normalFont))
			date := sr.Day + "/" + r.Month + "/" + r.Year
			drawCell(c, receitas_vendas_each_type_table[len(receitas_vendas_each_type_table)-1], date, normalFont, creator.CellHorizontalAlignmentCenter)
			drawCell(c, receitas_vendas_each_type_table[len(receitas_vendas_each_type_table)-1], sr.Description, normalFont, creator.CellHorizontalAlignmentCenter)
			drawCell(c, receitas_vendas_each_type_table[len(receitas_vendas_each_type_table)-1], fmt.Sprintf("%d", sr.Amount), normalFont, creator.CellHorizontalAlignmentCenter)
			drawCell(c, receitas_vendas_each_type_table[len(receitas_vendas_each_type_table)-1], fmt.Sprintf("%f", sr.Value), normalFont, creator.CellHorizontalAlignmentCenter)
		}

		drawTotalRow(c, receitas_vendas_each_type_table[len(receitas_vendas_each_type_table)-1], st.Total, boldFont)
	}

	for i := 0; i < len(receitas_vendas_each_type); i++ {
		receitas_vendas_each_type[i].Add(receitas_vendas_each_type_table[i])
		receitas_vendas.Add(receitas_vendas_each_type[i])
	}

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

	// Créditos de Membros

	creditos_membros := makeSubchapter(titulo, "Créditos de Membros", normalFont)
	titulo.Add(creditos_membros)

	creditos_membros_table := makeTable(c, []string{"Nome", "Valor"}, normalFont)

	for _, cr := range r.MembersAfterPaying {
		drawCell(c, creditos_membros_table, cr.Name, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, creditos_membros_table, fmt.Sprintf("%f", cr.Credit), normalFont, creator.CellHorizontalAlignmentCenter)
	}

	creditos_membros.Add(creditos_membros_table)

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

func drawTotalRow(c *creator.Creator, table *creator.Table, total float64, font *model.PdfFont) {
	drawCell(c, table, "Total", font, creator.CellHorizontalAlignmentCenter)
	drawCell(c, table, "", font, creator.CellHorizontalAlignmentCenter)
	drawCell(c, table, "", font, creator.CellHorizontalAlignmentCenter)
	drawCell(c, table, fmt.Sprintf("%f", total), font, creator.CellHorizontalAlignmentCenter)
}
