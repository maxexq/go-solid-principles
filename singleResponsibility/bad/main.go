package main

import "fmt"

type Player struct{}

func (p Player) Move() {
	fmt.Println("Moving")
}

func (p Player) Attack() {
	fmt.Println("Attacking")
}

func (p Player) DisplayScoreBoard() {
	fmt.Println("Display Score board")
}

func main() {
	// Creating an instance of Player
	player := &Player{}

	// Player actions
	player.Move()
	player.Attack()
	player.DisplayScoreBoard()
}

// The above code violates the Single Responsibility Principle (SRP) because the Player struct has multiple responsibilities: moving, attacking, and displaying the scoreboard.
// This makes the code less maintainable and harder to understand. Each responsibility should be encapsulated in its own struct or module.
// Displaying the scoreboard should be handled by a separate struct or function, allowing for better separation of concerns and easier testing.
