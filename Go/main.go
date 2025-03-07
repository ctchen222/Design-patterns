package main

import (
	"fmt"
	"sync"

	"github.com/ctchen1999/design-patterns/Creational_Patterns/Abstract_Factory"
	builder "github.com/ctchen1999/design-patterns/Creational_Patterns/Builder"
	"github.com/ctchen1999/design-patterns/Creational_Patterns/Singleton"
	"github.com/ctchen1999/design-patterns/Creational_Patterns/factory"
)

func main() {
	fmt.Println("Design Patterns in Go")

	// Singleton
	fmt.Println("[Singleton Start]")
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		instance1 := Singleton.GetInstance()
		fmt.Printf("instance1: %p\n", instance1)
	}()

	go func() {
		defer wg.Done()
		instance2 := Singleton.GetInstance()
		fmt.Printf("instance2: %p\n", instance2)
	}()

	wg.Wait()
	fmt.Println("[Singleton Done]")

	fmt.Println("[Factory Start]")
	ak47, _ := factory.GunFactory("ak47")
	musket, _ := factory.GunFactory("musket")
	factory.PrintDetailed(ak47)
	factory.PrintDetailed(musket)
	fmt.Println("[Factory Done]")

	fmt.Println("[Abstract Factory Start]")
	adidasFactory, _ := AbstractFactory.GetSportFactory("adidas")
	nikeFactory, _ := AbstractFactory.GetSportFactory("nike")
	adidasShoe := adidasFactory.MakeShoe()
	nikeShort := nikeFactory.MakeShort()
	fmt.Printf("Adidas Shoe Logo: %s\n", adidasShoe.GetLogo())
	fmt.Printf("Nike Shoe Logo: %s\n", nikeShort.GetLogo())

	fmt.Println("[Abstract Factory Done]")

	fmt.Println("[Builder Start]")
	normalBuilder := builder.GetBuilder("normal")
	iglooBuilder := builder.GetBuilder("igloo")
	normalBuilder.SetWindowType()
	normalBuilder.SetDoorType()
	normalBuilder.SetNumFloor()
	iglooBuilder.SetWindowType()
	iglooBuilder.SetDoorType()
	iglooBuilder.SetNumFloor()
	normalHouse := normalBuilder.GetHouse()
	iglooHouse := iglooBuilder.GetHouse()
	normalHouse.PrintHouse()
	iglooHouse.PrintHouse()
	fmt.Println("[Builder Done]")
}
