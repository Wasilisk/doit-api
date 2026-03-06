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

func NullableTime(t sql.NullTime) *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}
