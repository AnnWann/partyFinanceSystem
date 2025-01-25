package executors

import (
	"errors"
	"strconv"

	"github.com/AnnWann/pstu_finance_system/src/database"
	"github.com/AnnWann/pstu_finance_system/src/models"
)

func AddRegister(day string, month string, year string, nucleo string, type_of_register string, giver string, receiver string, amount string, value string, description string) (int, error) {
	type_of_registerInt , err := strconv.Atoi(type_of_register)
	if err != nil {
		return 0, errors.New("invalid type of register")
	}
	nucleoInt, err := strconv.Atoi(nucleo)
	if err != nil {
		return 0, errors.New("invalid nucleo")
	}

	_, err = database.GetDB().GetTypeOfRegisterDB().GetType(type_of_registerInt)
	if err != nil {
		return 0, errors.New("Tipo de registro não encontrado")
	}

	nucleoExists := database.GetDB().GetNucleoDB().NucleoExists(nucleoInt)
	if !nucleoExists {
		return 0, errors.New("Núcleo não encontrado")
	}

	personDB := database.GetDB().GetPersonDB()
	_, err = personDB.GetPersonById(giver)
	if err != nil {
		return 0, errors.New("Doador não encontrado")
	}

	_, err = personDB.GetPersonById(receiver)
	if err != nil {
		return 0, errors.New("Recebedor não encontrado")
	}

	dayIsValid, err := strconv.Atoi(day)
	if err != nil || dayIsValid < 1 || dayIsValid > 31 {
		return 0, errors.New("invalid day")
	}

	monthIsValid, err := strconv.Atoi(month)
	if err != nil || monthIsValid < 1 || monthIsValid > 12 {
		return 0, errors.New("invalid month")
	}

	_, err = strconv.Atoi(year)
	if err != nil {
		return 0, errors.New("invalid year")
	}

	amountINT, err := strconv.Atoi(amount)
	if err != nil {
		return 0, errors.New("invalid amount")
	}

	valueFLOAT, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, errors.New("invalid value")
	}

	registerDB := database.GetDB().GetRegisterDB()
	id, err := registerDB.GetNextId()

	register := models.Register{
		Id:          id,
		Day:         day,
		Month:       month,
		Year:        year,
		Nucleo:      nucleoInt,
		Type:        type_of_registerInt,
		Giver:       giver,
		Receiver:    receiver,
		Amount:      amountINT,
		Value:       float64(valueFLOAT),
		Description: description,
	}

	err = database.GetDB().GetRegisterDB().InsertRegister(register)
	return register.Id, err
}



func GetRegister(filterOptions map[string]string) ([]models.Register, error) {
	
	registers, err := database.GetDB().GetRegisterDB().GetRegister()
	if err != nil {
		return nil, err
	}

	return filterRegisters(registers, filterOptions), nil
	
}

func filterRegisters(registers []models.Register, filterOptions map[string]string) []models.Register {
	if filterOptions == nil {
		return registers
	}

	var filteredRegisters []models.Register
	for _, register := range registers {
		if filterRegister(register, filterOptions) {
			filteredRegisters = append(filteredRegisters, register)
		}
	}

	return filteredRegisters
}

func filterRegister(register models.Register, filterOptions map[string]string) bool {
	isValid := false
	for key, value := range filterOptions {
		switch key {
		case "--nucleo": 
			valueInt, err := strconv.Atoi(value)
			if err != nil {
				return false
			}
			isValid = register.Id == valueInt
		case "--day":
			isValid = register.Day == value
		case "--month":
			isValid = register.Month == value
		case "--year":
			isValid = register.Year == value
		case "--type":
			valueInt, err := strconv.Atoi(value)
			if err != nil {
				return false
			}
			isValid = register.Type == valueInt
		case "--giver":
			isValid = register.Giver == value
		case "--receiver":
			isValid = register.Receiver == value
		}
	}	
	return isValid
}

func AddTypeOfRegister(name string, nucleo string, description string, party_share string) error {
	nucleoInt , err := strconv.Atoi(nucleo)
	if err != nil {
		return errors.New("Núcleo invalido")
	}

	nucleoExists := database.GetDB().GetNucleoDB().NucleoExists(nucleoInt)
	if !nucleoExists {
		return errors.New("Núcleo não encontrado")
	}

	id, err := database.GetDB().GetTypeOfRegisterDB().GetNextId()
	if err != nil {
		return errors.New("Erro ao buscar id")
	}

	party_shareFLOAT, err := strconv.ParseFloat(party_share, 32)
	if err != nil {
		return errors.New("invalid party share")
	}

	typeOfRegister := models.TypeOfRegister{
		Id:          id,
		Name:        name,
		Nucleo:      nucleo,
		Description: description,
		PartyShare:  party_shareFLOAT,
	}
	
	return database.GetDB().GetTypeOfRegisterDB().InsertType(typeOfRegister)
}

func GetTypeOfRegister(nucleo string) ([]models.TypeOfRegister, error) {
	nucleoInt , err := strconv.Atoi(nucleo)
	if err != nil {
		return nil, errors.New("invalid nucleo")
	}
	
	types, err := database.GetDB().GetTypeOfRegisterDB().GetTypesByNucleo(nucleoInt)
	if err != nil {
		return nil, err
	}

	return types, nil
}

func UpdateTypeOfRegister(id string, partyShare string) error {
	if id == "" {
		return errors.New("arguments cannot be empty. the correct format is 'update <id> <partyShare>'")
	}
	partyShareFLOAT, err := strconv.ParseFloat(partyShare, 32)
	if err != nil {
		return errors.New("invalid party share")
	}
	err = database.GetDB().GetTypeOfRegisterDB().UpdatePartyShare(id, partyShareFLOAT)
	return err
}

func DeleteTypeOfRegister(id string) error {
	if id == "" {
		return errors.New("arguments cannot be empty. the correct format is 'delete <id>'")
	}
	err := database.GetDB().GetTypeOfRegisterDB().DeleteType(id)
	return err
}