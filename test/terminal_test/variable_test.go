package test

import (
	"testing"

	"github.com/AnnWann/pstu_finance_system/src/parser"
	"github.com/AnnWann/pstu_finance_system/src/views/terminal"
)

func TestVariableSucceding(t *testing.T) {
	names := []string{"test", "test2", "test3"}
	values := []string{"value", "value2", "value3"}

	defer terminal.ClearVariableTable()

	for i, name := range names {
		command := "$ " + name + " " + values[i]
		option, modifiers, arguments, err := parser.Parse(command)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		terminal.HandleOption(option, modifiers, arguments)
	}

	for i, name := range names {
		if terminal.GetVariable(name) != values[i] {
			t.Errorf("Expected %q but got %q", values[i], terminal.GetVariable(name))
		}
	}
}

func TestVariableFailing(t *testing.T) {
	names := []string{"test", "test2", "--modifier"}
	commands := []string{
		"$",
		"$ test",
		"$ test2 --modifier",
		"$ --modifier value",
	}

	defer terminal.ClearVariableTable()

	for _, command := range commands {
		option, modifiers, arguments, _ := parser.Parse(command)

		terminal.HandleOption(option, modifiers, arguments)
	}

	for i, name := range names {
		value := terminal.GetVariable(name)
		if value != "" {
			t.Errorf("At signature %d, expected no value but got %q", i+1, value)
		}
	}
}

func TestParsingWithVariable(t *testing.T) {
	option, modifiers, arguments, err := parser.Parse("$ test value")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	terminal.HandleOption(option, modifiers, arguments)

	command := "option --modifier $test"
	_, _, arguments, err = parser.Parse(command)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if arguments[0] != "value" {
		t.Errorf("Expected value but got %q", arguments[0])
	}
}