package terminal

import (
	"fmt"
	"os"

	"github.com/AnnWann/pstu_finance_system/src/executors"
	"github.com/joho/godotenv"
)

func (op *Options) Get() {
	if len(op.Modifiers) == 0 {
		fmt.Println("Obter o que? Use 'get --person', 'get --register' ou 'get --payday'")
		return
	}

	var firstKey string
	for key := range op.Modifiers {
		firstKey = key
		break
	}

	modifiersRest := make(map[string]string)

	for key, value := range op.Modifiers {
		if key != firstKey {
			modifiersRest[key] = value
		}
	}

	switch firstKey {
	case "--person":
		op.GetPerson(modifiersRest)
	case "--register":
		op.GetRegister(modifiersRest)
	case "--report":
		op.GetReport(modifiersRest)
	case "--payday":
		if len(op.Arguments) == 0 {
			fmt.Println("Commando invalido, considere: " + op.Commands["get"] + " " + op.GetModifiers["--payday"])
			return
		}
		op.GetPayday(op.Arguments[0])
	case "--typeOfRegister":
		if len(op.Arguments) == 0 {
			fmt.Println("Commando invalido, considere: " + op.Commands["get"] + " " + op.GetModifiers["--typeOfRegister"])
			return
		}
		op.GetTypeOfRegister(op.Arguments[0])
	case "--nucleo":
		op.GetNucleo(modifiersRest)
	case "--partido":
		op.GetPartido()

	default:
		fmt.Println("Modificador inválido")
	}
}

func (op *Options) GetPerson(modifiers map[string]string) {
	if len(modifiers) == 0 {
		fmt.Println("Commando invalido, considere: " + op.Commands["get"] + " " + op.GetModifiers["--person"])
		return
	}

	persons, err := executors.GetPerson(modifiers)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, person := range persons {
		fmt.Printf("dados da pessoa:\n id: %s\nnome: %s\npapel: %s\nnucleo: %d\ncontribuição mensal: %f\ncredito: %f\n", person.Id, person.Name, person.Role, person.Nucleo, person.MonthlyPayment, person.Credit)
	}	
}

func (op *Options) GetRegister(modifiers map[string]string) {
	if len(modifiers) == 0 {
		fmt.Println("Commando invalido, considere: " + op.Commands["get"] + " " + op.GetModifiers["--register"])
		return
	}

	register, err := executors.GetRegister(modifiers)
	if err != nil {
		fmt.Println(err)
	}

	for _, reg := range register {
		fmt.Printf("dados do registro:\n id: %d\ndia: %s\nmês: %s\nano: %squantidade: %dvalor: %f\ntipo: %d\nnucleo: %d\ndoador: %s\nrecebedor: %s\ndescrição: %s\n", reg.Id, reg.Day, reg.Month, reg.Year, reg.Amount, reg.Value, reg.Type, reg.Nucleo, reg.Giver, reg.Receiver, reg.Description)
	}
}

func (op *Options) GetReport(modifiers map[string]string) {
	if len(modifiers) == 0 {
		fmt.Println("Commando invalido, considere: " + op.Commands["get"] + " " + op.GetModifiers["--report"])
		return
	}

	err := executors.GetMonthlyReport(modifiers)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	pdf_folder := os.Getenv("PDF_FOLDER")

	fmt.Printf("Relatórios gerados em: %s\n", pdf_folder)
}

func (op *Options) GetTypeOfRegister(nucleoId string) {

	typeOfRegister, err := executors.GetTypeOfRegister(nucleoId)
	if err != nil {
		fmt.Println(err)
	}

	for _, tReg := range typeOfRegister {
		fmt.Printf("dados do tipo de registro:\n id: %d\nnome: %s\nnucleo: %s\ndescrição: %s\npartilha partidária: %f\n", tReg.Id, tReg.Name, tReg.Nucleo, tReg.Description, tReg.PartyShare)
	}
	}

func (op *Options) GetNucleo(modifiers map[string]string) {
	if len(modifiers) == 0 {
		fmt.Println("Commando invalido, considere: " + op.Commands["get"] + " " + op.GetModifiers["--nucleo"])
		return
	}

	nucleos, err := executors.GetNucleo(modifiers)
	if err != nil {
		fmt.Println(err)
	}

	for _, nucleo := range nucleos {
		fmt.Printf("dados do nucleo:\n id: %d\nnome: %s\ncidade: %s\nestado: %s\ncredito: %f\ndia de pagamento: %s\n", nucleo.Id, nucleo.Name, nucleo.City, nucleo.State, nucleo.Credit, nucleo.Payday)
	}
}

func (op *Options) GetPartido() {
	partido, err := executors.GetPartido()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("dados do partido:\n id: %d\nnome: %s\ncredito: %f", partido.Id, partido.Name, partido.Credit)
}


func (op *Options) GetPayday(nucleoId string) {
	payday, err := executors.GetPayday(nucleoId)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("dia de pagamento: %s", payday)
}




