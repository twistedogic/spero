package oddtype

type Hha struct {
	A          string `json:"A" odd:"type=hha"`
	H          string `json:"H" odd:"type=hha"`
	D          string `json:"D" odd:"type=hha"`
	ID         string `json:"ID" odd:"id"`
	POOLSTATUS string `json:"POOLSTATUS"`
	INPLAY     string `json:"INPLAY"`
	ALLUP      string `json:"ALLUP"`
	HG         string `json:"HG"`
	AG         string `json:"AG"`
	Cur        string `json:"Cur"`
}
