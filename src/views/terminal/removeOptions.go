package terminal

import (
	"fmt"

	"github.com/AnnWann/pstu_finance_system/src/executors"
)

func (op *Options) Remove() {
	if len(op.Modifiers) == 0 {
		fmt.Println("Remover o que? Use:\n '--pessoa:" + op.RemoveModifiers["--pessoa"] + "'\n'--tipoDeRegistro:" + op.RemoveModifiers["--tipoDeRegistro"] + "'\n'--nucleo:" + op.RemoveModifiers["--nucleo"] + "'")
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
	case "--pessoa":
		op.RemovePerson(op.Arguments[0])
	case "--tipoDeRegistro":
		op.RemoveTypeOfRegister(op.Arguments[0])
	case "--nucleo":
		op.RemoveNucleo(op.Arguments[0])

	default:
		fmt.Println("Modificador inválido")
	}
}

func (op *Options) RemovePerson(id string) {

	err := executors.DeleteMembro(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Pessoa removida com sucesso")
}

func (op *Options) RemoveTypeOfRegister(id string) {
	err := executors.DeleteTipoDeRegistro(id)
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
