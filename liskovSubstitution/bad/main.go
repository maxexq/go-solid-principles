package main

import "fmt"

type (
	Aircraft interface {
		Fly()

		// Fire() should not be here if not all aircraft can fire
		// This violates the Liskov Substitution Principle (LSP)
		// because a Boeing747 cannot fire, but it implements the Aircraft interface.
		Fire()
	}

	Boeing747 struct{}
	SplitFire struct{}
)

func (b Boeing747) Fly() {
	fmt.Println("Boeing747 is flying")
}

func (b Boeing747) Fire() {
	panic("Error: Boeing747 cannot fire")
}

func (s SplitFire) Fly() {
	fmt.Println("SplitFire is flying")
}

func (s SplitFire) Fire() {
	fmt.Println("SplitFire is firing")
}

func main() {
	boeing := Boeing747{}
	splitFire := SplitFire{}

	// Calling methods on Boeing747
	boeing.Fly()
	// boeing.Fire() // This will panic

	// Calling methods on SplitFire
	splitFire.Fly()
	splitFire.Fire()
}

// In this example, we have an interface `Aircraft` that defines two methods: `Fly()` and `Fire()`. The `Boeing747` struct implements the `Fly()` method but panics when calling the `Fire()` method. This violates the Liskov Substitution Principle because a `Boeing747` cannot be used as a substitute for an `Aircraft` without causing a panic when calling the `Fire()` method.
// In the above code, we have an interface `Aircraft` with methods `Fly()` and `Fire()`. The `Boeing747` struct implements the `Fly()` method but panics when calling the `Fire()` method. This violates the Liskov Substitution Principle because a `Boeing747` cannot be used as a substitute for an `Aircraft` without causing a panic when calling the `Fire()` method.
// To adhere to the Liskov Substitution Principle, we should ensure that all methods in the interface are implemented correctly by all structs that implement the interface. In this case, we could either remove the `Fire()` method from the `Aircraft` interface or create a separate interface for aircraft that can fire. This way, we can ensure that all implementations of the `Aircraft` interface can be used interchangeably without causing any issues.
