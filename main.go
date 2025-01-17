package main

import (
	"fmt"
	"strings"

	"github.com/AnnWann/pstu_finance_system/lib/database"
	executors "github.com/AnnWann/pstu_finance_system/lib/mainHelpers"
	"github.com/AnnWann/pstu_finance_system/lib/pdf"
	"github.com/AnnWann/pstu_finance_system/lib/reportManager"
)

const file = "/home/uanadev/databases/pstu.db"

func main() {
	database.InitDB(file)

	var command string

	for {
		fmt.Print(">>> ")
		fmt.Scan(&command)

		splitCommand := strings.Fields(command)

		switch splitCommand[0] {
		case "exit":
			return 
		
		case "add": 
			switch splitCommand[1] {
				case "person":
					if len(splitCommand) != 4 {
						fmt.Println("Invalid number of arguments")
						break
					}

					if splitCommand[2] == "" || splitCommand[3] == "" {
						fmt.Println("Invalid arguments")
						break
					}

					err := executors.AddPerson(splitCommand[2], splitCommand[3])
					if err != nil {
						fmt.Println("Error adding person: ", err)
						break
					}
					fmt.Println("Person added")

				case "register":
					if len(splitCommand) != 12 {
						fmt.Println("Invalid number of arguments")
						break
					}

					if splitCommand[2] == "" || splitCommand[3] == "" ||
					 	splitCommand[4] == "" || splitCommand[5] == "" ||
						splitCommand[6] == "" || splitCommand[7] == "" ||
						splitCommand[8] == "" || splitCommand[9] == "" ||
						splitCommand[10] == "" || splitCommand[11] == "" {
						fmt.Println("Invalid arguments")
						break
					}

					Giver, err := database.GetPersonByName(splitCommand[5])
					if err != nil {
						fmt.Println("Giver not found")
						break
					}

					Receiver, err := database.GetPersonByName(splitCommand[6])
					if err != nil {
						fmt.Println("Receiver not found")
						break
					}

					err = executors.AddRegister(splitCommand[2], splitCommand[3], splitCommand[4], splitCommand[5], Giver.Id, Receiver.Id, splitCommand[7], splitCommand[8], splitCommand[9], splitCommand[10])
					if err != nil {
						fmt.Println("Error adding register: ", err)
						break
					}

					fmt.Println("Register added")

				case "help":
					fmt.Println("Available commands:")
					fmt.Println("add person <name> <role> - Add a person to the database")
					fmt.Println("add register <day> <month> <year> <type> <giver> <receiver> <amount> <value> <partyShare> <description> - Add a register to the database")
					fmt.Println("add help - Show this help message")

				default:
					fmt.Println("Invalid arguments. Try: add help")

			}
		case "report":
			switch splitCommand[1] {
				case "monthly":
					if len(splitCommand) != 4 {
						fmt.Println("Invalid number of arguments")
						break
					}

					reportExists, err:= database.GetReport(splitCommand[2] + "-" + splitCommand[3])
					if err != nil {
						fmt.Println("Error getting report: ", err)
						break
					}
					if reportExists.Id == "" {
						fmt.Println("Report already exists")
						break
					}

					tx, err := database.DB.Begin()
					if err != nil {
						tx.Rollback()
						fmt.Println("Error starting transaction: ", err)
						break
					}

					report, err := reportManager.GetMonthlyReport(splitCommand[2], splitCommand[3])
					if err != nil {
						tx.Rollback()
						fmt.Println("Error generating report: ", err)
						break
					}

					err = pdfMaker.MakePDFReport(report)
					if err != nil {
						tx.Rollback()
						fmt.Println("Error generating PDF: ", err)
						break
					}

					fmt.Println("Report generated")
					tx.Commit()

				case "yearly": 
					fmt.Println("Not implemented yet")
				
				case "help":
					fmt.Println("Available commands:")
					fmt.Println("report monthly <month> <year> - Generate a report for a specific month")
					fmt.Println("report yearly <year> - Generate a report for a specific year")
					fmt.Println("report help - Show this help message")
				
				default:
				fmt.Println("Invalid arguments. Try: report help")
			}
		
		case "help":
			fmt.Println("Available commands:")
			fmt.Println("add - Add a person or a register to the database")
			fmt.Println("report - Generate a report")
			fmt.Println("exit - Exit the program")
			fmt.Println("help - Show this help message")
			fmt.Println("For more information, type <command> help")
	
		default:
			fmt.Println("Invalid command. Try: help")
		}
	}
}