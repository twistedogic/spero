package oddscraper

import (
	"testing"

	"github.com/twistedogic/spero/pkg/schema/odd"
)

func TestOddScraper_GetOdd(t *testing.T) {
	type args struct {
		oddURL string
		t      odd.OddEnum
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base", args{"https://bet.hkjc.com/football/getJSON.aspx", odd.HAD}, false},
		{"err", args{"http://localhost", odd.HAD}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := New(tt.args.oddURL)
			_, err := o.GetOdd(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("OddScraper.GetOdd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
