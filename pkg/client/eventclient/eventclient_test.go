package eventclient

import (
	"testing"

	"github.com/twistedogic/spero/pkg/schema/match"
)

func SetupClient() *Client {
	eventURL := "https://lsc.fn.sportradar.com/hkjc/en"
	return New(eventURL)
}

func TestGetMatchFullFeed(t *testing.T) {
	client := SetupClient()
	var fullfeed match.Fullfeed
	if err := client.GetMatchFullFeed(0, &fullfeed); err != nil {
		t.Error(err)
	}
	for _, match := range fullfeed.GetMatches() {
		if match.ID == 0 {
			t.Fail()
		}
	}
}
func TestGetMatchDetail(t *testing.T) {
	client := SetupClient()
	var detail match.Detail
	if err := client.GetMatchDetail(14970663, &detail); err != nil {
		t.Error(err)
	}
	t.Log(detail)
}
