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
	return variableTable[name]
}

func DeleteVariable(name string) {
	delete(variableTable, name)
}

func ClearVariableTable() {
	variableTable = nil
	InitVariableTable()
}

func InitVariableTable() {
	variableTable = make(map[string]string)
	variableTable["externo"] = "0"
	variableTable["partido"] = "1"
	variableTable["nucleo_geral"] = "2"
	variableTable["aspirante"] = "-100"
	variableTable["militante"] = "-200"
	variableTable["dirigente"] = "-300"
	variableTable["dirigente_financeiro"] = "-400"
	variableTable["contribuicao"] = "-100"
	variableTable["jornal"] = "-200"
	variableTable["pagamento_partido"] = "-300"
	variableTable["gasto"] = "-400"
}
