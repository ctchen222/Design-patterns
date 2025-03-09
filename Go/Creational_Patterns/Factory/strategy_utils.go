package factory

var gunFactoryStrategy = map[string]func() (IGun, error){
	"ak47":   func() (IGun, error) { return newAk47(), nil },
	"musket": func() (IGun, error) { return newMusket(), nil },
}

var animalFactoryStrategy = map[string]func() (IAnimal, error){
	"dog": func() (IAnimal, error) { return newDot(), nil },
	"cat": func() (IAnimal, error) { return newCat(), nil },
}
