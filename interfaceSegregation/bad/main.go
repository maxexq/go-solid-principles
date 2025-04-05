package main

type (
	Aircraft interface {
		Fly()
		Fire()
		Aircondition()
	}

	Boeing747 struct{}
	SplitFire struct{}
)

func (b Boeing747) Fly() {
	println("Boeing747 is flying")
}

func (b Boeing747) Fire() {
	panic("Error: Boeing747 cannot fire")
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

func (s SplitFire) Aircondition() {
	panic("Error: SplitFire cannot air condition")
}

func main() {
	boeing := &Boeing747{}
	splitFire := &SplitFire{}

	// Calling methods on Boeing747
	boeing.Fly()
	boeing.Fire()
	boeing.Aircondition()

	// Calling methods on SplitFire
	splitFire.Fly()
	splitFire.Fire()
	splitFire.Aircondition()
}

// interface segregation principle (ISP) states that no client should be forced to depend on methods it does not use. In this example, we have an interface Aircraft that defines three methods: Fly(), Fire(), and Aircondition(). The Boeing747 struct implements the Fly() and Aircondition() methods but panics when calling the Fire() method. This violates the ISP because a Boeing747 cannot be used as a substitute for an Aircraft without causing a panic when calling the Fire() method.
// To adhere to the ISP, we should ensure that all methods in the interface are implemented correctly by all structs that implement the interface. In this case, we could either remove the Fire() method from the Aircraft interface or create a separate interface for aircraft that can fire. This way, we can ensure that all implementations of the Aircraft interface can be used interchangeably without causing any issues.
