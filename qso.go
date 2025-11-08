package types

type Qso struct {
	ID int64 `json:"id"`
	QsoDetails
	ContactedStation
	LoggingStation
}
