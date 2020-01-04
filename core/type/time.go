package types

import (
	"time"

	"github.com/Sharykhin/go-payments/core"
)

type (
	Time time.Time
)

func TimeNow() Time {
	return Time(time.Now().UTC())
}

// MarshalJSON implements json.Marshaler interface
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(t).Format(`"` + core.ISO8601 + `"`)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (t *Time) UnmarshalJSON(data []byte) error {
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.Parse("2006-01-02T15:04:05Z", string(data))
	if err != nil {
		*t = Time{}
	} else {
		*t = Time(tt)
	}
	return nil
}
