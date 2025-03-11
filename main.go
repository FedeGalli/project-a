package main

import (
	game "github/FG412/project-a/game"
)

func main() {

	s := game.Session{}
	s.InitSession()
	s.StartGame()

}
