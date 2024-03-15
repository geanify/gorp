package rpg

type Inventory struct {
	Contents []*Item
}

func CreateNewInventory(size int) *Inventory {
	inventory := &Inventory{}
	inventory.Contents = make([]*Item, size)

	return inventory
}

func (inventory *Inventory) Add(item *Item) {
	inventory.Contents = append(inventory.Contents, item)
}
