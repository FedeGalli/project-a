package player

import (
	"fmt"
	"math/rand/v2"
)

type Player struct {
	Name              string
	Owned_territories map[string]*Territory
}

func bubbleSort(list *[]int) {
	n := len(*list)
	for i := range n - 1 {
		for j := i + 1; j < n; j++ {
			if (*list)[i] < (*list)[j] {
				tmp := (*list)[i]
				(*list)[i] = (*list)[j]
				(*list)[j] = tmp
			}
		}
	}
}

func (me *Player) InitializePlayerTerritories(territories []*Territory) {
	//initializing the map first
	me.Owned_territories = map[string]*Territory{}
	for _, territory := range territories {
		me.Owned_territories[territory.Name] = territory
		territory.Owner = me
	}
}

func (me *Player) GetAttackableTerritories() map[*Territory][]*Territory {

	attackable_territories := map[*Territory][]*Territory{}
	for _, owned_territory := range me.Owned_territories {
		if owned_territory.N_tanks > 1 { //you need at least 2 troups to attack a new territory
			for _, adj_territory := range owned_territory.Adj_territories {

				if adj_territory.Owner != me {
					fmt.Printf("from: %v adj territory: %v Owner: %v %v adj owner: %v %v: \n", owned_territory.Name, adj_territory.Name, me.Name, &me, adj_territory.Owner.Name, &adj_territory.Owner)
					attackable_territories[adj_territory] = append(attackable_territories[adj_territory], owned_territory)
				}
			}
		}
	}
	return attackable_territories
}

func (me *Player) Attack(from, to *Territory, n_dice *int) {

	mine_dices := []int{}
	opponent_dices := []int{}

	//min between the attacker choosen number of dices and the tanks available
	for range min(*n_dice, from.N_tanks) {
		mine_dices = append(mine_dices, rand.IntN(6)+1) // Returns a random integer from 1 to 6
	}

	//min between the attacker choosen number of dices and the tanks available for the defender
	for range min(*n_dice, to.N_tanks) {
		opponent_dices = append(opponent_dices, rand.IntN(6)+1) // Returns a random integer from 1 to 6
	}

	bubbleSort(&mine_dices)
	bubbleSort(&opponent_dices)

	fmt.Printf("Me: %v Opponent: %v \n", mine_dices, opponent_dices)

	//min between the attacker choosen number of dices and the tanks available for the defender
	for i := range min(*n_dice, to.N_tanks) {
		if mine_dices[i] <= opponent_dices[i] {
			from.N_tanks--
		} else {
			to.N_tanks--
		}
	}

	fmt.Printf("You have %v tanks, the opponent: %v\n", from.N_tanks, to.N_tanks)

	//if there aren't more troups on enemy territory, move troups
	moving_units := 0
	if to.N_tanks == 0 {
		//Conquered territory attack logic (you can move up to 3 units into the new territory)
		switch {
		case *n_dice < from.N_tanks:
			moving_units = *n_dice
		default:
			moving_units = *n_dice - 1
		}
		me.ConquerTerritory(&moving_units, from, to)
	}

}

func (me *Player) ConquerTerritory(n_troups *int, from, to *Territory) {
	delete(to.Owner.Owned_territories, to.Name) //deleting territory prev-ownership
	me.Owned_territories[to.Name] = to
	to.Owner = me

	//move troups to the new conquered territory
	me.MoveTroups(n_troups, from, to)
}

func (me *Player) MoveTroups(n_troups *int, from, to *Territory) {
	from.N_tanks -= *n_troups
	to.N_tanks += *n_troups
}
