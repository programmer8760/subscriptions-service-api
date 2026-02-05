package domain

import (
	"encoding/json"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

func NewDate(t time.Time) Date {
	return Date{Time: t}
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		time.Time(d.Time).Format("01-2006"),
	)
}

func (d *Date) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("01-2006", value)
	if err != nil {
		return err
	}

	d.Time = t

	return nil
}
