package pdfMaker

import (
	"fmt"
	"strconv"

	"github.com/AnnWann/pstu_finance_system/src/models"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

func PrintPDFMonthlyReport(r models.Relatorio_mensal_complexo, path string) error {

	c := creator.New()
	normalFont, err := model.NewStandard14Font("Helvetica")
	if err != nil {
		return err
	}
	boldFont, err := model.NewStandard14Font("Helvetica-Bold")
	if err != nil {
		return err
	}

	c.NewPage()

	titulo := makeChapter(c, "Relatório de Finanças de "+r.Mes+"/"+r.Ano+" -"+r.Partido.Nome+" - "+r.Nucleo.Nome, boldFont)
	// Receitas

	receitas := makeSubchapter(titulo, "Receitas", normalFont)

	// Membros

	receitas_membros := makeSubSubChapter(receitas, "Pagamentos de Membros", normalFont)

	receitas_membros_table := makeTable(c, []string{"Data", "Nome", "Designação", "Valor"}, normalFont)

	for _, mr := range r.Pagamentos_de_membros.Registros {
		date := mr.Dia + "/" + r.Mes + "/" + r.Ano
		person := r.Membros[mr.Pago_por]
		drawCell(c, receitas_membros_table, date, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, receitas_membros_table, strconv.Itoa(mr.Pago_por), normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, receitas_membros_table, person.Nome, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, receitas_membros_table, fmt.Sprintf("%.2f", mr.Valor), normalFont, creator.CellHorizontalAlignmentCenter)
	}

	drawTotalRow(c, receitas_membros_table, r.Pagamentos_de_membros.Total, boldFont, 2)

	receitas_membros.Add(receitas_membros_table)

	// Jornais

	receitas_jornais := makeSubSubChapter(receitas, "Venda de Jornais", normalFont)

	receitas_jornais_table := makeTable(c, []string{"Data", "Quantidade", "Valor"}, normalFont)
	for _, mr := range r.Vendas_jornal.Registros {
		date := mr.Dia + "/" + r.Mes + "/" + r.Ano
		drawCell(c, receitas_jornais_table, date, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, receitas_jornais_table, fmt.Sprintf("%d", mr.Quantidade), normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, receitas_jornais_table, fmt.Sprintf("%.2f", mr.Valor), normalFont, creator.CellHorizontalAlignmentCenter)
	}
	drawTotalRow(c, receitas_jornais_table, r.Vendas_jornal.Total, boldFont, 1)
	receitas_jornais.Add(receitas_jornais_table)

	// Vendas

	receitas_vendas := makeSubSubChapter(receitas, "Registros Especificos do Nucleo", normalFont)

	var receitas_vendas_each_type []*creator.Chapter
	var receitas_vendas_each_type_table []*creator.Table

	for _, st := range r.Registros_especificos.Tipos {
		receitas_vendas_each_type = append(receitas_vendas_each_type, makeSubSubChapter(receitas_vendas, st.Tipo, normalFont))
		receitas_vendas_each_type_table = append(receitas_vendas_each_type_table, makeTable(c, []string{"Data", "Descrição", "Quantidade", "Valor"}, normalFont))
		for _, sr := range st.Registros {
			date := sr.Dia + "/" + r.Mes + "/" + r.Ano
			drawCell(c, receitas_vendas_each_type_table[len(receitas_vendas_each_type_table)-1], date, normalFont, creator.CellHorizontalAlignmentCenter)
			drawCell(c, receitas_vendas_each_type_table[len(receitas_vendas_each_type_table)-1], sr.Descricao, normalFont, creator.CellHorizontalAlignmentCenter)
			drawCell(c, receitas_vendas_each_type_table[len(receitas_vendas_each_type_table)-1], fmt.Sprintf("%d", sr.Quantidade), normalFont, creator.CellHorizontalAlignmentCenter)
			drawCell(c, receitas_vendas_each_type_table[len(receitas_vendas_each_type_table)-1], fmt.Sprintf("%.2f", sr.Valor), normalFont, creator.CellHorizontalAlignmentCenter)
		}

		drawTotalRow(c, receitas_vendas_each_type_table[len(receitas_vendas_each_type_table)-1], st.Total, boldFont, 2)
	}

	for i := 0; i < len(receitas_vendas_each_type); i++ {
		receitas_vendas_each_type[i].Add(receitas_vendas_each_type_table[i])
	}

	// Despesas

	despesas := makeSubchapter(titulo, "Despesas", normalFont)

	despesas_table := makeTable(c, []string{"Data", "Descrição", "Gasto por", "Quantidade", "Valor"}, normalFont)

	for _, er := range r.Gastos.Registros {
		date := er.Dia + "/" + r.Mes + "/" + r.Ano
		person := r.Membros[er.Pago_por]
		drawCell(c, despesas_table, date, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, despesas_table, er.Descricao, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, despesas_table, person.Nome, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, despesas_table, fmt.Sprintf("%d", er.Quantidade), normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, despesas_table, fmt.Sprintf("%.2f", er.Valor), normalFont, creator.CellHorizontalAlignmentCenter)
	}

	drawTotalRow(c, despesas_table, r.Gastos.Total, boldFont, 3)

	despesas.Add(despesas_table)

	// Valores finais

	valores_finais := makeSubchapter(titulo, "Valores Finais", normalFont)

	total_arrecadado := makeStyledParagraph(c, "Total Arrecadado: "+fmt.Sprintf("%.2f", r.Total_Ganho), normalFont)
	valores_finais.Add(total_arrecadado)
	total_liquido := makeStyledParagraph(c, "Total Líquido: "+fmt.Sprintf("%.2f", r.Total_Liquido), normalFont)
	valores_finais.Add(total_liquido)
	pagamento_partido := makeStyledParagraph(c, "Parte do Partido: "+fmt.Sprintf("%.2f", r.Pagamento_Partidario), normalFont)
	valores_finais.Add(pagamento_partido)
	lucro_nucleo := makeStyledParagraph(c, "Lucro do Núcleo: "+fmt.Sprintf("%.2f", r.Lucro_Nucleo), normalFont)
	valores_finais.Add(lucro_nucleo)
	// Créditos de Membros

	creditos_membros := makeSubchapter(titulo, "Créditos de Membros", normalFont)

	creditos_membros_table := makeTable(c, []string{"Nome", "Valor"}, normalFont)

	for _, cr := range r.Membros_apos_pagamentos {
		drawCell(c, creditos_membros_table, cr.Nome, normalFont, creator.CellHorizontalAlignmentCenter)
		drawCell(c, creditos_membros_table, fmt.Sprintf("%.2f", cr.Credito), normalFont, creator.CellHorizontalAlignmentCenter)
	}

	creditos_membros.Add(creditos_membros_table)

	if err = c.Draw(titulo); err != nil {
		return err
	}

	return c.WriteToFile(path + "/RelatorioMensal" + strconv.Itoa(r.Nucleo.ID) + r.Mes + r.Ano + ".pdf")

}

func makeChapter(c *creator.Creator, title string, font *model.PdfFont) *creator.Chapter {
	chapter := c.NewChapter(title)
	heading := chapter.GetHeading()
	heading.SetFontSize(20)
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

func drawTotalRow(c *creator.Creator, table *creator.Table, total float64, font *model.PdfFont, spaces int) {
	drawCell(c, table, "Total", font, creator.CellHorizontalAlignmentCenter)
	for i := 0; i < spaces; i++ {
		drawCell(c, table, "", font, creator.CellHorizontalAlignmentCenter)
	}
	drawCell(c, table, fmt.Sprintf("%.2f", total), font, creator.CellHorizontalAlignmentCenter)
}
