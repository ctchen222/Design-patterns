package builder

import "fmt"

type IBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() House
}

type House struct {
	windowType string
	doorType   string
	floor      int
}

func (h *House) PrintHouse() {
	fmt.Println("Window Type: " + h.windowType + " Door Type: " + h.doorType + " Floor: " + string(h.floor))
}

func GetBuilder(builderType string) IBuilder {
	if builder, exists := builderStrategy[builderType]; exists {
		return builder()
	}
	// if builderType == "normal" {
	// 	return &normalBuilder{}
	// }
	// if builderType == "igloo" {
	// 	return &iglooBuilder{}
	// }

	return nil
}

type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (b *normalBuilder) SetWindowType() {
	b.windowType = "Wooden Windows"
}

func (b *normalBuilder) SetDoorType() {
	b.doorType = "Wooden Door"
}

func (b *normalBuilder) SetNumFloor() {
	b.floor = 2
}

func (b *normalBuilder) GetHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (b *iglooBuilder) SetWindowType() {
	b.windowType = "Snow Windows"
}

func (b *iglooBuilder) SetDoorType() {
	b.doorType = "Snow Door"
}

func (b *iglooBuilder) SetNumFloor() {
	b.floor = 1
}

func (b *iglooBuilder) GetHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
