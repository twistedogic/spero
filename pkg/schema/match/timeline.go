package match

type Player struct {
	ID        int    `json:"_id"`
	Name      string `json:"name"`
	Fullname  string `json:"fullname"`
	Birthdate struct {
		Time     string `json:"time"`
		Date     string `json:"date"`
		Tz       string `json:"tz"`
		Tzoffset int    `json:"tzoffset"`
		Uts      int    `json:"uts"`
	} `json:"birthdate"`
	Nationality struct {
		ID          int    `json:"_id"`
		A2          string `json:"a2"`
		Name        string `json:"name"`
		A3          string `json:"a3"`
		Ioc         string `json:"ioc"`
		Continentid int    `json:"continentid"`
		Continent   string `json:"continent"`
		Population  int    `json:"population"`
	} `json:"nationality"`
	Position struct {
		Type      string `json:"_type"`
		Name      string `json:"name"`
		Shortname string `json:"shortname"`
		Abbr      string `json:"abbr"`
	} `json:"position"`
	Primarypositiontype interface{} `json:"primarypositiontype"`
	Haslogo             bool        `json:"haslogo"`
}

type Event struct {
	ID         int    `json:"_id"`
	Sid        int    `json:"_sid"`
	Rcid       int    `json:"_rcid"`
	Tid        int    `json:"_tid"`
	Dc         bool   `json:"_dc"`
	Uts        int    `json:"uts"`
	UpdatedUts int    `json:"updated_uts"`
	Type       string `json:"type"`
	Matchid    int    `json:"matchid"`
	Disabled   int    `json:"disabled"`
	Time       int    `json:"time"`
	Seconds    int    `json:"seconds"`
	Name       string `json:"name"`
	Injurytime int    `json:"injurytime"`
	Team       string `json:"team"`
	Result     struct {
		Home int `json:"home"`
		Away int `json:"away"`
	} `json:"result,omitempty"`
	X           int      `json:"X,omitempty"`
	Y           int      `json:"Y,omitempty"`
	Scorer      Player   `json:"scorer,omitempty"`
	Header      bool     `json:"header,omitempty"`
	Owngoal     bool     `json:"owngoal,omitempty"`
	Penalty     bool     `json:"penalty,omitempty"`
	Assists     []Player `json:"assists,omitempty"`
	Card        string   `json:"card,omitempty"`
	Periodscore struct {
		Home int `json:"home"`
		Away int `json:"away"`
	} `json:"periodscore,omitempty"`
	MatchStatus struct {
		Doc  string `json:"_doc"`
		ID   int    `json:"_id"`
		Name string `json:"name"`
	} `json:"matchStatus,omitempty"`
	Playerout    Player `json:"playerout,omitempty"`
	Playerin     Player `json:"playerin,omitempty"`
	Shirtnumbers struct {
		In  string `json:"in"`
		Out string `json:"out"`
	} `json:"shirtnumbers,omitempty"`
	Situation   string `json:"situation,omitempty"`
	Coordinates []struct {
		Team string `json:"team"`
		X    int    `json:"X"`
		Y    int    `json:"Y"`
	} `json:"coordinates,omitempty"`
}
