package main

import "github.com/Abdujabbor/electrician/game"

//MoveResponse struct
type MoveResponse struct {
	Success bool     `json:"success"`
	Error   string   `json:"error"`
	Board   [][]bool `json:"board"`
}

//GamesList type struct
type GamesList map[string]*game.Game

//NewGamesList generates list of games on memory
func NewGamesList() *GamesList {
	return &GamesList{}
}

//Add adds new item to games list
func (gl *GamesList) Add(hash string, g *game.Game) {
	lst := *gl
	lst[hash] = g
	gl = &lst
}
