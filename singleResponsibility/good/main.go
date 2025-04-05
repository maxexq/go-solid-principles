package main

import "fmt"

type (
	Player  struct{}
	Display struct{}
)

func (p Player) Move() {
	fmt.Println("Moving")
}

func (p Player) Attack() {
	fmt.Println("Attacking")
}

func (d Display) DisplayScoreBoard() {
	fmt.Println("Display Score board")
}

func main() {
	// Creating instances of Player and Display
	player := &Player{}
	display := &Display{}

	// Player actions
	player.Move()
	player.Attack()

	// Displaying the scoreboard is now handled by the Display struct
	display.DisplayScoreBoard()
}

// The above code adheres to the Single Responsibility Principle (SRP) because the Player struct is responsible for player actions (moving and attacking), while the Display struct is responsible for displaying the scoreboard. This separation of concerns makes the code more maintainable and easier to understand. Each responsibility is encapsulated in its own struct, allowing for better organization and testing.
// The Display struct can be further extended to handle other display-related functionalities without affecting the Player struct. This modular approach enhances code readability and maintainability, making it easier to manage changes in the future.
