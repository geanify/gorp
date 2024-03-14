package rpg

type Stats struct {
	Health int
	Armor  int
	Attack int
}

func (stats *Stats) Add(_stats *Stats) {
	stats.Health += _stats.Health
	stats.Armor += _stats.Armor
	stats.Attack += _stats.Attack
}
