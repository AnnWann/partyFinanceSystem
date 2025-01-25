package terminal

var commands = map[string]string{
	"help":    "Mostrar esta ajuda",
	"exit":    "Sair do programa",
	"add":     "Adicionar algo",
	"get":     "Obter algo",
	"remove":  "Remover algo",
	"promote": "Promover um membro",
	"$":			 "Definir uma variável. Use $ <nome> <valor>",}

var addModifiers = map[string]string{
	"--person":         "Adicionar uma pessoa. Args: Nome, NucleoId",
	"--register":       "Adicionar um registro. Args: Dia, Mês, Ano, Nucleo, Tipo, Doador, Receptor, Valor, Descrição",
	"--typeOfRegister": "Adicionar um tipo de registro. Args: Nome, Nucleo, Descrição, PartilhaPartidária",
	"--nucleo":         "Adicionar um nucleo. Args: Nome, Cidade, Estado",
	"--partido":        "Adicionar um partido. Args: Nome",
	"--report":         "Gerar um relatório. Args: NucleoId, Month, Year"}

var getModifiers = map[string]string{
	"--person":         "Obter pessoas. Use --id, --name, --role, --nucleo para filtrar a busca ou sem argumentos para obter todas as pessoas",
	"--register":       "Obter registros. Use --nucleo --day, --month, --year, --type, --giver, --receiver para filtrar",
	"--payday":         "Obter o dia de pagamento. Args: idNucleo",
	"--typeOfRegister": "Obter tipos de registro. Args: idNucleo",
	"--nucleo":         "Obter nucleos. Use --id, --name, --city, --state para filtrar",
	"--partido":        "Obter o partido.",
	"--report":         "Obter relatórios. Use --nucleo, --month, --year para filtrar e --pdf para obter arquivo pdf"}

var removeModifiers = map[string]string{
	"--person":         "Remover uma pessoa. Args: id",
	"--typeOfRegister": "Remover um tipo de registro. Args: id",
	"--nucleo":         "Remover um nucleo. Args: id"}

type Options struct {
	Commands        map[string]string
	AddModifiers    map[string]string
	GetModifiers    map[string]string
	RemoveModifiers map[string]string
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
		Option:          option,
		Modifiers:       modifiers,
		Arguments:       arguments}
}
