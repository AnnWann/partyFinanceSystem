package terminal

import (
	"fmt"

	"github.com/AnnWann/pstu_finance_system/src/executors"
)

func (op *Options) Update() {
	if len(op.Arguments) == 0 {
		fmt.Println("Atualizar o que? Use:\n'update --pessoa: " + op.AddModifiers["--pessoa"] + "'\n" + "'update --tipoDeRegistro: " + op.AddModifiers["--tipoDeRegistro"] + "'\n" + "'update --diaDePagamento: " + op.AddModifiers["--diaDePagamento"] + "'")
		return
	}

	switch op.Arguments[0] {
	case "--pessoa":
		op.UpdatePerson()
	case "--tipoDeRegistro":
		op.UpdateTypeOfRegister(op.Arguments[1:])
	case "--diaDePagamento":
		op.UpdatePayday(op.Arguments[1:])
	default:
		fmt.Println("Modificador inválido")
	}
}

func (op *Options) UpdatePerson() {
	if len(op.Modifiers) == 0 {
		fmt.Println("Atualizar o que de uma pessoa? Use os modificadores --id, --nucleo, --payment")
		return
	}

	id := op.Modifiers["--id"]
	nucleo := op.Modifiers["--nucleo"]
	payment := op.Modifiers["--payment"]

	err := executors.UpdateMembro(id, nucleo, payment)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Pessoa atualizada com sucesso")
}

func (op *Options) UpdateTypeOfRegister(args []string) {
	if len(args) < 2 {
		fmt.Println("Atualizar o que de um tipo de registro? Use os modificadores --id, --name, --description")
		return
	}

	id := args[0]
	PartyShare := args[1]

	err := executors.UpdateTipoDeRegistro(id, PartyShare)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Tipo de registro atualizado com sucesso")
}

func (op *Options) UpdatePayday(args []string) {
	if len(args) < 2 {
		fmt.Println("Atualizar o que de um nucleo? Args: Id Nucleo, Dia")
		return
	}

	id := args[0]
	day := args[1]

	err := executors.UpdatePayday(id, day)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Dia de pagamento atualizado com sucesso")
}
