package player

type Territory struct {
	N_tanks         int
	Adj_territories []*Territory
}

func InitializeTerritories() map[string]*Territory {
	territories := map[string]*Territory{}

	territories["Europe"] = &Territory{N_tanks: 5, Adj_territories: nil}
	territories["America"] = &Territory{N_tanks: 2, Adj_territories: nil}

	territories["America"].Adj_territories = append(territories["America"].Adj_territories, territories["Europe"])
	territories["Europe"].Adj_territories = append(territories["Europe"].Adj_territories, territories["America"])

	return territories
}
