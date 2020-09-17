package server

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gopherdojo/gopherdojo-studyroom/kadai4/komazz/omikuji"
)

func TestOmikujiHandler(t *testing.T) {
	tests := []struct {
		name    string
		time    time.Time
		wantMsg string
	}{
		{name: "Omikuji Get Random Message", time: time.Now(), wantMsg: ""},
		{name: "Omikuji Get Daikichi Message", time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), wantMsg: "大吉"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &Server{omikuji.New(tt.time)}
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/omikuji/", nil)
			s.omikujiHandler(w, r)

			rw := w.Result()
			defer rw.Body.Close()

			var rs omikuji.Result
			err := json.NewDecoder(rw.Body).Decode(&rs)
			if err != nil {
				t.Error("Fail to Decode")
			}

			if tt.wantMsg != "" && rs.Msg != tt.wantMsg {
				t.Error("Got Wrong Message")
			}
		})
	}
}
