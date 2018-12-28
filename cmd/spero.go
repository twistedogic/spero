package main

import (
	"fmt"
  "time"
  "os"
  "log"
  "github.com/urfave/cli"

	"github.com/prometheus/client_golang/prometheus"
  "github.com/grpc-ecosystem/go-grpc-prometheus"
  "github.com/twistedogic/spero/cmd/server"
  "github.com/twistedogic/spero/cmd/poll"
  "github.com/twistedogic/spero/cmd/commands"
	"github.com/twistedogic/spero/pkg/metric"
	pb "github.com/twistedogic/spero/pb"
)

var (
  registry = prometheus.NewRegistry()
  grpcMetrics = grpc_prometheus.NewServerMetrics()
)
var promPort, rpcPort, interval int
var asProm bool
var baseURL string
var betTypes []string

func init() {
  registry.MustRegister(
    grpcMetrics,
    metric.OddMetric,
    metric.ClientMetric,
  )
}

func Run(c *cli.Context) error {
  if types := c.StringSlice("type"); len(types) > 0 {
    betTypes = types
  }
  period, err := time.ParseDuration(fmt.Sprintf("%ds",interval))
  if err != nil {
    return err
  }
  out := make(chan *pb.Odd, 1)
  poll := poll.New(baseURL, period)
  metricServer := server.NewMetricServer(registry, promPort)
  rpcServer, err := server.NewRpcServer(out, grpcMetrics, rpcPort)
  if err != nil {
    return err
  }
  go poll.Start(out, betTypes...)
  go metricServer.Start()
  go rpcServer.Start()
  return commands.Graceful(metricServer,rpcServer,poll)
}

func main() {
  app := cli.NewApp()
  app.Name = "spero"
  app.Flags = []cli.Flag{
    cli.BoolFlag{
      Name:"prom, m",
      Hidden: true,
      Usage:"export odd as promethus metric",
      Destination: &asProm,
    },
    cli.StringFlag{
      Name:"url, u",
      Value: "https://bet.hkjc.com/football/getJSON.aspx",
      Usage:"target url",
      Destination: &baseURL,
    },
    cli.IntFlag{
      Name:"interval, i",
      Value: 10,
      Usage:"poll interval in seconds",
      Destination: &interval,
    },
    cli.StringSliceFlag{
      Name:"type, t",
      Usage: "type to poll",
      Value: &cli.StringSlice{"HAD","FHA","CRS","FCS","FTS","OOE","TTG","HFT","HHA","HDC","HIL","FHL"},
    },
    cli.IntFlag{
      Name:"http, o",
      Value: 9090,
      Usage:"http port",
      Destination: &promPort,
    },
    cli.IntFlag{
      Name:"rpc, p",
      Value: 8080,
      Usage:"rpc port",
      Destination: &rpcPort,
    },
  }
  app.Action = Run
  
  if err := app.Run(os.Args); err != nil {
    log.Fatal(err)
  }
}
