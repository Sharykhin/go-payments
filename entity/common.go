package entity

import (
	"bytes"
	"database/sql"
	"time"
)

type (
	NullString sql.NullString
	NullTime   time.Time
)

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
