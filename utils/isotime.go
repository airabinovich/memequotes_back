package utils

import (
	"fmt"
	"time"
)

// ISO8601Time represents a date time that can be marshalled and unmarshalled following the ISO 8601 convention.
type ISO8601Time time.Time

// Layout represent the formatter to be used for time parsing. Go doesn't provide a time parser that support millis
// time.RFC3339 (no millis) or time.RFC3339nano (nano seconds)
const Layout = "2006-01-02T15:04:05.000Z07:00"

// MarshalJSON creates a JSON representation for a date, following the ISO 8601 convention.
func (t *ISO8601Time) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf(`"%s"`, time.Time(*t).Format(Layout))
	return []byte(stamp), nil
}

// UnmarshalJSON parses a JSON representation, following the ISO 8601 convention, to a time element.
func (t *ISO8601Time) UnmarshalJSON(b []byte) error {
	tt, err := time.Parse(fmt.Sprintf(`"%s"`, Layout), string(b))
	if err == nil {
		*t = ISO8601Time(tt)
	}
	return err
}

// String returns the string representation using ISO8601 format
func (t *ISO8601Time) String() string {
	return time.Time(*t).Format(Layout)
}
