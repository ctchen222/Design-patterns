package factory

import "fmt"

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getPower() int {
	return g.power
}

type Ak47 struct {
	Gun
}

type Musket struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47",
			power: 4,
		},
	}
}

func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket",
			power: 1,
		},
	}
}

func GunFactory(gunType string) (IGun, error) {
	if gunFactory, exists := factoryStrategy[gunType]; exists {
		return gunFactory()
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

func PrintDetailed(g IGun) {
	fmt.Printf("Gun %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
