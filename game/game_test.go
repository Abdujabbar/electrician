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
	if gm.IsFinished() {
		t.Error("Error while declaring game, on start user cannot be winner")
	}
}

func TestMoveAvailable(t *testing.T) {
	gm := NewGame(5)
	var err error
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if gm.isTurnedOn(i, j) {
				err = gm.Move(i, j)
				if err.Error() != "cell already turned on" {
					t.Error("Error on checking is turned on cell")
				}
			}
		}
	}
}
