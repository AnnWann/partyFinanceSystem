package executors

import (
	"crypto/sha256"
	"errors"
	"strconv"

	"github.com/AnnWann/pstu_finance_system/lib/database"
	"github.com/AnnWann/pstu_finance_system/lib/models"
)

func AddRegister(day string, month string, year string, Type string, giver string, receiver string, amount string, value string, partyShare string, description string) error {
	if day == "" || month == "" || year == "" || Type == "" || giver == "" || receiver == "" || amount == "" || value == "" || partyShare == "" || description == "" {
		err := errors.New("arguments cannot be empty. The correct format is 'add register <day> <month> <year> <type> <giver> <receiver> <amount> <value> <partyShare> <description>'")
		return err
	}
	if Type != "payment" && Type != "expense" && Type != "journal" && Type != "other" {
		err := errors.New("can't add register with type " + Type +
			". Type must be 'payment', 'expense', 'journal' or 'other'")
		return err
	}

	dayIsValid, err := strconv.Atoi(day)
	if err != nil || dayIsValid < 1 || dayIsValid > 31 {
		return errors.New("invalid day")
	}

	monthIsValid, err := strconv.Atoi(month)
	if err != nil || monthIsValid < 1 || monthIsValid > 12 {
		return errors.New("invalid month")
	}

	_, err = strconv.Atoi(year)
	if err != nil {
		return errors.New("invalid year")
	}
	
	hash := sha256.New()
	hash.Write([]byte(giver + receiver + day + description))

	id := string(hash.Sum(nil))

	amountINT, err := strconv.Atoi(amount)
	if err != nil {
		return err
	}

	valueFLOAT, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return err
	}

	partyShareFLOAT, err := strconv.ParseFloat(partyShare, 32)
	if err != nil {
		return err
	}

	register := models.Register{
		Id:          id,
		Day:         day,
		Month:       month,
		Year:        year,
		Type:        Type,
		Giver:       giver,
		Receiver:    receiver,
		Amount:      int16(amountINT),
		Value:       float32(valueFLOAT),
		PartyShare:  float32(partyShareFLOAT),
		Description: description,
	}

	err = database.GetDB().GetRegisterDB().InsertRegister(register)
	return err
}

func GetRegister(id string) (models.Register, error) {
	if id == "" {
		return models.Register{}, errors.New("arguments cannot be empty. The correct format is 'get register <id>'")
	}
	return database.GetDB().GetRegisterDB().GetRegister(id)
}

func GetRegisters () ([]models.Register, error) {
	return database.GetDB().GetRegisterDB().GetRegisters()
}

func GetRegisterByMonthAndYear (month string, year string) ([]models.Register, error) {
	if month == "" || year == "" {
		return []models.Register{}, errors.New("arguments cannot be empty. The correct format is 'get register <month> <year>'")
	}
	monthInt, err := strconv.Atoi(month)
    if err != nil || monthInt < 1 || monthInt > 12 {
        return []models.Register{}, errors.New("invalid month")
    }
	_, err = strconv.Atoi(year)
	if err != nil {
			return []models.Register{}, errors.New("invalid year")
	}
	return database.GetDB().GetRegisterDB().GetRegisterByMonthAndYear(month, year)
}

func GetRegistersByYear (year string) ([]models.Register, error) {
	_, err := strconv.Atoi(year)
	if err != nil {
			return []models.Register{}, errors.New("invalid year")
	}
	if year == "" {
		return []models.Register{}, errors.New("arguments cannot be empty. The correct format is 'get register <year>'")
	}
	return database.GetDB().GetRegisterDB().GetRegistersByYear(year)
}
