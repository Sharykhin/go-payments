package entity

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"reflect"
	"time"
)

type (
	// NullString represents string values that can be null
	NullString sql.NullString
	// NullTime represents datetime that can be null
	NullTime struct {
		Valid bool
		Time  time.Time
	}
)

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

// MarshalJSON implements json.Marshaler interface
func (nt NullTime) MarshalJSON() ([]byte, error) {
	if nt.Valid {
		return []byte(nt.Time.Format("2006-01-02T15:04:05Z")), nil
	}

	return []byte(`null`), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (nt *NullTime) UnmarshalJSON(data []byte) error {
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.Parse("2006-01-02T15:04:05Z", string(data))
	if err != nil {
		*nt = NullTime{Valid: false}
	} else {
		*nt = NullTime{Valid: true, Time: tt}
	}

	return nil
}

// Scan implements the Scanner interface.
func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ns = NullString{s.String, false}
	} else {
		*ns = NullString{s.String, true}
	}

	return nil
}

// Value implements the driver Valuer interface.
func (ns NullString) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}

	return ns.String, nil
}

// MarshalJSON implements json.Marshaler interface
func (ns NullString) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if ns.Valid == false {
		buf.WriteString(`null`)
		return []byte("null"), nil
	} else {
		buf.WriteString(`"` + ns.String + `"`)
	}

	return buf.Bytes(), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// Return null if string is considered as nullable
func (ns *NullString) UnmarshalJSON(in []byte) error {
	str := string(in)
	if str == `null` || str == `""` {
		return nil
	}

	ns.Valid, ns.String = true, str[1:len(str)-1]
	return nil
}
