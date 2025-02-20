# PARTY FINANCE SYSTEM

This is an app made to handle the financial needs of a political party that's organized into smaller cores.
It stores user's data and financial registers, builds monthly and yearly reports and prints a PDF file of said reports.

## Instalation

Pre-requisites: golang, make and sqlite3.

Download source from: https://github.com/AnnWann/partyCoreFinantialSystem.

## Usage

### Cli Tool

1. Create a file at the directory of your choice named <something>.db
2. Create a .env file with the following variables:

        PDF_FOLDER=/where/you/want/to/store/the/reports
        DB_PATH=/where/your/database/file/is.db
        PARTIDO=your_Party_Name
        UNIDOC_KEY=your_UNIDOC_API_Key

3. run `make`
4. you are now in the program interface, if all the steps above were followed, you should have configured your database correctly and can now use the app's function. To learn more about the usage, input `help` or access ["docs/CLI_TOOL_USAGE.md"](./docs/CLI_TOOL_USAGE.md).











