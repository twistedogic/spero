package commands

import (
	"fmt"
  "time"

  "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/twistedogic/spero/pkg/storage"
	"github.com/twistedogic/spero/cmd/service/metricservice"
	"github.com/twistedogic/spero/cmd/job/oddjob"
	"github.com/twistedogic/spero/cmd/job/metajob"
	"github.com/twistedogic/spero/cmd/job"
	"github.com/twistedogic/spero/cmd/service"
	"github.com/urfave/cli"
)

func Run(c *cli.Context) error {
  betTypes := c.StringSlice("type")
	oddPeriod, err := time.ParseDuration(fmt.Sprintf("%ds", oddInterval))
	if err != nil {
		return err
	}
	metaPeriod, err := time.ParseDuration(fmt.Sprintf("%ds", metaInterval))
	if err != nil {
		return err
  }
  db, err := gorm.Open("postgres", addr) 
  if err != nil {
    return err
  }
  store := storage.New(db)
	jobInstance := job.New(
    oddjob.New(oddPeriod, oddURL, betTypes, store), 
    metajob.New(metaPeriod, eventURL, rate, store),
  )
  serviceInstance := service.New(port, metricservice.New())
  go serviceInstance.Start()
  go jobInstance.Start()
	return Graceful(jobInstance, serviceInstance)
}
