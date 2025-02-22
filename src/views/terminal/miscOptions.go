package terminal

import (
	"fmt"
	"os"
)

func (op *Options) Help() {
	for k, v := range op.Commands {
		fmt.Printf("%s: %s\n", k, v)
		switch k {
		case "add":
			for mk, mv := range op.AddModifiers {
				fmt.Printf("  %s: %s\n", mk, mv)
			}
		case "get":
			for mk, mv := range op.GetModifiers {
				fmt.Printf("  %s: %s\n", mk, mv)
			}
		case "update":
			for mk, mv := range op.UpdateModifiers {
				fmt.Printf("  %s: %s\n", mk, mv)
			}
		case "remove":
			for mk, mv := range op.RemoveModifiers {
				fmt.Printf("  %s: %s\n", mk, mv)
			}
		}
	}
	fmt.Println("Use $<nome_variavel> para chamar uma variável. Exemplo: $nome\nColoque informações entre \"\" para enviar argumentos compostos. Ex: \"nome sobrenome\"")
}
func Exit() {
	fmt.Println("Saindo...")
	os.Exit(0)
}
