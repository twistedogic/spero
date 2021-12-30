package input

import (
	"log"

	"github.com/twistedogic/spero/pkg/input/odd"
	"github.com/twistedogic/spero/pkg/input/result"
)

func init() {
	if err := odd.Register(); err != nil {
		log.Fatal(err)
	}
	if err := result.Register(); err != nil {
		log.Fatal(err)
	}
}
