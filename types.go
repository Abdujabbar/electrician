package main

import "github.com/Abdujabbor/electrician/game"
import "github.com/Abdujabbor/electrician/randomstring"

//GamesList type struct
type GamesList map[string]game.Game

//NewGamesList generates list of games on memory
func NewGamesList() *GamesList {
	return &GamesList{}
}

//Add adds new item to games list
func (gl *GamesList) Add(g game.Game) {
	lst := *gl
	randomStringGenerator := randomstring.NewGenerator("abcdefghijklmnopqrstuvwxyz", 10)
	for {
		hash := randomStringGenerator.Generate()
		if _, ok := lst[hash]; !ok {
			lst[hash] = g
			break
		}
	}
	gl = &lst
}
