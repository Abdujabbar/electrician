package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/Abdujabbor/electrician/randomstring"

	"github.com/Abdujabbor/electrician/game"

	"github.com/gorilla/mux"
)

var gList = NewGamesList()

func main() {
	mxRouter := mux.NewRouter()
	mxRouter.HandleFunc("/", serveTemplate)
	mxRouter.HandleFunc("/move", move).Methods("POST")
	mxRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.ListenAndServe(":8000", mxRouter)
}

func move(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var currentGame *game.Game
	var row, col int
	var err error
	var ok bool
	moveResponse := MoveResponse{
		Success: true,
	}

	lst := *gList
	if currentGame, ok = lst[r.FormValue("hash")]; !ok {
		moveResponse.Success = false
		moveResponse.Error = "Game not exists please refresh page"
	}
	row, err = strconv.Atoi(r.FormValue("row"))
	if err != nil {
		moveResponse.Success = false
		moveResponse.Error = err.Error()
	}
	col, err = strconv.Atoi(r.FormValue("col"))
	if err != nil {
		moveResponse.Success = false
		moveResponse.Error = err.Error()
	}
	err = currentGame.Move(row, col)
	if err != nil {
		moveResponse.Error = err.Error()
		moveResponse.Success = false
	}
	moveResponse.Board = currentGame.GetBoard()
	moveResponse.MoveCounter = currentGame.GetMoveCounter()
	moveResponse.Finished = currentGame.IsFinished()
	response, _ := json.Marshal(moveResponse)
	w.Write(response)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	gm := game.NewGame(5)
	randomGenerator := randomstring.NewGenerator("abcdefg", 10)
	hash := randomGenerator.Generate()
	gList.Add(hash, gm)
	lp := filepath.Join("templates/layouts", "base.html")
	fp := filepath.Join("templates", "index.html")
	files := []string{fp, lp}
	tmpl, _ := template.ParseFiles(files...)
	data := struct {
		Game *game.Game
		Hash string
	}{
		gm,
		hash,
	}
	tmpl.ExecuteTemplate(w, "layout", data)
}
