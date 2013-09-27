package main

import (
	"strings"
)

var cmdRecord = &Command{
	Run:   runRecord,
	Usage: "record <command> [<args>]",
	Short: "Create, modify, or view records",
	Long: `
Create, modify, or view records

Usage:

  force record get <object> <id>

	force record create <object> [<fields>]

	force record update <object> <id> [<fields>]

	force record query <soql>

Examples:

  force record get User 00Ei0000000000

	force record create User Name:"David Dollar" Phone:0000000000

	force record update User 00Ei0000000000 State:GA

	force record query select id, name from user
`,
}

func runRecord(cmd *Command, args []string) {
	switch args[0] {
	case "get":
		runRecordGet(args[1:])
	case "create":
		runRecordCreate(args[1:])
	case "update":
		runRecordUpdate(args[1:])
	case "query":
		runRecordQuery(args[1:])
	default:
		ErrorAndExit("so such subcommand for record: %s", args[0])
	}
}

func runRecordGet(args []string) {
	force, _ := ActiveForce()
	if len(args) != 2 {
		ErrorAndExit("must specify type and id")
	}
	object, err := force.GetRecord(args[0], args[1])
	if err != nil {
		ErrorAndExit(err.Error())
	} else {
		DisplayForceRecord(object)
	}
}

func runRecordCreate(args []string) {
	ErrorAndExit("not implemented yet")
}

func runRecordUpdate(args []string) {
	ErrorAndExit("not implemented yet")
}

func runRecordQuery(args []string) {
	force, _ := ActiveForce()
	if len(args) < 1 {
		ErrorAndExit("must specify query")
	}
	query := strings.Join(args, " ")
	records, err := force.Query(query)
	if err != nil {
		ErrorAndExit(err.Error())
	} else {
		DisplayForceRecords(records)
	}
}
