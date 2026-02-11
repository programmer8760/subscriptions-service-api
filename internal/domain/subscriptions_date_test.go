package domain

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDateMarshalJSON(t *testing.T) {
	d := NewDate(time.Date(2006, time.January, 1, 0, 0, 0, 0, time.UTC))

	b, err := d.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON() error = %v", err)
	}

	if got, want := string(b), `"01-2006"`; got != want {
		t.Fatalf("MarshalJSON() = %s, want %s", got, want)
	}
}

func TestDateUnmarshalJSON(t *testing.T) {
	tests := []struct {
		input     string
		wantMonth time.Month
		wantYear  int
		wantErr   bool
	}{
		{`"04-2011"`, time.April, 2011, false},
		{`"2006-01"`, time.January, 0, true},
		{`"2006-01-02T15:04:05Z07:00"`, time.January, 0, true},
		{`"abcd"`, time.January, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			var d Date
			err := json.Unmarshal([]byte(tt.input), &d)
			if err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("UnmarshalJSON() error = %v", err)
			}
			if d.Month() != tt.wantMonth {
				t.Errorf("month = %v, want %v", d.Month(), tt.wantMonth)
			}
			if d.Year() != tt.wantYear {
				t.Errorf("year = %d, want %d", d.Year(), tt.wantYear)
			}
		})
	}
}
