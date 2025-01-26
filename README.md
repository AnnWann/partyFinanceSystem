# PARTY FINANCE SYSTEM / SISTEMA FINANCEIRO DE PARTIDO

This is an app made to handle the financial needs of a political party that's organized into smaller cores. \
<span style="color:#808b96">.Isto é um aplicativo feito para organizar as necessidades de um partido político organizado em nucleos.</span>


It stores user's data and financial registers, builds monthly and yearly reports and prints a PDF file of said reports. \
<span style="color:#808b96">Suas funções incluem armazenar dados de usuários e registros financeiros, construir relatórios mensais e anuais e imprimir os relatórios em arquivos PDF.</span>

## Instalation / Instalação

Download source from/<span style="color:#808b96">Instale fonte de</span>: https://github.com/AnnWann/partyCoreFinantialSystem \
Download binary from/<span style="color:#808b96">Instale binário de</span>: (not implemented) \
Use browser version at/<span style="color:#808b96">Use versão no navegador em</span>: (not implemented) \

## Usage / Uso

You use this app by two means: cli tool or browser interface (not implemented). \
<span style="color:#808b96">Existem dois meios de usar esse aplicativo: a ferramenta command-line e a interface de browser (não implementada)</span>

### Cli Tool / Ferramenta Command-line

**Pre-Requisites/Pré-requisitos**

There is no pre-requisite if you're running from the binary. \
<span style="color:#808b96">Não existe nenhum pré-requisito se você está rodando do binário. </span>

For the source version, you will need `make` and `go` installed, as well as their own pre-requisites. \
<span style="color:#808b96">Para rodar a versão de fonte, você vai precisar de `make` e `go` instalado, como também os seus respectivos pré-requisitos. </span>

**To run from source/<span style="color:#808b96">Para rodar da fonte/span>**

    $ make all

That will generate an executable file and execute it in your terminal. \
<span style="color:#808b96">Isto vai gerar um arquivo executavél e executá-lo dentro de seu terminal. </span>

**For this version**, the binary program will by default generate a .db file inside of `./database` and store the PDFs inside of `./reports`. \
<span style="color:#808b96">**Nesta versão**, por padrão o programa binário irá criar um arquivo .db dentro de `./database` e armazenar os PDFs dentro de `./reports`.</span>

If you are using the source version, you can change where the database and PDF are stored by using `$ echo -e "PDF_FOLDER=<your-pdf-folder>\nDB_PATH=<your-database-folder>" > .env` \
<span style="color:#808b96">Se você está usando a versão fonte, você pode alterar onde o banco e os PDFs vão ser armazenados usando `$ echo -e "PDF_FOLDER=<seu-diretorio-pdf>\nDB_PATH=<seu-diretorio-banco>" > env` </span>

To know more about the usage of the CLI tool, use `$ help` or access ["docs/CLI_TOOL_USAGE.md"](./docs/CLI_TOOL_USAGE.md) \
<span style="color:#808b96">Para saber mais sobre o uso do terminal, use `$ help` ou acesse ["docs/CLI_TOOL_USAGE.md"](./docs/CLI_TOOL_USAGE.md) </span>

### Browser interface

Go to ... \
Acesse ...









