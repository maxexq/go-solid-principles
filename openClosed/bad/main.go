package main

type (
	Player struct{}
	Steak  struct{}
)

func (p *Player) EatSteak(steak *Steak) {
	steak.Heal()
	steak.DamageBluff()

	// steak.DefendBluff() // This line violates the Open/Closed Principle (OCP) because it requires modifying the Steak struct to add new functionality or change existing behavior.
	// Instead, we should use an interface or a separate struct to handle the defense buff.
	steak.DefendBluff()
}

func (s *Steak) Heal() {
	println("Healing")
}

func (s *Steak) DamageBluff() {
	println("Damage increased by 10%")
}

// The Steak struct has a method called DefendBluff that increases the player's defense by 10%.
// This method is called when the player eats the steak. The method is defined within the Steak struct, which violates the Open/Closed Principle (OCP) because it requires modifying the Steak struct to add new functionality or change existing behavior.

func (s *Steak) DefendBluff() {
	println("Defend increased by 10%")
}

func main() {
	// Creating an instance of Player and Steak
	player := &Player{}
	steak := &Steak{}

	// Player eats steak
	player.EatSteak(steak)
}
