package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const jsonContentType = "application/json"

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

// request単位でrouterを実行しないために、
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := &PlayerServer{store, http.NewServeMux()}

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router
	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	leagueTable := p.store.GetLeague()
	//w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", jsonContentType)
	if err := json.NewEncoder(w).Encode(leagueTable); err != nil {
		log.Fatalf("cloud not encode %v", leagueTable)
	}
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

/*
// scoreの取得先をここで規定しない
// server内部に格納先storeと取得方法getを定義しておいて、
// serverを定義した時に方法を変えられるようにしておく
func (p *PlayerServer) GetPlayerScore(name string) string {
	switch name {
	case "Pepper":
		return "20"
	case "Floyd":
		return "10"
	}

	return ""
}
*/
