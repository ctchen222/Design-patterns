package factory

var factoryStrategy = map[string]func() (IGun, error){
	"ak47":   func() (IGun, error) { return newAk47(), nil },
	"musket": func() (IGun, error) { return newMusket(), nil },
}
