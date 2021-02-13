package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

/*
// scoreの取得先をここで規定しない
// server内部に格納先storeと取得方法getを定義しておいて、定義した時に方法を変えられるようにしておく
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
