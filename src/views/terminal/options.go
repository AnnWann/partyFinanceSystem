package terminal

var commands = map[string]string{
	"help":      "Mostrar esta ajuda",
	"exit":      "Sair do programa",
	"add":       "Adicionar algo",
	"get":       "Obter algo",
	"remove":    "Remover algo",
	"update":    "Atualizar algo",
	"promote":   "Promover um membro",
	"$":         "Definir uma variável. Use $ <nome> <valor>",
	"variables": "Mostrar todas as variáveis",
}

var addModifiers = map[string]string{
	"--membro":         "Adicionar uma membro. Args: Nome, NucleoId",
	"--registro":       "Adicionar um registro. Args: Dia, Mês, Ano, Nucleo, Tipo, Pagante, Cobrante, Quantidade, Valor, Descrição",
	"--tipoDeRegistro": "Adicionar um tipo de registro. Args: Nome, Nucleo, Descrição, PartilhaPartidária",
	"--nucleo":         "Adicionar um nucleo. Args: Nome, Cidade, Estado, Dia de pagamento",
	"--relatorio":      "Gerar um relatório. Args: NucleoId, Month, Year"}

var getModifiers = map[string]string{
	"--membro":         "Obter membros. Use --id, --nome, --cargo, --nucleo para filtrar a busca ou sem argumentos para obter todas as pessoas",
	"--registro":       "Obter registros. Use --nucleo --dia, --mes, --ano, --tipo, --pagante, --cobrante para filtrar",
	"--diaDePagamento": "Obter o dia de pagamento. Args: idNucleo",
	"--tipoDeRegistro": "Obter tipos de registro. Args: idNucleo",
	"--nucleo":         "Obter nucleos. Use --id, --nome, --cidade, --estado para filtrar",
	"--partido":        "Obter o partido.",
	"--relatorio":      "Obter relatórios. Use --nucleo, --mes, --ano para filtrar e --pdf para obter arquivo pdf"}

var updateModifiers = map[string]string{
	"--diaDePagamento": "Atualizar o dia de pagamento de um nucleo. Args: Id Nucleo, Dia",
	"--tipoDeregistro": "Atualizar o tipo de registro. Args: Id, PartilhaPartidária",
	"--pessoa":         "Atualizar o nucleo de uma pessoa. Use --id --contribuicao --nucleo"}

var removeModifiers = map[string]string{
	"--membro":         "Remover um membro. Args: id",
	"--tipoDeRegistro": "Remover um tipo de registro. Args: id",
	"--nucleo":         "Remover um nucleo. Args: id"}

type Options struct {
	Commands        map[string]string
	AddModifiers    map[string]string
	GetModifiers    map[string]string
	RemoveModifiers map[string]string
	UpdateModifiers map[string]string
	Option          string
	Modifiers       map[string]string
	Arguments       []string
}

func NewOptions(option string, modifiers map[string]string, arguments []string) Options {
	return Options{
		Commands:        commands,
		AddModifiers:    addModifiers,
		GetModifiers:    getModifiers,
		RemoveModifiers: removeModifiers,
		UpdateModifiers: updateModifiers,
		Option:          option,
		Modifiers:       modifiers,
		Arguments:       arguments}
}
