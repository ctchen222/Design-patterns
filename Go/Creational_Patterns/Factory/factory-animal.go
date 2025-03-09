package factory

import "fmt"

type IAnimal interface {
	Speak() string
}

type dog struct{}

func (d *dog) Speak() string {
	return "Woof!"
}

func newDot() IAnimal {
	return &dog{}
}

type cat struct{}

func (c *cat) Speak() string {
	return "Meow!"
}

func newCat() IAnimal {
	return &cat{}
}

func AnimalFactory(animal_type string) (IAnimal, error) {
	// switch animal_type {
	// case "dog":
	// 	return &dog{}, nil
	// case "cat":
	// 	return &cat{}, nil
	// default:
	// 	return nil, fmt.Errorf("Wrong animal type passed!")
	// }
	if animalFactory, exists := animalFactoryStrategy[animal_type]; exists {
		return animalFactory()
	}
	return nil, fmt.Errorf("No match animal")
}
