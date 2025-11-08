package types

// ContactedStation represents details of the station contacted during a QSO, and is compatible with the ADI format.
type ContactedStation struct {
	Address      string `json:"address"`
	Age          string `json:"age"`
	Altitude     string `json:"altitude"`
	Call         string `json:"call"`
	Cont         string `json:"cont"` // the contacted station's Continent
	ContactedOp  string `json:"contacted_op"`
	Country      string `json:"country"`
	CQZ          string `json:"cqz"`
	DXCC         string `json:"dxcc"`
	Email        string `json:"email"`
	EqCall       string `json:"eq_call"` // the contacted station's owner's callsign
	Gridsquare   string `json:"gridsquare"`
	Iota         string `json:"iota"`
	IotaIslandId string `json:"iota_island_id"`
	ITUZ         string `json:"ituz"`
	Lat          string `json:"lat"`
	Lon          string `json:"lon"`
	Name         string `json:"name"`
	QTH          string `json:"qth"`
	Rig          string `json:"rig"`
	Web          string `json:"web"`
	WwffRef      string `json:"wwff_ref"`
}
