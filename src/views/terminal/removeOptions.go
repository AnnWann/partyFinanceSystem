package terminal

import (
	"fmt"

	"github.com/AnnWann/pstu_finance_system/src/executors"
)

func (op *Options) Remove() {
	if len(op.Modifiers) == 0 {
		fmt.Println("Remover o que? Use 'remove --person', 'remove --register' ou 'remove --payday'")
		return
	}

	if len(op.Arguments) == 0 {
		fmt.Println("Argumento inválido. Use 'remove --person <id>', 'remove --typeOfRegister <id>' ou 'remove --nucleo <id>'")
		return
	}

	var firstKey string
	for key := range op.Modifiers {
		firstKey = key
		break
	}

	switch firstKey {
	case "--person":
		op.RemovePerson(op.Arguments[0])
	case "--typeOfRegister":
		op.RemoveTypeOfRegister(op.Arguments[0])
	case "--nucleo":
		op.RemoveNucleo(op.Arguments[0])

	default:
		fmt.Println("Modificador inválido")
	}
}

func (op *Options) RemovePerson(id string) {

	err := executors.DeletePerson(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Pessoa removida com sucesso")
}

func (op *Options) RemoveTypeOfRegister(id string) {
	err := executors.DeleteTypeOfRegister(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Tipo de registro removido com sucesso")
}

func (op *Options) RemoveNucleo(id string) {
	err := executors.DeleteNucleo(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Nucleo removido com sucesso")
}