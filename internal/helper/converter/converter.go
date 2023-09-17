package converter

import (
	"database/sql"
	"time"
)

func NullStringToString(nullStr sql.NullString) string {
	if nullStr.Valid {
		return nullStr.String
	}
	return ""
}

func StringToNullString(str string) sql.NullString {
	if str == "" {
		return sql.NullString{Valid: false, String: ""}
	}
	return sql.NullString{Valid: true, String: str}
}

func TimeToNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: !t.IsZero()}
}

func NullTimeToTime(nt sql.NullTime) time.Time {
	if nt.Valid {
		return nt.Time
	}
	return time.Time{}
}

func Int64ToNullInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: true}
}

func NullInt64ToInt64(ni sql.NullInt64) int64 {
	if ni.Valid {
		return ni.Int64
	}
	return 0 // or any other default value you prefer
}

func StringToTime(dateStr string, layout string) (time.Time, error) {
	// Parse the string using the provided layout
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func TimeToString(t time.Time, layout string) string {
	return t.Format(layout)
}

func ConvertCategory(s string) uint64 {
	if s == "Makanan" {
		return 1
	} else if s == "Pakaian" {
		return 2
	}
	return 3
}
