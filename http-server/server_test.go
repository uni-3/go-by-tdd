package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
	}
	server := &PlayerServer{&store}

	t.Run("200", func(t *testing.T) {
		cases := []struct {
			name     string
			player   string
			expected string
			status   int
		}{
			{
				name:     "return score 20",
				player:   "Pepper",
				expected: "20",
				status:   200,
			},
			{
				name:     "return score 10",
				player:   "Floyd",
				expected: "10",
				status:   200,
			},
		}
		for _, test := range cases {

			t.Run(test.name, func(t *testing.T) {
				req := newGetScoreRequest(test.player)
				res := httptest.NewRecorder()
				server.ServeHTTP(res, req)

				assert.Equal(t, res.Code, test.status)
				assert.Equal(t, test.expected, res.Body.String())
			})
		}

	})

	t.Run("404", func(t *testing.T) {
		cases := []struct {
			name     string
			player   string
			expected string
			status   int
		}{
			{
				name:     "missing player",
				player:   "Appolo",
				expected: "",
				status:   http.StatusNotFound,
			},
		}
		for _, test := range cases {

			t.Run(test.name, func(t *testing.T) {
				req := newGetScoreRequest(test.player)
				res := httptest.NewRecorder()
				server.ServeHTTP(res, req)

				assert.Equal(t, res.Code, test.status, "check for response status")
			})
		}
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
	}
	server := &PlayerServer{&store}

	cases := []struct {
		name     string
		expected string
	}{
		{
			name:     "accepted recodes win on POST",
			expected: "Pepper",
		},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			player := "Pepper"
			req := newPostWinRequest(player)
			res := httptest.NewRecorder()

			server.ServeHTTP(res, req)

			assert.Equal(t, http.StatusAccepted, res.Code, "check for status")
			assert.Len(t, store.winCalls, 1, "check for length")
			assert.Equal(t, test.expected, store.winCalls[0], "check for post name")
		})
	}
}

// Test .
func TestLeague(t *testing.T) {
	store := StubPlayerStore{}
	server := &PlayerServer{&store}
	t.Run("return 200 on league", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/league", nil)
		res := httptest.NewRecorder()
		server.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, name, nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}
