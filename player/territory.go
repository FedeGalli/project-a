package player

type Territory struct {
	Name            string
	N_tanks         int
	Adj_territories []*Territory
	Owner           *Player
}

func InitializeTerritories() map[string]*Territory {
	territories := map[string]*Territory{}

	territories["Europe"] = &Territory{Name: "Europe", N_tanks: 5, Adj_territories: nil}
	territories["America"] = &Territory{Name: "America", N_tanks: 2, Adj_territories: nil}
	territories["Antartica"] = &Territory{Name: "Antartica", N_tanks: 2, Adj_territories: nil}

	territories["America"].Adj_territories = append(territories["America"].Adj_territories, territories["Europe"])
	territories["Europe"].Adj_territories = append(territories["Europe"].Adj_territories, territories["America"])

	territories["America"].Adj_territories = append(territories["America"].Adj_territories, territories["Antartica"])
	territories["Antartica"].Adj_territories = append(territories["Antartica"].Adj_territories, territories["America"])

	return territories
}
