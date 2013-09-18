package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/anthonybishopric/nac"
	"github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {

	datafile := flag.String("datafile", "/usr/local/var/nac/nac.db", "The name of the file nac will use to store program data")

	db, err := initDB(datafile)
	if err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
	}

	flag.Parse()

}

func initDB(datafile *string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", *datafile)
	if err != nil {
		return db, err
	}

	tx, err := db.Begin()
	if err != nil {
		return db, err
	}
	// setup tables

	return db, nil
}

func subcmd(name, signature, description string) *flag.FlagSet {
	flags := flag.NewFlagSet(name, flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Printf("\nUsage: docker %s %s\n\n%s\n\n", name, signature, description)
		flags.PrintDefaults()
	}
	return flags
}
