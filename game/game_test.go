package game

import (
	"fmt"
	"testing"
)

func TestNewGamePanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	NewGame(2)
}

func TestMoveOnGame(t *testing.T) {
	gm := NewGame(5)
	err := gm.Move(-1, 0)
	if err == nil {
		t.Error("Error while moving by wrong index")
	}
}

func TestWinOnGame(t *testing.T) {
	gm := NewGame(5)
	if gm.isWin() {
		t.Error("Error while declaring game, on start user cannot be winner")
	}
}
