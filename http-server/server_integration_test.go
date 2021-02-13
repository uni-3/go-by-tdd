package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	// mainで定義されているものを使う
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}

	cases := []struct {
		name string
	}{
		{
			name: "record 3 times, then get how many store",
		},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			player := "Peppar"

			server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
			server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
			server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
			res := httptest.NewRecorder()
			server.ServeHTTP(res, newGetScoreRequest(player))

			assert.Equal(t, res.Code, http.StatusOK)
			assert.Equal(t, "3", res.Body.String())
		})
	}
}
