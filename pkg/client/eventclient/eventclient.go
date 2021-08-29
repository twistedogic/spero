package eventclient

import (
	"fmt"

	"github.com/twistedogic/spero/pkg/client/base"
)

const DefaultURL = "https://lsc.fn.sportradar.com/hkjc/en"

type Client struct {
	base.Client
	BaseURL string
}

func New(eventURL string) Client {
	client := base.New()
	return Client{client, eventURL}
}

func NewWithDefault() Client {
	return New(DefaultURL)
}

//https://lsc.fn.sportradar.com/hkjc/en/Asia:Shanghai/gismo/event_fullfeed/0/55
func (c Client) GetMatchFullFeed(offset int) ([]byte, error) {
	u := fmt.Sprintf("%s/Asia:Shanghai/gismo/event_fullfeed/%d/55", c.BaseURL, offset)
	return c.GetByte(u)
}

//https://lsc.fn.sportradar.com/hkjc/en/Etc:UTC/gismo/match_timelinedelta/14736623
func (c Client) GetMatchEvent(matchID int) ([]byte, error) {
	u := fmt.Sprintf("%s/Etc:UTC/gismo/match_timelinedelta/%d", c.BaseURL, matchID)
	return c.GetByte(u)
}

//https://lsc.fn.sportradar.com/hkjc/en/Etc:UTC/gismo/match_info/14736623
func (c Client) GetMatchInfo(matchID int) ([]byte, error) {
	u := fmt.Sprintf("%s/Etc:UTC/gismo/match_info/%d", c.BaseURL, matchID)
	return c.GetByte(u)
}

//https://lsc.fn.sportradar.com/hkjc/en/Etc:UTC/gismo/event_get/1
func (c Client) GetCurrentEvent(value interface{}) ([]byte, error) {
	u := fmt.Sprintf("%s/Etc:UTC/gismo/event_get/1", c.BaseURL)
	return c.GetByte(u)
}

//https://lsc.fn.sportradar.com/hkjc/en/Etc:UTC/gismo/match_details/14971653
func (c Client) GetMatchDetail(matchID int) ([]byte, error) {
	u := fmt.Sprintf("%s/Etc:UTC/gismo/match_details/%d", c.BaseURL, matchID)
	return c.GetByte(u)
}

//https://lsc.fn.sportradar.com/hkjc/en/Etc:UTC/gismo/match_situations/14971653
func (c Client) GetMatchSituation(matchID int) ([]byte, error) {
	u := fmt.Sprintf("%s/Etc:UTC/gismo/match_situations/%d", c.BaseURL, matchID)
	return c.GetByte(u)
}

//https://lsc.fn.sportradar.com/hkjc/en/Etc:UTC/gismo/match_detailsextended/14971653
func (c Client) GetMatchExtendedDetail(matchID int) ([]byte, error) {
	u := fmt.Sprintf("%s/Etc:UTC/gismo/match_detailsextended/%d", c.BaseURL, matchID)
	return c.GetByte(u)
}

//https://lsc.fn.sportradar.com/hkjc/en/Etc:UTC/gismo/match_timeline/14971653 <- matchID
func (c Client) GetMatchTimeline(matchID int) ([]byte, error) {
	u := fmt.Sprintf("%s/Etc:UTC/gismo/match_timeline/%d", c.BaseURL, matchID)
	return c.GetByte(u)
}

//https://lsc.fn.sportradar.com/hkjc/en/Etc:UTC/gismo/livescore_season_fixtures/55279 <- season fixtures
func (c Client) GetSeasonFixtures(seasonID int) ([]byte, error) {
	u := fmt.Sprintf("%s/Etc:UTC/gismo/livescore_season_fixtures/%d", c.BaseURL, seasonID)
	return c.GetByte(u)
}

//https://lsc.fn.sportradar.com/hkjc/en/Asia:Shanghai/gismo/bet_get/hkjc/0 for hkjc match id translation
func (c Client) GetBet(offset int) ([]byte, error) {
	u := fmt.Sprintf("%s/Asia:Shanghai/gismo/bet_get/hkjc/%d", c.BaseURL, offset)
	return c.GetByte(u)
}

//https://lsc.fn.sportradar.com/hkjc/en/Etc:UTC/gismo/match_bookmakerodds/14736623 (use hkjc instead)
func (c Client) GetMatchOdd(matchID int) ([]byte, error) {
	u := fmt.Sprintf("%s/Etc:UTC/gismo/match_bookmakerodds/%d", c.BaseURL, matchID)
	return c.GetByte(u)
}

//https://lsc.fn.sportradar.com/hkjc/en/Etc:UTC/gismo/stats_team_nextx/39/5
//https://lsc.fn.sportradar.com/hkjc/en/Etc:UTC/gismo/stats_team_lastx/39/5
//https://lsc.fn.sportradar.com/hkjc/en/Etc:UTC/gismo/stats_season_uniqueteamstats/54785
//https://lsc.fn.sportradar.com/hkjc/en/Etc:UTC/gismo/stats_team_versusrecent/14/34
