package oddtype

type Hft struct {
	DA         string `json:"DA" odd:"type=hft"`
	HH         string `json:"HH" odd:"type=hft"`
	AH         string `json:"AH" odd:"type=hft"`
	AD         string `json:"AD" odd:"type=hft"`
	AA         string `json:"AA" odd:"type=hft"`
	DD         string `json:"DD" odd:"type=hft"`
	DH         string `json:"DH" odd:"type=hft"`
	HD         string `json:"HD" odd:"type=hft"`
	HA         string `json:"HA" odd:"type=hft"`
	ID         string `json:"ID" odd:"id"`
	POOLSTATUS string `json:"POOLSTATUS"`
	INPLAY     string `json:"INPLAY"`
	ALLUP      string `json:"ALLUP"`
	Cur        string `json:"Cur"`
}
