package omikuji

import (
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name    string
		time    time.Time
		wantMsg string
	}{
		{name: "Get Random Message", time: time.Now(), wantMsg: ""},
		{name: "Get Daikichi Message", time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), wantMsg: "大吉"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			o := New(tt.time)
			r := o.Run()

			if tt.wantMsg != "" && r.Msg != tt.wantMsg {
				t.Error("Got Wrong Message")
			}
		})
	}
}
