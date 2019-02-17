package oddscraper

import (
	"github.com/twistedogic/spero/pkg/client/oddclient"
	"github.com/twistedogic/spero/pkg/metric"
	"github.com/twistedogic/spero/pkg/schema/odd"
)

type OddScraper struct {
	*oddclient.Client
	BetType []odd.OddEnum
}

func New(oddURL string, types ...odd.OddEnum) *OddScraper {
	return &OddScraper{
		oddclient.New(oddURL),
		types,
	}
}

func (o *OddScraper) GetOdd(t odd.OddEnum) ([]odd.Match, error) {
	var oddTable odd.OddTable
	err := o.GetOddTable(t.Key(), &oddTable)
	metric.UpdatePollMetric(t, err)
	return oddTable.Matches, err
}

func (o *OddScraper) Poll(ch chan interface{}) error {
	for _, t := range o.BetType {
		matches, err := o.GetOdd(t)
		if err != nil {
			return err
		}
		for _, match := range matches {
			ch <- match
		}
	}
	return nil
}
