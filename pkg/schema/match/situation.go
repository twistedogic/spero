package match

type Situation struct {
	Doc        string `json:"_doc"`
	ID         int    `json:"_id"`
	Matchid    int    `json:"matchid"`
	Uts        int    `json:"uts"`
	Team       string `json:"team"`
	Situation  string `json:"situation"`
	X          int    `json:"x"`
	Y          int    `json:"y"`
	Time       int    `json:"time"`
	Injurytime int    `json:"injurytime"`
}
