package player

import (
	"fmt"
	"math/rand/v2"
)

type Player struct {
	name              string
	Owned_territories []*Territory
}

func bubbleSort(list *[]int) {
	n := len(*list)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if (*list)[i] < (*list)[j] {
				tmp := (*list)[i]
				(*list)[i] = (*list)[j]
				(*list)[j] = tmp
			}
		}
	}
}

func (me *Player) Attack(from, to *Territory, n_dice *int) {

	mine_dices := []int{}
	opponent_dices := []int{}

	for range *n_dice {
		mine_dices = append(mine_dices, rand.IntN(6)+1) // Returns a random integer from 1 to 6
	}

	//min between the attacker choosen number of dices and the tanks available for the defender
	for range min(*n_dice, to.N_tanks) {
		opponent_dices = append(opponent_dices, rand.IntN(6)+1) // Returns a random integer from 1 to 6
	}

	bubbleSort(&mine_dices)
	bubbleSort(&opponent_dices)

	fmt.Printf("Me: %v Opponent: %v", mine_dices, opponent_dices)

	//min between the attacker choosen number of dices and the tanks available for the defender
	for i := range min(*n_dice, to.N_tanks) {
		if mine_dices[i] <= opponent_dices[i] {
			from.N_tanks--
		} else {
			to.N_tanks--
		}
	}

	fmt.Printf("You have %v tanks, the opponent: %v\n", from.N_tanks, to.N_tanks)
}
