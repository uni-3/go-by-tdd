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

	cases := []struct {
		name     string
		player   string
		expected string
	}{
		{
			name:     "return score 20",
			player:   "Pepper",
			expected: "20",
		},
		{
			name:     "return score 10",
			player:   "Floyd",
			expected: "10",
		},
	}
	for _, test := range cases {

		t.Run(test.name, func(t *testing.T) {
			req := newGetScoreRequest(test.player)
			res := httptest.NewRecorder()
			server.ServeHTTP(res, req)

			actual := res.Body.String()
			assert.Equal(t, test.expected, actual)
		})
	}
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
