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

func (me *Player) GetAttackableTerritories() map[string]*Territory {

	attackable_territories := map[string]*Territory{}
	for _, owned_territory := range me.Owned_territories {
		if owned_territory.N_tanks > 1 {
			for _, adj_territory := range owned_territory.Adj_territories {
				if attackable_territories[adj_territory.Name] == nil && adj_territory.Owner != me {
					attackable_territories[adj_territory.Name] = adj_territory
				}
			}
		}
	}
	return attackable_territories
}

func (me *Player) GetOwnedSideTerritories(target *Territory) map[int]*Territory {

	owned_side_territories := map[int]*Territory{}
	counter := 0
	for _, owned_territory := range me.Owned_territories {
		if owned_territory.N_tanks > 1 {
			for _, adj_territory := range owned_territory.Adj_territories {
				if adj_territory == target {
					owned_side_territories[counter] = owned_territory
					counter++
				}
			}
		}
	}
	return owned_side_territories
}

func (me *Player) ConquerTerritory(target *Territory) {
	me.Owned_territories[target.Name] = target
	target.Owner = me
}

func (me *Player) Attack(from, to *Territory, n_dice *int) bool {

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

	return to.N_tanks == 0
}
