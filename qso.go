package types

type Qso struct {
	ID int64 `json:"id"`
	QsoDetails
	ContactedStation
	LoggingStation
}

// QsoDetails represents the details of a QSO, and is compatible with the ADI format.
type QsoDetails struct {
	AIndex      string `json:"a_index"`
	AntPath     string `json:"ant_path"` // ADIF, section II.B.1 - currently, we only use S and L
	Band        string `json:"band"`
	BandRx      string `json:"band_rx"` //in a split frequency QSO, the logging station's receiving band
	Comment     string `json:"comment"`
	ContestId   string `json:"contest_id"`
	Distance    string `json:"distance"` // km
	Freq        string `json:"freq"`
	FreqRx      string `json:"freq_rx"`
	Mode        string `json:"mode"`
	Submode     string `json:"submode"`
	MySig       string `json:"my_sig"`
	MySigInfo   string `json:"my_sig_info"`
	Notes       string `json:"notes"` // information of interest to the logging station's operator
	QsoDate     string `json:"qso_date"`
	QsoDateOff  string `json:"qso_date_off"`
	QsoRandom   string `json:"qso_random"`
	QsoComplete string `json:"qso_complete"`
	RstRcvd     string `json:"rst_rcvd"`
	RstSent     string `json:"rst_sent"`
	RxPwr       string `json:"rx_pwr"`   // the contacted station's transmitter power in Watts with a value greater than or equal to 0
	Sig         string `json:"sig"`      // the name of the contacted station's special activity or interest group
	SigInfo     string `json:"sig_info"` // information associated with the contacted station's activity or interest group
	SRX         string `json:"srx"`      // contest QSO received serial number with a value greater than or equal to 0
	STX         string `json:"stx"`      // contest QSO transmitted serial number with a value greater than or equal to 0
	TimeOff     string `json:"time_off"`
	TimeOn      string `json:"time_on"`
	TxPwr       string `json:"tx_pwr"` // the logging station's power in Watts with a value greater than or equal to 0
}

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

// LoggingStation represents the details of the station logging the QSO, including location, equipment, and operator info.
// It is compatible with the ADI format.
type LoggingStation struct {
	AntennaAzimuth  string `json:"ant_az"` // the bearing from the logging station to the contacted station
	MyAltitude      string `json:"my_altitude"`
	MyAntenna       string `json:"my_antenna"`
	MyCity          string `json:"my_city"`
	MyCountry       string `json:"my_country"`
	MyCqZone        string `json:"my_cq_zone"`
	MyDXCC          string `json:"my_dxcc"`
	MyGridsquare    string `json:"my_gridsquare"`
	MyIota          string `json:"my_iota"`
	MyIotaIslandID  string `json:"my_iota_island_id"`
	MyITUZone       string `json:"my_itu_zone"`
	MyLat           string `json:"my_lat"`
	MyLon           string `json:"my_lon"`
	MyMorseKeyInfo  string `json:"my_morse_key_info"`
	MyMorseKeyType  string `json:"my_morse_key_type"`
	MyName          string `json:"my_name"`
	MyPostalCode    string `json:"my_postal_code"`
	MyRig           string `json:"my_rig"`
	MyStreet        string `json:"my_street"`
	MyWwffRef       string `json:"my_wwff_ref"`
	Operator        string `json:"operator"` // the logging operator's callsign if STATION_CALLSIGN is absent, OPERATOR shall be treated as both the logging station's callsign and the logging operator's callsign
	OwnerCallsign   string `json:"owner_callsign"`
	StationCallsign string `json:"station_callsign"`
}
