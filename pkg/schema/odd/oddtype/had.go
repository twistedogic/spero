package oddtype

type Had struct {
	H          string `json:"H" odd:"type=had"`
	D          string `json:"D" odd:"type=had"`
	A          string `json:"A" odd:"type=had"`
	ID         string `json:"ID" odd:"id"`
	POOLSTATUS string `json:"POOLSTATUS"`
	INPLAY     string `json:"INPLAY"`
	ALLUP      string `json:"ALLUP"`
	Cur        string `json:"Cur"`
}
