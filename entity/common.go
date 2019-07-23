package entity

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"reflect"
	"time"
)

type (
	NullString sql.NullString
	NullTime   struct {
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

func (ns NullString) Value() (driver.Value, error) {
	if ns.Valid {
		return ns.String, nil
	}

	return nil, nil
}

func (s NullString) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if s.Valid == false {
		buf.WriteString(`null`)
		return []byte("null"), nil
	} else {
		buf.WriteString(`"` + s.String + `"`)
	}

	return buf.Bytes(), nil
}

func (s *NullString) UnmarshalJSON(in []byte) error {
	str := string(in)
	if str == `null` {
		return nil
	}

	s.Valid, s.String = true, str
	return nil
}
