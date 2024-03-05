package rpg

type Unit struct {
	Inventory []*Item
	Stats     *Stats
}

func (unit *Unit) GetStats() *Stats {
	stats := *unit.Stats

	for i := 0; i < len(unit.Inventory); i++ {
		stats.Add(&unit.Inventory[i].Stats)
	}

	return &stats
}

func (unit *Unit) Attack(defendingUnit *Unit) {

	// attackerStats := unit.GetStats()
	// defenderStats := defendingUnit.GetStats()

}
