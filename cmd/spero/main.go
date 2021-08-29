package main

import (
	"context"

	_ "github.com/twistedogic/spero/pkg/input"

	"github.com/Jeffail/benthos/v3/public/service"
)

func main() {
	service.RunCLI(context.Background())
}
