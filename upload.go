package types

import "time"

type QsoUpload struct {
	ID            int64     `json:"id"`
	UploadedAt    time.Time `json:"uploaded_at"`
	NextAttemptAt time.Time `json:"next_attempt_at"`
	ModifiedAt    time.Time `json:"modified_at"`
	QsoID         int64     `json:"qso_id"`
	Service       string    `json:"service"`
	Status        string    `json:"status"`
	Attempts      int64     `json:"attempts"`
	LastError     string    `json:"last_error"`
	Qso           Qso
}
