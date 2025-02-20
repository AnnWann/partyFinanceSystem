package models

type Relatorio_mensal_complexo struct {
	ID                      string                       `json:"id"`
	Mes                     string                       `json:"mes"`
	Ano                     string                       `json:"ano"`
	Membros                 map[int]Membro               `json:"membros"`
	Membros_apos_pagamentos map[int]Membro               `json:"membros_após_pagamentos"`
	Nucleo                  Nucleo                       `json:"nucleo"`
	Partido                 Partido                      `json:"partido"`
	Pagamentos_de_membros   SubRelatorio                 `json:"pagamentos_de_membros"`
	Vendas_jornal           SubRelatorio                 `json:"vendas_jornal"`
	Gastos                  SubRelatorio                 `json:"gastos"`
	Registros_especificos   Registros_Especificos_Nucleo `json:"registros_especificos"`
	Total_Ganho             float64                      `json:"total_ganho"`
	Total_Liquido           float64                      `json:"total_liquido"`
	Pagamento_Partidario    float64                      `json:"pagamento_partidario"`
	Lucro_Nucleo            float64                      `json:"lucro_nucleo"`
	Link_Arquivo            string                       `json:"link_arquivo"`
}

type Relatorio_mensal struct {
	ID                   string   `json:"id"`
	Mes                  string   `json:"mes"`
	Ano                  string   `json:"ano"`
	Membros              []string `json:"membros"`
	Nucleo               int      `json:"nucleo"`
	Registros            []string `json:"registros"`
	Total_Ganho          float64  `json:"total_earned"`
	Total_Liquido        float64  `json:"total_liquid"`
	Pagamento_Partidario float64  `json:"party_debts"`
	Lucro_Nucleo         float64  `json:"core_surplus"`
	Link_Arquivo         string   `json:"file_link"`
}

type SubRelatorio struct {
	Registros []Registro
	Tipo      string
	Total     float64
}

type Registros_Especificos_Nucleo struct {
	Tipos map[int]SubRelatorio
	Total float64
}
