package commands

import (
	"github.com/urfave/cli"
)

var port, oddInterval, metaInterval, rate int
var oddURL, eventURL, addr string

func RunCLI() *cli.App {
	app := cli.NewApp()
	app.Name = "spero"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "odd_url, o",
			Value:       "https://bet.hkjc.com/football/getJSON.aspx",
			Usage:       "target odd url",
			Destination: &oddURL,
		},
		cli.StringFlag{
			Name:        "event_url, e",
			Value:       "https://lsc.fn.sportradar.com/hkjc/en",
			Usage:       "target event url",
			Destination: &eventURL,
		},
		cli.StringFlag{
			Name:        "addr, a",
			Value:       "postgresql://maxroach@database:26257/spero?sslmode=disable",
			Usage:       "postgresql connection string",
			Destination: &addr,
		},
		cli.IntFlag{
			Name:        "odd_interval, i",
			Value:       10,
			Usage:       "odd poll interval in seconds",
			Destination: &oddInterval,
		},
		cli.IntFlag{
			Name:        "meta_interval, m",
			Value:       86400,
			Usage:       "meta poll interval in seconds",
			Destination: &metaInterval,
		},
		cli.StringSliceFlag{
			Name:  "type, t",
			Usage: "type to poll",
			Value: &cli.StringSlice{"HAD", "FHA", "HFT", "HHA"},
		},
		cli.IntFlag{
			Name:        "port, p",
			Value:       3000,
			Usage:       "http port",
			Destination: &port,
		},
	}
	app.Action = Run
	return app
}
