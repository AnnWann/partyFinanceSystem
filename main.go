package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AnnWann/pstu_finance_system/lib/database"
	"github.com/AnnWann/pstu_finance_system/lib/executors"
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
		
			case "get":
				switch splitCommand[1] {
				case "person":
					if len(splitCommand) != 3 {
						fmt.Println("invalid number of arguments. Try 'get person --help'")
						continue
					}
					switch splitCommand[2] {
					case "--help":
						fmt.Println("get person --all - Gets all persons")
						fmt.Println("get person --members - Gets all members")
						fmt.Println("get person <id> - Gets one person")
						fmt.Println("get person --byName - Gets one person by name")
						continue
					case "--all":
						persons, err := executors.GetAllPersons()
						if err != nil {
							fmt.Println(err)
							continue
						}
						fmt.Print("Persons:\n")
						fmt.Print("id\tname\trole\n")
						for _, p := range persons {
							fmt.Printf("%s\t%s\t%s\n", p.Id, p.Name, p.Role)
						}
					case "--members":
						persons, err := executors.GetMembers()
						if err != nil {
							fmt.Println(err)
							continue
						}
						fmt.Print("Members:\n")
						fmt.Print("id\tname\trole\n")
						for _, p := range persons {
							fmt.Printf("%s\t%s\t%s\n", p.Id, p.Name, p.Role)
						}
					
					case "--byName": 
						var name string
						fmt.Print("Enter the name of the person: ")
						fmt.Scan(&name)
						person, err := executors.GetPerson(name)
						if err != nil {
							fmt.Println(err)
							continue
						}
						fmt.Printf("id: %s\nname: %s\nrole: %s\nMonthlyPayment: %f\nCredit: %f\n", person.Id, person.Name, person.Role, person.MonthlyPayment, person.Credit)

					default:
						person, err := executors.GetPerson(splitCommand[2])
						if err != nil {
							fmt.Println(err)
							continue
						}
						fmt.Printf("id: %s\nname: %s\nrole: %s\nMonthlyPayment: %f\nCredit: %f\n", person.Id, person.Name, person.Role, person.MonthlyPayment, person.Credit)
					}
				case "register":
					if len(splitCommand) == 3 {
						if splitCommand[2] == "--help" {
							fmt.Println("get register <id> - Gets one register")
							fmt.Println("get register --all - Gets all registers")
							fmt.Println("get register --byMonthAndYear <month> <year> - Gets all registers by month and year")
							fmt.Println("get register --byYear <year> - Gets all registers by year")
							continue
						} else if splitCommand[2] == "--all" {
							registers, err := executors.GetRegisters()
							if err != nil {
								fmt.Println(err)
								continue
							}
							fmt.Print("Registers:\n")
							fmt.Print("id\tdate\ttype\tgiver\treceiver\tamount\tvalue\tpartyShare\tdescription\n")
							for _, r := range registers {
								fmt.Printf("%s\t%s-%s-%s\t%s\t%s\t%s\t%d\t%f\t%f\t%s\n", r.Id, r.Day, r.Month, r.Year, r.Type, r.Giver, r.Receiver, r.Amount, r.Value, r.PartyShare, r.Description)
							}
						} else {
							register, err := executors.GetRegister(splitCommand[2])
							if err != nil {
								fmt.Println(err)
								continue
							}
							fmt.Printf("id: %s\ndate: %s-%s-%s\ntype: %s\ngiver: %s\nreceiver: %s\namount: %d\nvalue: %f\npartyShare: %f\ndescription: %s\n", register.Id, register.Day, register.Month, register.Year, register.Type, register.Giver, register.Receiver, register.Amount, register.Value, register.PartyShare, register.Description)
						}
					} else if len(splitCommand) == 4 {
						if splitCommand[2] == "--byYear" {
							registers, err := executors.GetRegistersByYear(splitCommand[3])
							if err != nil {
								fmt.Println(err)
								continue
							}
							fmt.Print("Registers:\n")
							fmt.Print("id\tdate\ttype\tgiver\treceiver\tamount\tvalue\tpartyShare\tdescription\n")
							for _, r := range registers {
								fmt.Printf("%s\t%s-%s-%s\t%s\t%s\t%s\t%d\t%f\t%f\t%s\n", r.Id, r.Day, r.Month, r.Year, r.Type, r.Giver, r.Receiver, r.Amount, r.Value, r.PartyShare, r.Description)
							}

						} else {	
							fmt.Println("invalid arguments. Try 'get register --help'")
						}

					} else if len(splitCommand) == 5 {
						if splitCommand[2] == "--byMonthAndYear" {
							registers, err := executors.GetRegisterByMonthAndYear(splitCommand[3], splitCommand[4])
							if err != nil {
								fmt.Println(err)
								continue
							}
							fmt.Print("Registers:\n")
							fmt.Print("id\tdate\ttype\tgiver\treceiver\tamount\tvalue\tpartyShare\tdescription\n")
							for _, r := range registers {
								fmt.Printf("%s\t%s-%s-%s\t%s\t%s\t%s\t%d\t%f\t%f\t%s\n", r.Id, r.Day, r.Month, r.Year, r.Type, r.Giver, r.Receiver, r.Amount, r.Value, r.PartyShare, r.Description)
							}
						} else {
							fmt.Println("invalid arguments. Try 'get register --help'")
						}
					} else {
						fmt.Println("invalid arguments. Try 'get register --help'")
					}
				
				case "report":
					if len(splitCommand) == 3 {
						if splitCommand[2] == "--help" {
							fmt.Println("get report <month> <year> - Gets one report")
							fmt.Println("get report <year> - Gets the report of the year (not implemented)")
							continue
						}

						_, err := strconv.Atoi(splitCommand[2])
						if err != nil {
							fmt.Println("invalid year")
							continue
						}  

						fmt.Print("not implemented yet")
					} else if len(splitCommand) == 4 {
						month, err := strconv.Atoi(splitCommand[2])
						if err != nil || month < 1 || month > 12 {
							fmt.Println("invalid month")
							continue
						}
						_, err = strconv.Atoi(splitCommand[3])
						if err != nil {
							fmt.Println("invalid year")
							continue
						}

						err = executors.GetMonthlyReport(splitCommand[2], splitCommand[3])
						if err != nil {
							fmt.Println(err)
							continue
						}
					} else {
						fmt.Println("invalid arguments. Try 'get report --help'")
					}
				
				case "payday":
					if len(splitCommand) == 2 {
						payday, err := executors.GetPayday()
						if err != nil {
							fmt.Println(err)
							continue
						}
						fmt.Printf("Payday: %s\n", payday)
					} else {
						fmt.Println("invalid arguments. Try 'get payday'")
					}
				
				case "help":
					fmt.Println("get person --all - Gets all persons")
					fmt.Println("get person --members - Gets all members")
					fmt.Println("get person <id> - Gets one person")
					fmt.Println("get register <id> - Gets one register")
					fmt.Println("get register --all - Gets all registers")
					fmt.Println("get register --byMonthAndYear <month> <year> - Gets all registers by month and year")
					fmt.Println("get register --byYear <year> - Gets all registers by year")
					fmt.Println("get report <month> <year> - Gets one report")
					fmt.Println("get report <year> - Gets the report of the year (not implemented)")
					fmt.Println("get payday - Gets the payday")
					continue

				default:
					fmt.Println("invalid arguments. Try 'get --help'")
			}	
			
			case "add":
				switch splitCommand[1] {
				case "person":
					if len(splitCommand) == 2 {
						var name string
						fmt.Print("Enter the name of the person: ")
						fmt.Scan(&name)
						var role string
						fmt.Print("select a role: \n1 - aspirant\n2 - militant\n3 - leader\n4 - financeLeader\n")
						var option int
						fmt.Scan(&option)
						switch option {
						case 1:
							role = "aspirant"
						case 2:
							role = "militant"
						case 3:
							role = "leader"
						case 4:
							role = "financeLeader"
						default:
							fmt.Println("invalid option")
							continue
						}

						err := executors.AddPerson(name, role)
						if err != nil {
							fmt.Println(err)
							continue
						}

						fmt.Println("Person added successfully")
					if len(splitCommand) == 3 {
						if splitCommand[2] == "--help" {
							fmt.Println("add person - Adds a person")
							continue
						}
					} else {
						fmt.Println("invalid arguments. Try 'add person --help'")
					} 
				}
				case "register":
					if len(splitCommand) == 11 {
						day := splitCommand[2]
						month := splitCommand[3]
						year := splitCommand[4]
						Type := splitCommand[5]
						giver := splitCommand[6]
						receiver := splitCommand[7]
						amount := splitCommand[8]
						value := splitCommand[9]
						partyShare := splitCommand[10]
						description := splitCommand[11]

						err := executors.AddRegister(day, month, year, Type, giver, receiver, amount, value, partyShare, description)
						if err != nil {
							fmt.Println(err)
							continue
						}
						fmt.Println("Register added successfully")
					} else if len(splitCommand) == 3 {
						if splitCommand[2] == "--help" {
							fmt.Println("add register <day> <month> <year> <type> <giver> <receiver> <amount> <value> <partyShare> <description> - Adds a register")
							continue
						}
					} else {
						fmt.Println("invalid arguments. Try 'add register --help'")
					}
				case "--help":
					fmt.Println("add person - Adds a person")
					fmt.Println("add register <day> <month> <year> <type> <giver> <receiver> <amount> <value> <partyShare> <description> - Adds a register")
					continue
				default:
					fmt.Println("invalid arguments. Try 'add --help'")
					continue
			}

			case "set":
				switch splitCommand[1] {
					case "payday":
						if len(splitCommand) == 3 {
							err := executors.SetPayday(splitCommand[2])
							if err != nil {
								fmt.Println(err)
								continue
							}
							fmt.Println("Payday set successfully")
						} else {
							fmt.Println("invalid arguments. Try 'set payday <day>'")
						}
					case "--help":
						fmt.Println("set payday <day> - Sets the payday")
						continue
					default:
						fmt.Println("invalid arguments. Try 'set --help'")
						continue
				}
			
			case "promote":
				if len(splitCommand) == 2 {
					if splitCommand[1] == "--help" {
						fmt.Println("promote <id> - Promotes a person to militant")
						fmt.Println("promote <promoteeId> <demoteeId> - swaps roles (demotee must be some kind of leader)")
					err := executors.PromoteToMilitant(splitCommand[1])
					if err != nil {
						fmt.Println(err)
						continue
					}
					fmt.Println("Person promoted successfully")
				} else if len(splitCommand) == 3 {
					err := executors.PromoteNewLeader(splitCommand[1], splitCommand[2])
					if err != nil {
						fmt.Println(err)
						continue
					}
					fmt.Println("Person promoted successfully")
				} else {
					fmt.Println("invalid arguments. Try 'promote --help'")
				}	
			}
			case "--help":
				fmt.Println("get person --all - Gets all persons")
				fmt.Println("get person --members - Gets all members")
				fmt.Println("get person <id> - Gets one person")
				fmt.Println("get person --byName - Gets one person by name")
				fmt.Println("get register <id> - Gets one register")
				fmt.Println("get register --all - Gets all registers")
				fmt.Println("get register <month> <year> - Gets all registers by month and year")
				fmt.Println("get register --byyear <year> - Gets all registers by year")
				fmt.Println("get report <month> <year> - Gets one report")
				fmt.Println("get report <year> - Gets the report of the year (not implemented)")
				fmt.Println("get payday - Gets the payday")
				fmt.Println("add person - Adds a person")
				fmt.Println("add register <day> <month> <year> <type> <giver> <receiver> <amount> <value> <partyShare> <description> - Adds a register")
				fmt.Println("set payday <day> - Sets the payday")
				fmt.Println("promote <id> - Promotes a person to militant")
				fmt.Println("promote <promoteeId> <demoteeId> - swaps roles (demotee must be some kind of leader")
				continue

			default:
				fmt.Println("invalid command. Try '--help'")
		}
	}
}