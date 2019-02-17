package poll

import (
	"reflect"
	"testing"
	"time"

	"github.com/twistedogic/spero/pkg/schema/odd"
	"github.com/twistedogic/spero/pkg/scraper/metascraper"
	"github.com/twistedogic/spero/pkg/scraper/oddscraper"
)

func TestPoller_Start(t *testing.T) {
	type fields struct {
		poller       Poller
		testInterval string
		interval     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"oddscraper", fields{oddscraper.New("https://bet.hkjc.com/football/getJSON.aspx", odd.HAD), "10s", "1s"}, false},
		{"metascraper", fields{metascraper.New("https://lsc.fn.sportradar.com/hkjc/en", 5), "10s", "1s"}, false},
		{"err", fields{oddscraper.New("http://localhost", odd.HAD), "10s", "1s"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testInterval, err := time.ParseDuration(tt.fields.testInterval)
			if err != nil {
				t.Error(err)
			}
			interval, err := time.ParseDuration(tt.fields.interval)
			if err != nil {
				t.Error(err)
			}
			p := New(tt.fields.poller, interval)
			ch := make(chan interface{})
			go func() {
				time.Sleep(testInterval)
				p.Stop()
			}()
			go func() {
				for v := range ch {
					if reflect.DeepEqual(v, odd.Outcome{}) {
						t.Error(v)
					}
				}
			}()
			if err := p.Start(ch); (err != nil) != tt.wantErr {
				t.Errorf("Poller.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
