package executors

import (
	"crypto/sha256"
	"strconv"

	"github.com/AnnWann/pstu_finance_system/lib/database"
	"github.com/AnnWann/pstu_finance_system/lib/models"
)

func AddPerson(name string, role string) error {
	hash := sha256.New()
	hash.Write([]byte(name))

	person := models.Person{
		Id:     string(hash.Sum(nil)),
		Name:   name,
		Role:   role,
		Credit: 0,
	}

	err := database.InsertPerson(person)
	return err
}

func AddRegister(day string, month string, year string, Type string, giver string, receiver string, amount string, value string, partyShare string, description string) error {
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

	err = database.InsertRegister(register)
	return err
}
