package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/AnnWann/pstu_finance_system/src/views/terminal"
)

func Parse(str string) (string, map[string]string, []string, error) {
	var option string
	var modifiers = make(map[string]string)
	var arguments []string

	lexemes, err := getLexemes(str)
	if err != nil {
		return "", nil, nil, err
	}

	if len(lexemes) == 0 {
		return "", nil, nil, nil
	}

	if len(lexemes) < 2 {
		return lexemes[0], nil, nil, nil
	}

	err = checkSyntax(lexemes)
	if err != nil {
		return "", nil, nil, err
	}

	option = lexemes[0]
	firstModifier := lexemes[1]
	modifiers[firstModifier] = ""
	lexemesRest := lexemes[2:]

	for i := 0; i < len(lexemesRest); i++ {
		if strings.HasPrefix(lexemesRest[i], "--") {
			if i+1 < len(lexemesRest) && !strings.HasPrefix(lexemesRest[i+1], "--") {
				if strings.HasPrefix(lexemesRest[i+1], `$`) {
					modifiers[lexemesRest[i]] = terminal.GetVariable(strings.Trim(lexemesRest[i+1], `$`))
				} else {
					modifiers[lexemesRest[i]] = strings.Trim(lexemesRest[i+1], `"`)
					i++
				}
			} else {
				modifiers[lexemesRest[i]] = ""
			}
		} else {
			if strings.HasPrefix(lexemesRest[i], `$`) {
				arguments = append(arguments, terminal.GetVariable(strings.Trim(lexemesRest[i], `$`)))
			} else {
				arguments = append(arguments, strings.Trim(lexemesRest[i], `"`))
			}
		}
	}

	return option, modifiers, arguments, nil
}

func getLexemes(str string) ([]string, error) {
	re := regexp.MustCompile(`"[^"]*"|\S+`)
	splitStr := re.FindAllString(str, -1)

	if len(splitStr) == 0 {
		return nil, errors.New("no lexemes found")
	}

	return splitStr, nil
}

func checkSyntax(lexemes []string) error {
	if !expectOption(lexemes[0]) {
		return errors.New("esperava opção, mas recebeu algo diferente, em '" + lexemes[0] + " " + lexemes[1] + "'") 
	}

	if !expectModifier(lexemes[1]) {
		return errors.New("esperava modificador, mas recebeu algo diferente, em '" + lexemes[0] + " " + lexemes[1] + "'")
	}

	lexemes = lexemes[2:]

	if len(lexemes) == 0 {
		return nil
	}

	if expectArgument(lexemes[0]) {
		for i := 0; i < len(lexemes); i++ {
			if !expectArgument(lexemes[i]) {
				return errors.New("esperava argumento, mas recebeu algo diferente, em '" + lexemes[i-1] + " " + lexemes[i] + "'")
			}
		}
		return nil
	}

	var expect expectFn
	expect = expectModifier

	for i := 0; i < len(lexemes); i++ {
		if expect(lexemes[i]) {
			expect = flipExpect(expect)
			continue
		}
		return errors.New(
			"Syntax error: Recebeu " + Name(expect) +
				" mas esperava " + Name(flipExpect(expect)) +
				"em '" + lexemes[i-2] + " " + lexemes[i-1] + " " + lexemes[i] + "'")
	}
	return nil
}

type expectFn func(string) bool

func expectOption(lexemes string) bool {
	return !strings.HasPrefix(lexemes, "--")
}

func expectModifier(lexemes string) bool {
	return strings.HasPrefix(lexemes, "--")
}

func expectArgument(lexemes string) bool {
	return !strings.HasPrefix(lexemes, `--`)
}

func Name(fn func(string) bool) string {
	if getFunctionName(fn) == getFunctionName(expectOption) {
		return "Opção"
	}
	if getFunctionName(fn) == getFunctionName(expectModifier) {
		return "Modificador"
	}
	if getFunctionName(fn) == getFunctionName(expectArgument) {
		return "Argumento"
	}
	return "Desconhecido"
}

func flipExpect(expect expectFn) expectFn {
	if getFunctionName(expect) == getFunctionName(expectModifier) {
		return expectArgument
	}
	return expectModifier
}

func getFunctionName(i interface{}) string {
	return strings.TrimPrefix(strings.TrimSuffix(strings.TrimPrefix(strings.TrimSuffix(fmt.Sprintf("%v", i), "func(string) bool"), "func("), ")"), "main.")
}
