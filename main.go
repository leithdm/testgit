package main


/*
The key points that make this example helpful for understanding interfaces:
The Animal interface only cares about behavior (making a sound), not what the animal actually is
Each animal type (Dog, Cat, Cow) implements the interface by providing its own version of MakeSound()
The AnimalSound function doesn't care what type of animal it gets - it just needs something that can make a sound
You can easily add new animals without changing any existing code:
Apply to main.go
!
This demonstrates the main benefits of interfaces:
Code reuse (one function works with many types)
Flexibility (easy to add new types)
Abstraction (code only depends on the behavior it needs)
A more practical example might be saving data to different places:

*/

import "fmt"

// Animal interface defines a behavior (making a sound)
type Animal interface {
	MakeSound() string
}

// Concrete implementations
type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

type Cow struct {
	Name string
}

// Each type implements the Animal interface by defining MakeSound()
func (d Dog) MakeSound() string {
	return d.Name + " says: Woof!"
}

func (c Cat) MakeSound() string {
	return c.Name + " says: Meow!"
}

func (c Cow) MakeSound() string {
	return c.Name + " says: Moo!"
}

// This function can work with ANY animal
func AnimalSound(a Animal) {
	fmt.Println(a.MakeSound())
}

func main() {
	// Create some animals
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}
	cow := Cow{Name: "Bessie"}

	// We can call AnimalSound with any animal
	AnimalSound(dog) // Prints: Buddy says: Woof!
	AnimalSound(cat) // Prints: Whiskers says: Meow!
	AnimalSound(cow) // Prints: Bessie says: Moo!

	// We can also make a slice of animals
	farm := []Animal{dog, cat, cow}

	// Loop through all animals
	fmt.Println("\nAll farm animals:")
	for _, animal := range farm {
		AnimalSound(animal)
	}
}
