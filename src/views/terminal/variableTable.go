package terminal

var variableTable map[string]string

func SetVariable(args []string) {
	if len(args) < 2 {
		return
	}

	if variableTable == nil {
		variableTable = make(map[string]string)
	}

	variableTable[args[0]] = args[1]
}

func GetVariable(name string) string {
	if variableTable == nil {
		return ""
	}

	return variableTable[name]
}