package schema

import (
	"encoding/json"
	"strconv"
	"time"

	pb "github.com/twistedogic/spero/pb"
)

const (
	TIMESTAMP_FORAMT = "2006-01-02T15:04:05-07:00"
	DATE_FORMAT      = "2006-01-02-07:00"
)

type Team struct {
	TeamID     string `json:"teamID"`
	TeamNameCH string `json:"teamNameCH"`
	TeamNameEN string `json:"teamNameEN"`
}

func (t Team) ToProto() *pb.Team {
	id, _ := strconv.ParseInt(t.TeamID, 10, 64)
	return &pb.Team{
		Id:     id,
		NameEn: t.TeamNameEN,
		NameCn: t.TeamNameCH,
	}
}

type League struct {
	LeagueID        string `json:"leagueID"`
	LeagueShortName string `json:"leagueShortName"`
	LeagueNameCH    string `json:"leagueNameCH"`
	LeagueNameEN    string `json:"leagueNameEN"`
}

func (l League) ToProto() *pb.League {
	id, _ := strconv.ParseInt(l.LeagueID, 10, 64)
	return &pb.League{
		Id:        id,
		Shortname: l.LeagueShortName,
		NameCn:    l.LeagueNameCH,
		NameEn:    l.LeagueNameEN,
	}
}

type Match struct {
	MatchID           string `json:"matchID" odd:"id"`
	MatchIDinofficial string `json:"matchIDinofficial"`
	MatchNum          string `json:"matchNum"`
	MatchDate         string `json:"matchDate"`
	MatchDay          string `json:"matchDay"`
	Coupon            struct {
		CouponID        string `json:"couponID"`
		CouponShortName string `json:"couponShortName"`
		CouponNameCH    string `json:"couponNameCH"`
		CouponNameEN    string `json:"couponNameEN"`
	} `json:"coupon"`
	League            League `json:"league,omitempty" odd:"league"`
	HomeTeam          Team   `json:"homeTeam" odd:"home"`
	AwayTeam          Team   `json:"awayTeam" odd:"away"`
	MatchStatus       string `json:"matchStatus,omitempty"`
	MatchTime         string `json:"matchTime" odd:"match_time"`
	Statuslastupdated string `json:"statuslastupdated,omitempty"`
	Inplaydelay       string `json:"inplaydelay,omitempty"`
	LiveEvent         struct {
		IlcLiveDisplay  bool          `json:"ilcLiveDisplay"`
		HasLiveInfo     bool          `json:"hasLiveInfo"`
		IsIncomplete    bool          `json:"isIncomplete"`
		MatchIDbetradar string        `json:"matchIDbetradar"`
		Matchstate      string        `json:"matchstate"`
		StateTS         string        `json:"stateTS"`
		Liveevent       []interface{} `json:"liveevent"`
	} `json:"liveEvent,omitempty"`
	Accumulatedscore []struct {
		Periodvalue  string `json:"periodvalue"`
		Periodstatus string `json:"periodstatus"`
		Home         string `json:"home"`
		Away         string `json:"away"`
	} `json:"accumulatedscore,omitempty"`
	Livescore struct {
		Home string `json:"home"`
		Away string `json:"away"`
	} `json:"livescore,omitempty"`
	Cornerresult string `json:"cornerresult,omitempty"`
	Cur          string `json:"Cur"`
	IsDef        string `json:"isDef,omitempty"`
	HasWebTV     bool   `json:"hasWebTV"`
	Hadodds      struct {
		H          string `json:"H" odd:"value,type=had"`
		D          string `json:"D" odd:"value,type=had"`
		A          string `json:"A" odd:"value,type=had"`
		ID         string `json:"ID"`
		POOLSTATUS string `json:"POOLSTATUS"`
		INPLAY     string `json:"INPLAY"`
		ALLUP      string `json:"ALLUP"`
		Cur        string `json:"Cur"`
	} `json:"hadodds,omitempty"`
	Fhaodds struct {
		H          string `json:"H" odd:"value,type=fha"`
		D          string `json:"D" odd:"value,type=fha"`
		A          string `json:"A" odd:"value,type=fha"`
		ID         string `json:"ID"`
		POOLSTATUS string `json:"POOLSTATUS"`
		INPLAY     string `json:"INPLAY"`
		ALLUP      string `json:"ALLUP"`
		Cur        string `json:"Cur"`
	} `json:"fhaodds,omitempty"`
	Crsodds struct {
		S0402      string `json:"S0402"`
		S0205      string `json:"S0205"`
		SM1MH      string `json:"SM1MH"`
		S0200      string `json:"S0200"`
		S0301      string `json:"S0301"`
		S0002      string `json:"S0002"`
		SM1MA      string `json:"SM1MA"`
		S0303      string `json:"S0303"`
		S0501      string `json:"S0501"`
		S0001      string `json:"S0001"`
		S0005      string `json:"S0005"`
		S0000      string `json:"S0000"`
		S0401      string `json:"S0401"`
		S0100      string `json:"S0100"`
		S0105      string `json:"S0105"`
		S0101      string `json:"S0101"`
		S0400      string `json:"S0400"`
		S0502      string `json:"S0502"`
		S0003      string `json:"S0003"`
		SM1MD      string `json:"SM1MD"`
		S0300      string `json:"S0300"`
		S0102      string `json:"S0102"`
		S0202      string `json:"S0202"`
		S0500      string `json:"S0500"`
		S0302      string `json:"S0302"`
		S0204      string `json:"S0204"`
		S0201      string `json:"S0201"`
		S0004      string `json:"S0004"`
		S0203      string `json:"S0203"`
		S0104      string `json:"S0104"`
		S0103      string `json:"S0103"`
		ID         string `json:"ID"`
		POOLSTATUS string `json:"POOLSTATUS"`
		INPLAY     string `json:"INPLAY"`
		ALLUP      string `json:"ALLUP"`
		Cur        string `json:"Cur"`
	} `json:"crsodds,omitempty"`
	Fcsodds struct {
		S0100      string `json:"S0100"`
		S0203      string `json:"S0203"`
		SM1MA      string `json:"SM1MA"`
		S0204      string `json:"S0204"`
		S0202      string `json:"S0202"`
		S0402      string `json:"S0402"`
		S0004      string `json:"S0004"`
		S0303      string `json:"S0303"`
		S0104      string `json:"S0104"`
		S0401      string `json:"S0401"`
		S0000      string `json:"S0000"`
		S0201      string `json:"S0201"`
		S0103      string `json:"S0103"`
		SM1MH      string `json:"SM1MH"`
		S0101      string `json:"S0101"`
		S0500      string `json:"S0500"`
		S0300      string `json:"S0300"`
		S0301      string `json:"S0301"`
		S0501      string `json:"S0501"`
		S0400      string `json:"S0400"`
		S0502      string `json:"S0502"`
		S0002      string `json:"S0002"`
		S0205      string `json:"S0205"`
		S0302      string `json:"S0302"`
		S0003      string `json:"S0003"`
		S0005      string `json:"S0005"`
		S0102      string `json:"S0102"`
		SM1MD      string `json:"SM1MD"`
		S0200      string `json:"S0200"`
		S0001      string `json:"S0001"`
		S0105      string `json:"S0105"`
		ID         string `json:"ID"`
		POOLSTATUS string `json:"POOLSTATUS"`
		INPLAY     string `json:"INPLAY"`
		ALLUP      string `json:"ALLUP"`
		Cur        string `json:"Cur"`
	} `json:"fcsodds,omitempty"`
	Ooeodds struct {
		O          string `json:"O" odd:"value,type=ooe"`
		E          string `json:"E" odd:"value,type=ooe"`
		ID         string `json:"ID"`
		POOLSTATUS string `json:"POOLSTATUS"`
		INPLAY     string `json:"INPLAY"`
		ALLUP      string `json:"ALLUP"`
		Cur        string `json:"Cur"`
	} `json:"ooeodds,omitempty"`
	Ttgodds struct {
		P5         string `json:"P5"`
		P1         string `json:"P1"`
		P2         string `json:"P2"`
		P6         string `json:"P6"`
		P3         string `json:"P3"`
		P4         string `json:"P4"`
		P0         string `json:"P0"`
		M7         string `json:"M7"`
		ID         string `json:"ID"`
		POOLSTATUS string `json:"POOLSTATUS"`
		INPLAY     string `json:"INPLAY"`
		ALLUP      string `json:"ALLUP"`
		Cur        string `json:"Cur"`
	} `json:"ttgodds,omitempty"`
	Hftodds struct {
		DA         string `json:"DA"`
		HH         string `json:"HH"`
		AH         string `json:"AH"`
		AD         string `json:"AD"`
		AA         string `json:"AA"`
		DD         string `json:"DD"`
		DH         string `json:"DH"`
		HD         string `json:"HD"`
		HA         string `json:"HA"`
		ID         string `json:"ID"`
		POOLSTATUS string `json:"POOLSTATUS"`
		INPLAY     string `json:"INPLAY"`
		ALLUP      string `json:"ALLUP"`
		Cur        string `json:"Cur"`
	} `json:"hftodds,omitempty"`
	Hhaodds struct {
		A          string `json:"A"`
		H          string `json:"H"`
		D          string `json:"D"`
		ID         string `json:"ID"`
		POOLSTATUS string `json:"POOLSTATUS"`
		INPLAY     string `json:"INPLAY"`
		ALLUP      string `json:"ALLUP"`
		HG         string `json:"HG"`
		AG         string `json:"AG"`
		Cur        string `json:"Cur"`
	} `json:"hhaodds,omitempty"`
	Hdcodds struct {
		A          string `json:"A"`
		H          string `json:"H"`
		ID         string `json:"ID"`
		POOLSTATUS string `json:"POOLSTATUS"`
		INPLAY     string `json:"INPLAY"`
		ALLUP      string `json:"ALLUP"`
		HG         string `json:"HG"`
		AG         string `json:"AG"`
		Cur        string `json:"Cur"`
	} `json:"hdcodds,omitempty"`
	Hilodds struct {
		LINELIST []struct {
			LINENUM    string `json:"LINENUM"`
			MAINLINE   string `json:"MAINLINE"`
			LINESTATUS string `json:"LINESTATUS"`
			LINEORDER  string `json:"LINEORDER"`
			LINE       string `json:"LINE"`
			H          string `json:"H"`
			L          string `json:"L"`
		} `json:"LINELIST"`
		ID         string `json:"ID"`
		POOLSTATUS string `json:"POOLSTATUS"`
		INPLAY     string `json:"INPLAY"`
		ALLUP      string `json:"ALLUP"`
		Cur        string `json:"Cur"`
	} `json:"hilodds,omitempty"`
	Fhlodds struct {
		LINELIST []struct {
			LINENUM    string `json:"LINENUM"`
			MAINLINE   string `json:"MAINLINE"`
			LINESTATUS string `json:"LINESTATUS"`
			LINEORDER  string `json:"LINEORDER"`
			LINE       string `json:"LINE"`
			L          string `json:"L"`
			H          string `json:"H"`
		} `json:"LINELIST"`
		ID         string `json:"ID"`
		POOLSTATUS string `json:"POOLSTATUS"`
		INPLAY     string `json:"INPLAY"`
		ALLUP      string `json:"ALLUP"`
		Cur        string `json:"Cur"`
	} `json:"fhlodds,omitempty"`
	Chlodds struct {
		LINELIST []struct {
			LINENUM    string `json:"LINENUM"`
			MAINLINE   string `json:"MAINLINE"`
			LINESTATUS string `json:"LINESTATUS"`
			LINEORDER  string `json:"LINEORDER"`
			LINE       string `json:"LINE"`
			H          string `json:"H"`
			L          string `json:"L"`
		} `json:"LINELIST"`
		ID         string `json:"ID"`
		POOLSTATUS string `json:"POOLSTATUS"`
		INPLAY     string `json:"INPLAY"`
		ALLUP      string `json:"ALLUP"`
		Cur        string `json:"Cur"`
	} `json:"chlodds,omitempty"`
	HasExtraTimePools bool `json:"hasExtraTimePools"`
	Results           struct {
	} `json:"results"`
	DefinedPools []string `json:"definedPools"`
	InplayPools  []string `json:"inplayPools"`
}

func (m Match) GetLastUpdate() (time.Time, error) {
	return time.Parse(TIMESTAMP_FORAMT, m.Statuslastupdated)
}

func (m Match) GetMatchDate() (time.Time, error) {
	return time.Parse(DATE_FORMAT, m.MatchDate)
}

func (m Match) String() string {
	b, _ := json.MarshalIndent(m, "", "  ")
	return string(b)
}

type Matches []Match

func (m Matches) Len() int      { return len(m) }
func (m Matches) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

type ByLastUpdate struct{ Matches }

func (s ByLastUpdate) Less(i, j int) bool {
	a, _ := s.Matches[i].GetLastUpdate()
	b, _ := s.Matches[j].GetLastUpdate()
	return a.After(b)
}
