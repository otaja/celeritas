package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func setup(arg1, arg2 string) {
	if arg1 != "new" && arg1 != "version" && arg1 != "help" {
		err := godotenv.Load()
		if err != nil {
			exitGracefully(err)
		}

		path, err := os.Getwd()
		if err != nil {
			exitGracefully(err)
		}

		cel.RootPath = path
		cel.DB.DatabaseType = os.Getenv("DATABASE_TYPE")
	}
}

func getDSN() string {
	dbType := cel.DB.DatabaseType

	if dbType == "pgx" {
		dbType = "postgres"
	}

	if dbType == "postgres" {
		var dsn string
		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASS"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"))
		} else {
			dsn = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"))
		}
		return dsn
	}

	return "mysql://" + cel.BuildDSN()

}

func showHelp() {
	color.Blue(`Available commands:
	help                    - show the help page
	version                 - print the application version
	make migration <name>   - create a migrition file out of templates 
	make auth               - creates migrations for auth tables and code for auth / middleware
	migrate                 - runs all up migrations for the database
	migrate down            - runs all down migrations for the database
	migrate reset           - reset all migrations and run all up migrations again	
	make handler <name>     - create a stub handler in the handlers directory
	make model <name>       - creates a new model in the data directory
	make session			- creates a table in the database as a session store 
	make mail <name>        - create some E-Mail Templates for E-Mail Integration
	
	`)
}

func updateSourceFiles(path string, fi os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if fi.IsDir() {
		return nil
	}

	// only check go files
	matched, err := filepath.Match("*.go", fi.Name())
	if err != nil {
		return err
	}

	if matched {
		read, err := os.ReadFile(path)
		if err != nil {
			exitGracefully(err)
		}

		newContents := strings.Replace(string(read), "myapp", appURL, -1)

		err = os.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			exitGracefully(err)
		}
	}

	return nil
}

func updateSource() {
	// walk entire Project folders including subfolders
	err := filepath.Walk(".", updateSourceFiles)
	if err != nil {
		exitGracefully(err)
	}
}
