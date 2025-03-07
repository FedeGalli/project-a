package main

import (
	"github/FG412/tanks/player"
)

func main() {

	territories := player.InitializeTerritories()

	federico := player.Player{}
	federico.Owned_territories = append(federico.Owned_territories, territories["Europe"])
	nunzio := player.Player{}
	nunzio.Owned_territories = append(nunzio.Owned_territories, territories["America"])
	n_dice := 3
	federico.Attack(federico.Owned_territories[0], nunzio.Owned_territories[0], &n_dice)
	federico.Attack(federico.Owned_territories[0], nunzio.Owned_territories[0], &n_dice)

}
