package rpg

type Unit struct {
	Inventory Inventory
	Stats     *Stats
}

func (unit *Unit) GetStats() *Stats {
	stats := *unit.Stats

	for i := 0; i < len(unit.Inventory.Contents); i++ {
		stats.Add(&unit.Inventory.Contents[i].Stats)
	}

	return &stats
}

func (unit *Unit) Attack(defendingUnit *Unit) {

	attackerStats := unit.GetStats()
	defenderStats := defendingUnit.GetStats()

	damage := attackerStats.Attack - defenderStats.Armor

	defendingUnit.Stats.Health -= damage

}
