package metascraper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/twistedogic/spero/pkg/client/eventclient"
	"github.com/twistedogic/spero/pkg/metric"
	"github.com/twistedogic/spero/pkg/schema/match"
	"github.com/twistedogic/spero/pkg/schema/sportradar"
)

type MetaScraper struct {
	client        *eventclient.Client
	ratePerSecond int
}

func New(eventURL string, rate int) *MetaScraper {
	return &MetaScraper{eventclient.New(eventURL), rate}
}

func (m *MetaScraper) GetMatchDetail(id int) (match.Detail, error) {
	var detail match.Detail
	err := m.client.GetMatchDetail(id, &detail)
	metric.UpdateScrapeMetric(detail, err)
	if err != nil {
		log.Printf("MatchDetail %d %v", id, err)
		return detail, err
	}
	return detail, nil
}

func (m *MetaScraper) GetMatchEvents(id int) ([]match.Event, error) {
	out := make([]match.Event, 0)
	var container sportradar.Container
	err := m.client.GetMatchTimeline(id, &container)
	metric.UpdateScrapeMetric(out, err)
	if err != nil {
		log.Printf("Events %d %v", id, err)
		return out, err
	}
	for _, data := range container.GetData() {
		var events struct {
			Events []match.Event `json:"events"`
		}
		if err := json.Unmarshal(data, &events); err != nil {
			return out, err
		}
		out = append(out, events.Events...)
	}
	return out, nil
}

func (m *MetaScraper) GetMatchSituations(id int) ([]match.Situation, error) {
	out := make([]match.Situation, 0)
	var container sportradar.Container
	err := m.client.GetMatchSituation(id, &container)
	metric.UpdateScrapeMetric(out, err)
	if err != nil {
		log.Printf("Situations %d %v", id, err)
		return out, err
	}
	for _, data := range container.GetData() {
		var situations struct {
			Situations []match.Situation `json:"situations"`
		}
		if err := json.Unmarshal(data, &situations); err != nil {
			return out, err
		}
		out = append(out, situations.Situations...)
	}
	return out, nil
}

func (m *MetaScraper) GetMatchMeta(data match.Match) (match.MatchData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id := data.ID
	var wg sync.WaitGroup
	ch := make(chan interface{})
	errCh := make(chan error)
	done := make(chan bool)
	wg.Add(3)
	go func() {
		defer wg.Done()
		if detail, err := m.GetMatchDetail(id); err != nil {
			errCh <- err
		} else {
			ch <- detail
		}
	}()
	go func() {
		defer wg.Done()
		if events, err := m.GetMatchEvents(id); err != nil {
			errCh <- err
		} else {
			ch <- events
		}
	}()
	go func() {
		defer wg.Done()
		if situations, err := m.GetMatchSituations(id); err != nil {
			errCh <- err
		} else {
			ch <- situations
		}
	}()
	go func() {
		wg.Wait()
		close(ch)
		close(errCh)
		done <- true
	}()
	out := match.MatchData{Match: data}
	for {
		select {
		case data := <-ch:
			switch v := data.(type) {
			case match.Detail:
				out.Detail = v
			case []match.Event:
				out.Events = v
			case []match.Situation:
				out.Situations = v
			}
		case err := <-errCh:
			return out, err
		case <-ctx.Done():
			close(ch)
			close(errCh)
			return out, fmt.Errorf("time out")
		case <-done:
			break
		}
	}
	return out, nil
}

func (m *MetaScraper) GetMatches(offset int) ([]match.MatchData, error) {
	var wg sync.WaitGroup
	interval := time.Duration(1000/m.ratePerSecond) * time.Millisecond
	ticker := time.Tick(interval)
	out := make([]match.MatchData, 0)
	var fullfeed match.Fullfeed
	if err := m.client.GetMatchFullFeed(0, &fullfeed); err != nil {
		return out, err
	}
	ch := make(chan match.MatchData)
	errCh := make(chan error, 1)
	done := make(chan bool)
	for _, rawMatch := range fullfeed.GetMatches() {
		<-ticker
		wg.Add(1)
		go func(matchEntry match.Match) {
			defer wg.Done()
			if data, err := m.GetMatchMeta(matchEntry); err != nil {
				errCh <- err
			} else {
				ch <- data
			}
		}(rawMatch)

	}
	go func() {
		wg.Wait()
		close(ch)
		close(errCh)
		done <- true
	}()
	for {
		select {
		case data := <-ch:
			out = append(out, data)
		case err := <-errCh:
			return out, err
		case <-done:
			break
		}
	}
	return out, nil
}

func (m *MetaScraper) Poll(ch chan interface{}) error {
	matches, err := m.GetMatches(-1)
	if err != nil {
		return err
	}
	for _, match := range matches {
		ch <- match
	}
	return nil
}
