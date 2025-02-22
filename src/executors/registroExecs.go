package executors

import (
	"errors"
	"strconv"

	"github.com/AnnWann/pstu_finance_system/src/database"
	"github.com/AnnWann/pstu_finance_system/src/models"
)

func AddRegistro(dia string, mes string, ano string, nucleo string, tipo_de_registro string, pago_por string, cobrado_por string, quantidade string, valor string, descricao string) (int, error) {
	type_of_registerInt, err := strconv.Atoi(tipo_de_registro)
	if err != nil {
		return 0, errors.New("tipo de registro inválido")
	}

	nucleoInt, err := strconv.Atoi(nucleo)
	if err != nil {
		return 0, errors.New("núcleo inválido")
	}

	pago_porInt, err := strconv.Atoi(pago_por)
	if err != nil {
		return 0, errors.New("pago_por inválido")
	}

	cobrado_porInt, err := strconv.Atoi(cobrado_por)
	if err != nil {
		return 0, errors.New("cobrado_por inválido")
	}

	dayIsValid, err := strconv.Atoi(dia)
	if err != nil || dayIsValid < 1 || dayIsValid > 31 {
		return 0, errors.New("dia inválido")
	}

	monthIsValid, err := strconv.Atoi(mes)
	if err != nil || monthIsValid < 1 || monthIsValid > 12 {
		return 0, errors.New("mês inválido")
	}

	_, err = strconv.Atoi(ano)
	if err != nil {
		return 0, errors.New("ano inválido")
	}

	amountINT, err := strconv.Atoi(quantidade)
	if err != nil {
		return 0, errors.New("quantidade inválida")
	}

	valueFLOAT, err := strconv.ParseFloat(valor, 32)
	if err != nil {
		return 0, errors.New("valor inválido")
	}

	t, err := database.GetDB().GetTiposDeRegistroDB().GetTipo(type_of_registerInt)
	if err != nil {
		return 0, errors.New("tipo de registro não encontrado")
	}
	if (t.Nucleo != nucleoInt && t.ID > 0) && t.Nucleo != database.GetDB().GetPessoasDB().GetNucleoGeral() {
		return 0, errors.New("tipo de registro não pertence ao núcleo")
	}

	nucleoExists := database.GetDB().GetNucleoDB().NucleoExists(nucleoInt)
	if !nucleoExists {
		return 0, errors.New("núcleo não encontrado")
	}

	personDB := database.GetDB().GetPessoasDB()
	if !personDB.PessoaExists(pago_porInt) {
		return 0, errors.New("pago_por não encontrado")
	}

	if !personDB.PessoaExists(cobrado_porInt) {
		return 0, errors.New("cobrado_por não encontrado")
	}

	registerDB := database.GetDB().GetRegisterDB()
	id := registerDB.GetNextId()

	register := models.Registro{
		ID:         id,
		Dia:        dia,
		Mes:        mes,
		Ano:        ano,
		Nucleo:     nucleoInt,
		Tipo:       type_of_registerInt,
		Pagante:    pago_porInt,
		Cobrante:   cobrado_porInt,
		Quantidade: amountINT,
		Valor:      valueFLOAT,
		Descricao:  descricao,
	}

	err = database.GetDB().GetRegisterDB().InsertRegister(register)
	return register.ID, err
}

func GetRegistro(filterOptions map[string]string) ([]models.Registro, error) {

	registros, err := database.GetDB().GetRegisterDB().GetRegister()
	if err != nil {
		return nil, err
	}

	if len(filterOptions) > 0 {
		registros = filterRegistros(registros, filterOptions)
	}

	if len(registros) == 0 {
		return nil, errors.New("nenhum membro encontrado")
	}

	return registros, nil

}

func filterRegistros(registros []models.Registro, filterOptions map[string]string) []models.Registro {
	if filterOptions == nil {
		return registros
	}

	var filteredregistros []models.Registro
	for _, register := range registros {
		if filterRegistro(register, filterOptions) {
			filteredregistros = append(filteredregistros, register)
		}
	}

	return filteredregistros
}

func filterRegistro(register models.Registro, filterOptions map[string]string) bool {
	isValid := false
	for key, value := range filterOptions {
		switch key {
		case "--nucleo":
			valueInt, err := strconv.Atoi(value)
			if err != nil {
				isValid = false
			} else {
				isValid = register.ID == valueInt
			}
		case "--dia":
			isValid = register.Dia == value
		case "--mes":
			isValid = register.Mes == value
		case "--ano":
			isValid = register.Ano == value
		case "--tipo":
			valueInt, err := strconv.Atoi(value)
			if err != nil {
				isValid = false
			} else {
				isValid = register.Tipo == valueInt
			}
		case "--pagante":
			valueInt, err := strconv.Atoi(value)
			if err != nil {
				isValid = false
			} else {
				isValid = register.Pagante == valueInt
			}
		case "--cobrante":
			valueInt, err := strconv.Atoi(value)
			if err != nil {
				isValid = false
			} else {
				isValid = register.Cobrante == valueInt
			}
		}
	}
	return isValid
}

func AddTipoDeRegistro(nome string, nucleo string, descricao string, partilha_partidaria string) error {
	nucleoInt, err := strconv.Atoi(nucleo)
	if err != nil {
		return errors.New("núcleo inválido")
	}

	nucleoExists := database.GetDB().GetNucleoDB().NucleoExists(nucleoInt)
	if !nucleoExists {
		return errors.New("núcleo não encontrado")
	}

	id, err := database.GetDB().GetTiposDeRegistroDB().GetNextId()
	if err != nil {
		return errors.New("erro ao buscar id")
	}

	party_shareFLOAT, err := strconv.ParseFloat(partilha_partidaria, 32)
	if err != nil {
		return errors.New("parcela partidária inválida")
	}

	typeOfRegister := models.Tipo_de_registro{
		ID:                 id,
		Nome:               nome,
		Nucleo:             nucleoInt,
		Descricao:          descricao,
		Parcela_partidaria: party_shareFLOAT,
	}

	return database.GetDB().GetTiposDeRegistroDB().InsertTipo(typeOfRegister)
}

func GetTipoDeRegistro(nucleo string) ([]models.Tipo_de_registro, error) {
	nucleoInt, err := strconv.Atoi(nucleo)
	if err != nil {
		return nil, errors.New("núcleo inválido")
	}

	types, err := database.GetDB().GetTiposDeRegistroDB().GetTipoPorNucleo(nucleoInt)
	if err != nil {
		return nil, err
	}

	tipos_gerais, err := database.GetDB().GetTiposDeRegistroDB().GetTiposGeral()
	if err != nil {
		return nil, err
	}
	types = append(types, tipos_gerais...)

	return types, nil
}

func UpdateTipoDeRegistro(id string, partilha_partidaria string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("id inválido")
	}
	partilhaPartidariaFLOAT, err := strconv.ParseFloat(partilha_partidaria, 32)
	if err != nil {
		return errors.New("parcela partidária inválida")
	}
	err = database.GetDB().GetTiposDeRegistroDB().UpdatePartilhaPartidaria(idInt, partilhaPartidariaFLOAT)
	return err
}

func DeleteTipoDeRegistro(id string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("id inválido")
	}
	err = database.GetDB().GetTiposDeRegistroDB().DeleteTipo(idInt)
	return err
}
