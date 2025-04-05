package main

type (
	Aircraft interface {
		Fly()
	}

	// Removed Fire() method from the Aircraft interface to adhere to LSP
	// This ensures that all implementations of the Aircraft interface can be used interchangeably without causing any issues.

	Boeing747 struct {
		Aircraft
	}

	SplitFire struct {
		Aircraft
	}
)

func (b Boeing747) Fly() {
	println("Boeing747 is flying")
}

func (s SplitFire) Fly() {
	println("SplitFire is flying")
}

func (s SplitFire) Fire() {
	println("SplitFire is firing")
}

func main() {
	boeing := &Boeing747{}
	splitFire := &SplitFire{}

	// Calling methods on Boeing747
	boeing.Fly()

	// Calling methods on SplitFire
	splitFire.Fly()
	splitFire.Fire()
}

// In this example, we have an interface `Aircraft` that defines the `Fly()` method. The `Boeing747` struct implements the `Fly()` method, and the `SplitFire` struct implements both the `Fly()` and `Fire()` methods. This adheres to the Liskov Substitution Principle because all implementations of the `Aircraft` interface can be used interchangeably without causing any issues.
// In this case, we have removed the `Fire()` method from the `Aircraft` interface, ensuring that all implementations of the `Aircraft` interface can be used interchangeably without causing any issues. The `SplitFire` struct can still implement the `Fire()` method without violating the Liskov Substitution Principle.
