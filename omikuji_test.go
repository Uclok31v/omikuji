package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_Handler(t *testing.T) {
	cases := []struct {
		name     string
		datetime Time
		expected string
	}{
		{
			name:     "January 1st",
			datetime: Time{time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)},
			expected: "{\"omikuji\":\"大吉\"}",
		},
		{
			name:     "January 3rd",
			datetime: Time{time.Date(2020, 1, 3, 23, 59, 59, 0, time.Local)},
			expected: "{\"omikuji\":\"大吉\"}",
		},
		{
			name:     "January 4th",
			datetime: Time{time.Date(2020, 1, 4, 0, 0, 0, 0, time.Local)},
			expected: "{\"omikuji\":\"大凶\"}",
		},
	}

	for _, c := range cases {
		t.Helper()
		t.Run(c.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			c.datetime.handler(w, r)
			rw := w.Result()
			defer rw.Body.Close()

			if rw.StatusCode != http.StatusOK {
				t.Fatal("unexpected status code")
			}

			b, err := ioutil.ReadAll(rw.Body)
			if err != nil {
				t.Fatal("unexpected error")
			}

			actual := string(b)
			if c.expected != actual {
				t.Fatalf("unexpected response: %s", actual)
			}
		})
	}

}
