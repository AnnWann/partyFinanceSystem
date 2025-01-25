package terminal

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/AnnWann/pstu_finance_system/src/executors"
)

func (op *Options) Add() {
	if len(op.Arguments) == 0 {
		fmt.Println("Adicionar o que? Use 'add --person' ou 'add --register'")
		return
	}

	switch op.Arguments[0] {
	case "--person":
		op.AddPerson(op.Arguments[1:])
	case "--register":
		op.AddRegister(op.Arguments[1:])
	case "--typeOfRegister":
		op.AddTypeOfRegister(op.Arguments[1:])
	case "--nucleo":
		op.AddNucleo(op.Arguments[1:])
	case "--partido":
		op.AddPartido(op.Arguments[1:])
	case "--report":
		op.AddReport(op.Arguments[1:])
	default:
		fmt.Println("Modificador inválido")
	}
}

func (op *Options) AddPerson(Arguments []string) {
	if len(Arguments) == 0 {
		fmt.Println("Commando invalido, considere: " + op.Commands["add"] + " " + op.AddModifiers["--person"])
		return
	}

	nome := Arguments[0]
	nucleo := Arguments[1]
	id, err := executors.AddPerson(nome, nucleo)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s adicionada com sucesso com id: %s\n", nome, id)
}

func (op *Options) AddRegister(Arguments []string) {
	if len(Arguments) < 10 {
		fmt.Println("Commando invalido, considere: " + op.Commands["add"] + " " + op.AddModifiers["--register"])
		return
	}

	day := Arguments[0]
	month := Arguments[1]
	year := Arguments[2]
	nucleo := Arguments[3]
	tipo := Arguments[4]
	doador := Arguments[5]
	receptor := Arguments[6]
	quantidade := Arguments[7]
	valor := Arguments[8]
	descricao := Arguments[9]
	id, err := executors.AddRegister(day, month, year, nucleo, tipo, doador, receptor, quantidade, valor, descricao)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Registro adicionado com sucesso com id: %d\n", id)
}

func (op *Options) AddTypeOfRegister(Arguments []string) {
	if len(Arguments) < 4 {
		fmt.Println("Commando invalido, considere: " + op.Commands["add"] + " " + op.AddModifiers["--typeOfRegister"])
		return
	}

	nome := Arguments[0]
	nucleo := Arguments[1]
	descricao := Arguments[2]
	partilhaPartidaria := Arguments[3]

	err := executors.AddTypeOfRegister(nome, nucleo, descricao, partilhaPartidaria)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Tipo de registro adicionado com sucesso")
}

func (op *Options) AddNucleo(Arguments []string) {
	if len(Arguments) < 4 {
		fmt.Println("Commando invalido, considere: " + op.Commands["add"] + " " + op.AddModifiers["--nucleo"])
		return
	}

	nome := Arguments[0]
	cidade := Arguments[1]
	estado := Arguments[2]
	payday := Arguments[3]
	id, err := executors.AddNucleo(nome, cidade, estado, payday)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Nucleo adicionado com sucesso com id: %d\n", id)
}

func (op *Options) AddPartido(Arguments []string) {
	if len(Arguments) == 0 {
		fmt.Println("Commando invalido, considere: " + op.Commands["add"] + " " + op.AddModifiers["--partido"])
		return
	}

	err := executors.AddPartido(Arguments[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Partido adicionado com sucesso com id: 1")
}

func (op *Options) AddReport(Arguments []string) {
	if len(Arguments) == 0 {
		fmt.Println("Commando invalido, considere: " + op.Commands["add"] + " " + op.AddModifiers["--report"])
		return
	}

	var path_to_pdf string
	var err error
	var id string
	if len(Arguments) == 2 {
		nucleo := Arguments[0]
		year := Arguments[1]
		id, path_to_pdf, err = AddYearReport(nucleo, year)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		nucleo := Arguments[0]
		month := Arguments[1]
		year := Arguments[2]
		id, path_to_pdf, err = AddMonthReport(nucleo, month, year)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Printf("Relatório com id(%s) adicionado com sucesso em %s, abrindo arquivo..\n.", id, path_to_pdf)

	cmd := exec.Command("xdg-open", path_to_pdf)
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
}

func AddMonthReport(nucleo, month, year string) (string, string, error) {
	id, path_to_pdf, err := executors.AddMonthlyReport(nucleo, month, year)
	if err != nil {
		return "", "", err
	}

	return id, path_to_pdf, nil
}

func AddYearReport(nucleo, year string) (string, string, error) {
	return "", "", errors.New("not implemented")
	/* path_to_pdf, err := executors.AddYearlyReport(nucleo, year)
	if err != nil {
		return "", err
	}

	return path_to_pdf, nil */
}
