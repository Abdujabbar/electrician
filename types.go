package main

import "github.com/Abdujabbor/electrician/game"

//GamesList type struct
type GamesList map[string]game.Game

//NewGamesList generates list of games on memory
func NewGamesList() *GamesList {
	return &GamesList{}
}
