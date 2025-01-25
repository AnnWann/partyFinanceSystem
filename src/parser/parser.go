package parser

import (
	"regexp"
	"strings"

	"github.com/AnnWann/pstu_finance_system/src/views/terminal"
)

func Parse(str string) (string, map[string]string, []string) {
	var option string
	var pairs = make(map[string]string)
	var arguments []string

	// Use regex to split the command while preserving quoted strings
	re := regexp.MustCompile(`"[^"]*"|\S+`)
	splitStr := re.FindAllString(str, -1)

	if len(splitStr) == 0 {
		return "", nil, nil
	}

	if len(splitStr) < 2 {
		return splitStr[0], nil, nil
	}

	option = splitStr[0]
	firstModifier := splitStr[1]
	pairs[firstModifier] = ""
	strRest := splitStr[1:]

	for i := 0; i < len(strRest); i++ {
		if strings.HasPrefix(strRest[i], "--") {
			if i+1 < len(strRest) && !strings.HasPrefix(strRest[i+1], "--") {
				if strings.HasPrefix(strRest[i+1], `$`) {
					pairs[strRest[i]] = terminal.GetVariable(strings.Trim(strRest[i+1], `$`))
				} else {
				pairs[strRest[i]] = strings.Trim(strRest[i+1], `"`)
				i++}
			} else {
				pairs[strRest[i]] = ""
			}
		} else {
			if strings.HasPrefix(strRest[i], `$`) {
				arguments = append(arguments, terminal.GetVariable(strings.Trim(strRest[i], `$`)))
			} else { 
				arguments = append(arguments, strings.Trim(strRest[i], `"`))
			}
		}
	}

	return option, pairs, arguments
}
