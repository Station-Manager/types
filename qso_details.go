package types

// QsoDetails represents the details of a QSO and is compatible with the ADI format.
type QsoDetails struct {
	AIndex      string `json:"a_index"`
	AntPath     string `json:"ant_path"` // ADIF, section II.B.1 - currently, we only use S and L
	Band        string `json:"band"`
	BandRx      string `json:"band_rx"` //in a split frequency QSO, the logging station's receiving band
	Comment     string `json:"comment"`
	ContestId   string `json:"contest_id"`
	Distance    string `json:"distance"` // km
	Freq        string `json:"freq" validate:"required"`
	FreqRx      string `json:"freq_rx"`
	Mode        string `json:"mode" validate:"required"`
	Submode     string `json:"submode"`
	Notes       string `json:"notes"` // information of interest to the logging station's operator
	QsoDate     string `json:"qso_date" validate:"required"`
	QsoDateOff  string `json:"qso_date_off"`
	QsoRandom   string `json:"qso_random"`
	QsoComplete string `json:"qso_complete"`
	RstRcvd     string `json:"rst_rcvd"`
	RstSent     string `json:"rst_sent"`
	RxPwr       string `json:"rx_pwr"` // the contacted station's transmitter power in Watts with a value greater than or equal to 0
	SRX         string `json:"srx"`    // contest QSO received serial number with a value greater than or equal to 0
	STX         string `json:"stx"`    // contest QSO transmitted serial number with a value greater than or equal to 0
	TimeOff     string `json:"time_off"`
	TimeOn      string `json:"time_on"`
	TxPwr       string `json:"tx_pwr"` // the logging station's power in Watts with a value greater than or equal to 0
}
