package oddtype

type Fha struct {
	H          string `json:"H" odd:"type=fha"`
	D          string `json:"D" odd:"type=fha"`
	A          string `json:"A" odd:"type=fha"`
	ID         string `json:"ID" odd:"id"`
	POOLSTATUS string `json:"POOLSTATUS"`
	INPLAY     string `json:"INPLAY"`
	ALLUP      string `json:"ALLUP"`
	Cur        string `json:"Cur"`
}
