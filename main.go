package main

import (
	game "github/FG412/tanks/game"
	player "github/FG412/tanks/player"
)

func main() {

	territories := player.InitializeTerritories()

	federico := player.Player{Name: "Federico"}
	federico.InitializePlayerTerritories([]*player.Territory{territories["Europe"]})
	nunzio := player.Player{Name: "Nunzio"}
	nunzio.InitializePlayerTerritories([]*player.Territory{territories["America"], territories["Antartica"]})

	game.MilitaryPhase(federico)

}
