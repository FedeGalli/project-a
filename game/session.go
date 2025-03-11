package ui

import (
	"fmt"
	player "github/FG412/project-a/player"
)

type Session struct {
	Id          int32
	players     []*player.Player
	Territories map[string]*player.Territory
}

func (s *Session) InitSession() {
	// method testing implementation
	s.Territories = player.InitializeTerritories()
	s.players = []*player.Player{&player.Player{Name: "Federico"}, &player.Player{Name: "Nunzio"}}

	s.players[0].InitializePlayerTerritories([]*player.Territory{s.Territories["Europe"]})
	s.players[1].InitializePlayerTerritories([]*player.Territory{s.Territories["America"], s.Territories["Antartica"]})
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

func MilitaryPhase(p *player.Player) {
	fmt.Printf("%v Select your next action: \n", p.Name)
	fmt.Println("1. Attack ")
	fmt.Println("2. End military phase ")

	var user_input int

	switch fmt.Scanf("%d\n", &user_input); user_input {
	case 1:
		for user_input != 0 {

			fmt.Println("Select the territory to attack: ")
			i := 1
			attackable_territories := p.GetAttackableTerritories()
			user_choice_map := map[int]*player.Territory{}
			for territory := range attackable_territories {
				fmt.Printf("%d. Region: %v \tOwner:%v \tTanks:%v \t\n", i, territory.Name, territory.Owner.Name, territory.N_tanks)
				user_choice_map[i] = territory
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
				target_territory = user_choice_map[user_input]
				fmt.Printf("You are attacking %v! \nNow select from which of your territories you want to attack: \n",
					target_territory.Name)

				i = 1
				for _, from := range attackable_territories[target_territory] {
					fmt.Printf("%d. Region: %v Tanks:%v \n", i, from.Name, from.N_tanks)
					user_choice_map[i] = from
					i++
				}

				fmt.Scanf("%d\n", &user_input)
				from_territory = user_choice_map[user_input]

				fmt.Println("Insert the number of dice you want to roll: ")
				fmt.Scanf("%d\n", &user_input)

				p.Attack(from_territory, target_territory, &user_input)

			}
		}

	case 2:
		return
	}

}
