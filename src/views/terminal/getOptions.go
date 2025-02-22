package terminal

import (
	"fmt"

	"github.com/AnnWann/pstu_finance_system/src/executors"
)

func (op *Options) Get() {
	if len(op.Modifiers) == 0 {
		fmt.Println("Obter o que? Use:\n'--membro: " + op.Modifiers["--membro"] + "'\n'--registro: " + op.Modifiers["--registro"] + "'\n'--relatorio: " + op.Modifiers["--relatorio"] + "'\n'--diaDePagamento: " + op.Modifiers["--diaDePagamento"] + "'\n'--tipoDeRegistro: " + op.Modifiers["--tipoDeRegistro"] + "'\n'--nucleo: " + op.Modifiers["--nucleo"] + "'\n'-partido: " + op.Modifiers["--partido"] + "'")
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
	case "--membro":
		op.GetPerson(modifiersRest)
	case "--registro":
		op.GetRegister(modifiersRest)
	case "--relatorio":
		op.GetReport(modifiersRest)
	case "--diaDePagamento":
		if len(op.Arguments) == 0 {
			fmt.Println("Commando invalido, considere: " + op.Commands["get"] + " " + op.GetModifiers["--payday"])
			return
		}
		op.GetPayday(op.Arguments[0])
	case "--tipoDeRegistro":
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

	persons, err := executors.GetMembro(modifiers)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, person := range persons {
		fmt.Printf("dados da pessoa:\nid: %d\nnome: %s\ncargo: %d\nnucleo: %d\ncontribuição mensal: %.2f\ncredito: %.2f\n", person.ID, person.Nome, person.Cargo, person.Nucleo, person.Contribuicao_mensal, person.Credito)
	}
}

func (op *Options) GetRegister(modifiers map[string]string) {

	register, err := executors.GetRegistro(modifiers)
	if err != nil {
		fmt.Println(err)
	}

	for _, reg := range register {
		fmt.Printf("dados do registro:\nid: %d\ndia: %s\nmês: %s\nano: %s\nquantidade: %d\nvalor: %.2f\ntipo: %d\nnucleo: %d\npago_por: %d\ncobrado_por: %d\ndescrição: %s\n", reg.ID, reg.Dia, reg.Mes, reg.Ano, reg.Quantidade, reg.Valor, reg.Tipo, reg.Nucleo, reg.Pagante, reg.Cobrante, reg.Descricao)
	}
}

func (op *Options) GetReport(modifiers map[string]string) {

	relatorios, err := executors.GetRelatorioMensal(modifiers)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("urls dos relatórios mensais:")
	for _, relatorio := range relatorios {
		fmt.Printf("%s/%s: %s\n", relatorio.Mes, relatorio.Ano, relatorio.Url)
	}
}

func (op *Options) GetTypeOfRegister(nucleoId string) {

	typeOfRegister, err := executors.GetTipoDeRegistro(nucleoId)
	if err != nil {
		fmt.Println(err)
	}

	for _, tReg := range typeOfRegister {
		fmt.Printf("dados do tipo de registro:\nid: %d\nnome: %s\nnucleo: %d\ndescrição: %s\npartilha partidária: %.2f\n", tReg.ID, tReg.Nome, tReg.Nucleo, tReg.Descricao, tReg.Parcela_partidaria)
	}
}

func (op *Options) GetNucleo(modifiers map[string]string) {

	nucleos, err := executors.GetNucleo(modifiers)
	if err != nil {
		fmt.Println(err)
	}

	for _, nucleo := range nucleos {
		fmt.Printf("dados do nucleo:\nid: %d\nnome: %s\ncidade: %s\nestado: %s\nreserva: %.2f\ndia de pagamento: %s\n", nucleo.ID, nucleo.Nome, nucleo.Cidade, nucleo.Estado, nucleo.Reserva, nucleo.Dia_de_Pagamento)
	}
}

func (op *Options) GetPartido() {
	partido, err := executors.GetPartido()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("dados do partido:\nid: %d\nnome: %s\nreserva: %.2f", partido.ID, partido.Nome, partido.Reserva)
}

func (op *Options) GetPayday(nucleoId string) {
	payday, err := executors.GetPayday(nucleoId)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("dia de pagamento: %s", payday)
}
