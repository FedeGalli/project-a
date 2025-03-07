package ui

import (
	"fmt"
	player "github/FG412/tanks/player"
)

type Session struct {
	players []player.Player
}

func (s *Session) InitSession() {
	s.players = []player.Player{}
	//add players in the session
}

func (s *Session) StartGame() {
	someone_wins := false
	for !someone_wins {
		for _, user := range s.players {
			//EconomicPhase(user)
			MilitaryPhase(user)
			//MovePhase()
		}
	}
}

func MilitaryPhase(p player.Player) {
	fmt.Println("Select your next action: ")
	fmt.Println("1. Attack ")
	fmt.Println("2. End military phase ")

	var user_input int

	switch fmt.Scanf("%d\n", &user_input); user_input {
	case 1:
		for user_input != 0 {

			fmt.Println("Select the territory to attack: ")
			i := 1
			attackable_territories := p.GetAttackableTerritories()
			for _, territory := range attackable_territories {
				fmt.Printf("%d. Region: %v \tOwner:%v \tTanks:%v \t\n", i, territory.Name, territory.Owner.Name, territory.N_tanks)
				i++
			}

			fmt.Print("0. Stop attacking \n")

			switch fmt.Scanf("%d\n", &user_input); user_input {
			case 0:
				return
			//getting the right territory to attack
			default:
				var target_territory *player.Territory
				var from_territory *player.Territory
				target_territory = attackable_territories[user_input-1]
				fmt.Printf("You are attacking %v! \nNow select from which of your territories you want to attack: \n",
					target_territory.Name)

				owned_side_territories := p.GetOwnedSideTerritories(target_territory)
				i = 1
				for _, territory := range owned_side_territories {
					fmt.Printf("%d. Region: %v Tanks:%v \n", i, territory.Name, territory.N_tanks)
					i++
				}

				fmt.Scanf("%d\n", &user_input)
				from_territory = owned_side_territories[user_input-1]

				fmt.Println("Insert the number of dice you want to roll: ")
				fmt.Scanf("%d\n", &user_input)

				p.Attack(from_territory, target_territory, &user_input)

			}
		}

	case 2:
		return
	}

}
