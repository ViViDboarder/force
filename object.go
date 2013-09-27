package main

import (
	"fmt"
)

var cmdObjects = &Command{
	Run:   runObjects,
	Usage: "objects",
	Short: "List force.com objects",
	Long: `
List force.com objects

Examples:

  force objects
`,
}

func runObjects(cmd *Command, args []string) {
	force, _ := ActiveForce()
	objects, err := force.ListObjects()
	if err != nil {
		ErrorAndExit(fmt.Sprintf("ERROR: %s\n", err))
	} else {
		DisplayStringSlice(objects)
	}
}
