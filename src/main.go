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
	godotenv.Load()
	file := os.Getenv("DB_FILE")
	database.GetDB().InitDB(file)
	reader := bufio.NewReader(os.Stdin)
	var input string

	for {
		input, _ = reader.ReadString('\n')
		command, modifiers, arguments, err := parser.Parse(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		terminal.HandleOption(command, modifiers, arguments)
	}

}

