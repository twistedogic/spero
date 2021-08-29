package input

import (
	"log"

	"github.com/twistedogic/spero/pkg/input/odd"
)

func init() {
	if err := odd.Register(); err != nil {
		log.Fatal(err)
	}
}
