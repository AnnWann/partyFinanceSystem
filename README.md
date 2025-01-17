# PSTU FINANCE SYSTEM

This is an app made to handle the finantial necessities of a political party's core.

It stores user's data and finantial registers, builds monthly and yearly reports (not implemeted yet) and prints a PDF file of said reports.

## Instalation

Download source from https://github.com/AnnWann/partyCoreFinantialSystem
Download binary from (not implemented)
Use browser version at (not implemented)

## Usage

You use this app by two means: cli tool or browser interface (not implemented). 

### Cli Tool

To run:
  $ ./pstu_finance_system

If running from source:
  $ go build
  $ ./pstu_finance_system

The commands are:
  $ get person --all - Gets all persons
  $ get person --members - Gets all members
  $ get person \<id\> - Gets one person
  $ get person --byName - Gets one person by name
  $ get register \<id\> - Gets one register
  $ get register --all - Gets all registers
  $ get register --byMonthAndYear \<month\> \<year\> - Gets all registers by month and year
  $ get register --byYear \<year\> - Gets all registers by year
  $ get report \<month\> \<year\> - Gets one report
  $ get report \<year\> - Gets the report of the year (not implemented)
  $ get payday - Gets the payday
  $ add person - Adds a person
  $ add register \<day\> \<month\> \<year\> \<type\> \<giver\> \<receiver\> \<amount\> \<value\> \<partyShare\> \<description\> - Adds a register
  $ set payday \<day\> - Sets the payday
  $ promote \<id\> - Promotes a person to militant
  $ promote \<promoteeId\> \<demoteeId\> - swaps roles (demotee must be some kind of leader)

To understand what each variable is, see at [CLI VARIABLES](#cli-variables)

### Browser interface

The browser interface is very intuitive. 

## CLI VARIABLES

person.name string 
  the person's name

person.role string
  the person's role
can be:
  "party", "core", "leader", "financesLeader", "militant", "aspirant", "outsider"
contraints:
  There must be a "party", "core" and "outsider" in the database
  There is only one of "party", "core" and "outsider" in the database.
  There can only be one "leader" and "financesLeader"
  There must be at least 2 militants to add an aspirant
  If the database has 1 militant, this militant will be promoted to leader
  If the database has 2 militants, the second will be promoted to "financesLeader"
  If a "Leader" or "financesLeader" is demoted, a militant must be promoted with the corresponding role

register.day string
  day of the month (number) that the register was made. 

register.month string
  month (number) that the register was made. 

register.year string
  year (number) that the register was made. 

register.type string
  the type of register
can be: 
  "payment", "journal", "other", "expense"

register.giver string
  id of the person who paid for it

register.receiver string
  id of the person who paid for it

register.amount int
  the amount of items involved in the register

register.value float
  the monetary value of the register

register.partyShare float
  the share of the register that belongs to the party

register.description string
  a brief description of the register










