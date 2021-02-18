package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	// mainで定義されているものを使う
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	player := "Pepper"
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {

		res := httptest.NewRecorder()
		server.ServeHTTP(res, newGetScoreRequest(player))

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "3", res.Body.String())
	})

	t.Run("get league", func(t *testing.T) {

		res := httptest.NewRecorder()
		server.ServeHTTP(res, newLeagueRequest())

		var got []Player
		err := json.NewDecoder(res.Body).Decode(&got)
		if err != nil {
			t.Fatalf("unable to parse response from server %q into slice of player, '%v", res.Body, err)
		}

		want := []Player{
			{"Pepper", 3},
		}
		assert.Equal(t, res.Code, http.StatusOK)
		assert.Equal(t, want, got)
	})
}
