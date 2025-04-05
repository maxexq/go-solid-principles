package main

import "fmt"

type (
	Player struct{}

	Item struct {
		Name      string
		Abilities []Ability
	}

	Heal        struct{}
	DamageBluff struct{}
	DefendBluff struct{}

	Ability interface {
		Execute()
	}
)

func (p *Player) Use(item *Item) {
	fmt.Println("Using item:", item.Name)

	for _, item := range item.Abilities {
		item.Execute()
	}
}

func (h *Heal) Execute() {
	fmt.Println("Healing")
}

func (d *DamageBluff) Execute() {
	fmt.Println("Damage Bluff increased by 10%")
}

// The Item struct has a method called Execute that increases the player's defense by 10%.
// This method is called when the player uses the item. The method is defined within the Item struct, which adheres to the Open/Closed Principle (OCP) because it allows us to add new abilities without modifying the Item struct itself.
func (d *DefendBluff) Execute() {
	fmt.Println("Defend Bluff increased by 10%")
}

func main() {
	// Creating an instance of Player and Item
	player := &Player{}
	item := &Item{
		Name: "Steak",
		// when the player eats the steak, it heals and increases damage bluff
		// The item can have multiple abilities, and we can add more abilities without modifying the Item struct.
		// This adheres to the Open/Closed Principle (OCP) because we can extend the functionality of the Item struct without modifying it.
		// For example, we can add a new ability called "DefendBluff" that increases the player's defense by 10% when the item is used.
		Abilities: []Ability{
			&Heal{},
			&DamageBluff{},
			// The DefendBluff ability is added here, which increases the player's defense by 10% when the item is used.
			// This adheres to the Open/Closed Principle (OCP) because we can add new abilities without modifying the Item struct itself.
			&DefendBluff{},
		},
	}

	// Player uses item
	player.Use(item)
}
