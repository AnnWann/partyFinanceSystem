package terminal

import "fmt"

func HandleOption(option string, modifiers map[string]string, arguments []string) {
	op := NewOptions(option, modifiers, arguments)
	switch op.Option {
	case "help":
		op.Help()
	case "exit":
		Exit()
	case "add":
		op.Add()
	case "get":
		op.Get()
	case "remove":
		op.Remove()
	case "promote":
		op.Promote()
	case "$":
		SetVariable(op.Arguments)
	case "variables":
		ShowVariables()
	default:
		op.Help()
	}
}

func ShowVariables() {
	for key, value := range variableTable {
		fmt.Printf("%s: %s\n", key, value)
	}
}
