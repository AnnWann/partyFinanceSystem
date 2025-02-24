package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AnnWann/pstu_finance_system/src/database"
	"github.com/AnnWann/pstu_finance_system/src/parser"
	"github.com/AnnWann/pstu_finance_system/src/views/terminal"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		fmt.Println("Exiting...")
		os.Exit(1)
	}

	file := os.Getenv("DB_PATH")
	err = database.GetDB().InitDB(file)
	if err != nil {
		fmt.Println("Error initializing database: ", err)
		fmt.Println("Exiting...")
		os.Exit(1)
	}

	terminal.InitVariableTable()
	reader := bufio.NewReader(os.Stdin)
	var input string

	nomePartido := os.Getenv("PARTIDO")
	fmt.Printf("Sistema Financeiro do %s\nUse \"help\" para aprender como usá-lo", nomePartido)
	for {
		print("\n\n>>> ")
		input, _ = reader.ReadString('\n')
		command, modifiers, arguments, err := parser.Parse(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		terminal.HandleOption(command, modifiers, arguments)
	}

}
