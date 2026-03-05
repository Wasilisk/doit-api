package utils

import (
	"database/sql"
	"time"
)

func UnixToNullTime(unix *int64) sql.NullTime {
	if unix == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: time.Unix(*unix, 0), Valid: true}
}

func NullTimeToUnix(t sql.NullTime) *int64 {
	if !t.Valid {
		return nil
	}
	unix := t.Time.Unix()
	return &unix
}

func NullTimeToStringPtr(nt sql.NullTime) *string {
	if !nt.Valid {
		return nil
	}
	s := nt.Time.Format(time.RFC3339)
	return &s
}
