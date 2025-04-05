package main

type (
	Aircraft interface {
		Fly()
	}

	FighterAircraft interface {
		Aircraft
		Fire()
	}

	CommercialAircraft interface {
		Aircraft
		Aircondition()
	}

	Boeing747 struct{}
	SplitFire struct{}
)

func (b Boeing747) Fly() {
	println("Boeing747 is flying")
}

func (b Boeing747) Aircondition() {
	println("Boeing747 is air conditioning")
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
	boeing.Aircondition()

	// Calling methods on SplitFire
	splitFire.Fly()
	splitFire.Fire()
}

// In this example, we have separated the `Aircraft` interface into two interfaces: `FighterAircraft` and `CommercialAircraft`. The `FighterAircraft` interface includes the `Fire()` method, while the `CommercialAircraft` interface includes the `Aircondition()` method. This adheres to the Interface Segregation Principle (ISP) because no client is forced to depend on methods it does not use. The `Boeing747` struct implements the `CommercialAircraft` interface, while the `SplitFire` struct implements the `FighterAircraft` interface.
// This way, we ensure that all implementations of the `Aircraft` interface can be used interchangeably without causing any issues. The `Boeing747` struct can implement the `Aircondition()` method without violating the Interface Segregation Principle, and the `SplitFire` struct can implement the `Fire()` method without violating the Interface Segregation Principle.
