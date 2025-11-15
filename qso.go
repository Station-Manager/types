package types

type Qso struct {
	ID int64 `json:"id"`

	// LogbookID represents the foreign key to the logbook associated with a QSO entry.
	// Every QSO entry MUST have a logbook associated with it.
	LogbookID int64 `json:"logbook_id" validate:"required"`

	/*
		All the below fields are compatible with the ADI format and are populated by the adapter.
		The only exception to this is the ID field, which is required by database functions.
	*/
	QsoDetails
	ContactedStation
	LoggingStation
}
