package main

import (
	"log"
	"os"

	"github.com/twistedogic/spero/cmd/commands"
)

func main() {
	app := commands.RunCLI()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
