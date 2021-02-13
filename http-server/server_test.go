package main

import (
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

				assert.Equal(t, test.status, res.Code, "check for response status")
			})
		}
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, name, nil)
	return req
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}
