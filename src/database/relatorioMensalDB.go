package database

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/AnnWann/pstu_finance_system/src/models"
)

type RelatorioMensalDB struct {
	*sql.DB
}

func (db *DBWrapper) GetRelatorioMensalDB() *RelatorioMensalDB {
	return &RelatorioMensalDB{db.DB}
}

func (db *RelatorioMensalDB) GetNextId() (string, error) {
	var id string
	err := db.QueryRow("SELECT COUNT(*) FROM relatorios_mensais").Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (rDb *RelatorioMensalDB) InsertRelatorio(r models.Relatorio_mensal_complexo) error {

	DB := GetDB()
	PersonDB := DB.GetMembroDB()
	for _, m := range r.Membros_apos_pagamentos {
		err := PersonDB.UpdateCredito(m.ID, m.Credito)
		if err != nil {
			return err
		}
	}

	err := DB.GetNucleoDB().UpdateReserva(r.Nucleo.ID, r.Nucleo.Reserva)
	if err != nil {
		return err
	}

	membersIds := make([]int, len(r.Membros))
	for _, value := range r.Membros {
		membersIds = append(membersIds, value.ID)
	}

	membersIdsStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(membersIds)), ", "), "[]")

	sizeRegister := len(r.Pagamentos_de_membros.Registros) + len(r.Gastos.Registros) + len(r.Registros_especificos.Tipos)
	for _, salesType := range r.Registros_especificos.Tipos {
		sizeRegister += len(salesType.Registros)
	}

	registerIds := make([]int, sizeRegister)
	for i, payment := range r.Pagamentos_de_membros.Registros {
		registerIds[i] = payment.ID
	}

	for i, expense := range r.Gastos.Registros {
		registerIds[i] = expense.ID
	}

	for _, e_t := range r.Registros_especificos.Tipos {
		for _, r := range e_t.Registros {
			registerIds = append(registerIds, r.ID)
		}
	}

	registerIdsStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(registerIds)), ", "), "[]")

	_, err = rDb.Exec(
		"INSERT INTO"+
			"relatorios_mensais (id, mes, ano, nucleo, membros, registros, total_ganho, total_liquido, pagamento_partido, lucro_nucleo, link_arquivo)"+
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		r.ID, r.Mes, r.Ano, r.Nucleo.ID, membersIdsStr, registerIdsStr, r.Total_Ganho,
		r.Total_Liquido, r.Pagamento_Partidario, r.Lucro_Nucleo, r.Link_Arquivo)

	return err
}

func (db RelatorioMensalDB) RelatorioExists(nucleo int, month string, year string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM relatorios_mensais WHERE nucleo = ? AND mes = ? AND ano = ?", nucleo, month, year).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (db *RelatorioMensalDB) GetRelatorioMensal() ([]models.Relatorio_mensal, error) {
	row, err := db.Query("SELECT * FROM relatorios_mensais")
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var relatorios []models.Relatorio_mensal
	for row.Next() {
		var r models.Relatorio_mensal
		err = row.Scan(&r.ID, &r.Mes, &r.Ano, &r.Membros, &r.Registros, &r.Total_Ganho,
			&r.Total_Liquido, &r.Pagamento_Partidario, &r.Lucro_Nucleo, &r.Link_Arquivo)
		if err != nil {
			return nil, err
		}
		relatorios = append(relatorios, r)
	}

	return relatorios, nil
}

func (db *RelatorioMensalDB) GetRelatorioMensalById(id string) (models.Relatorio_mensal, error) {
	var r models.Relatorio_mensal
	err := db.QueryRow("SELECT * FROM relatorios_mensais WHERE id = ?", id).
		Scan(&r.ID, &r.Mes, &r.Ano, &r.Membros, &r.Registros, &r.Total_Ganho,
			&r.Total_Liquido, &r.Pagamento_Partidario, &r.Lucro_Nucleo, &r.Link_Arquivo)
	if err != nil {
		return r, err
	}

	return r, nil
}
