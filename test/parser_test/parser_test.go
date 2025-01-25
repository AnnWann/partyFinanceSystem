package parser_test

import (
	"reflect"
	"testing"

	"github.com/AnnWann/pstu_finance_system/src/parser"
)

func TestParseSucceding(t *testing.T) {
	type Expected struct {
		option    string
		modifiers map[string]string
		arguments []string
	}

	tests := []struct {
		signature int
		input    string
		expected Expected
	}{
		{1, "", Expected{"", nil, nil}},
		{2, "command", Expected{"command", nil, nil}},
		{3,"command --option", Expected{"command", map[string]string{"--option": ""}, nil}},
		{4, "command --option value", Expected{"command", map[string]string{"--option": ""}, []string{"value"}}},
		{5, "command --option value value2", Expected{"command", map[string]string{"--option": ""}, []string{"value", "value2"}}},
		{6, "command --option --option2", Expected{"command", map[string]string{"--option": "", "--option2": ""}, nil}},
		{7, "command --option --option2 value", Expected{"command", map[string]string{"--option": "", "--option2": "value"}, nil}},
		{8, "command --option --option2 value --option3 value2", Expected{"command", map[string]string{"--option": "", "--option2": "value", "--option3": "value2"}, nil}},
	}

	for _, test := range tests {
		option, modifiers, arguments, err := parser.Parse(test.input)
		if option != test.expected.option {
			t.Errorf("On signature %d, expected option %s, but got %s, with error message: %s", test.signature, test.expected.option, option, err)
		}
		if !reflect.DeepEqual(modifiers, test.expected.modifiers) {
			t.Errorf("On signature %d, expected modifiers %v, but got %v, with error message: %s", test.signature, test.expected.modifiers, modifiers, err)
		}
		if !reflect.DeepEqual(arguments, test.expected.arguments) {
			t.Errorf("On signature %d, expected arguments %v, but got %v, with error message: %s", test.signature, test.expected.arguments, arguments, err)
		}
	}
}

func TestParseFailing(t *testing.T) {
	tests := []string{
		"command --option value --option2",
		"command --option value --option2 value --option3",
		"command --option value --option2 value --option3 value",
	}

	for _, test := range tests {
		_, _, _, err := parser.Parse(test)
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	}
}
