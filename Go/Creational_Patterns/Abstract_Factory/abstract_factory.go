package AbstractFactory

import (
	"fmt"
)

type ISportFactory interface {
	MakeShoe() IShoe
	MakeShort() IShort
}

// Abstract Factory
func GetSportFactory(brand string) (ISportFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}

type Adidas struct{}

func (a *Adidas) MakeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) MakeShort() IShort {
	return &AdidasShort{
		Short: Short{
			logo: "adidas",
			size: 14,
		},
	}
}

type Nike struct{}

func (n *Nike) MakeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *Nike) MakeShort() IShort {
	return &NikeShort{
		Short: Short{
			logo: "nike",
			size: 14,
		},
	}
}

type IShoe interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) SetSize(size int) {
	s.size = size
}

func (s *Shoe) GetLogo() string {
	return s.logo
}

func (s *Shoe) GetSize() int {
	return s.size
}

type AdidasShoe struct {
	Shoe
}

type NikeShoe struct {
	Shoe
}

type IShort interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

type Short struct {
	logo string
	size int
}

func (s *Short) SetLogo(logo string) {
	s.logo = logo
}

func (s *Short) SetSize(size int) {
	s.size = size
}

func (s *Short) GetLogo() string {
	return s.logo
}

func (s *Short) GetSize() int {
	return s.size
}

type AdidasShort struct {
	Short
}

type NikeShort struct {
	Short
}
