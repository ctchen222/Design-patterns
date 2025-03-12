package observer

import (
	"fmt"
	"slices"
)

type Observer interface {
	Update(news string)
}

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	SetNews(new string)
	Notify()
}

type Agency struct {
	observers []Observer
	latesNews string
}

func (a *Agency) Attach(obs Observer) {
	a.observers = append(a.observers, obs)
}

func (a *Agency) Detach(obs Observer) {
	for i, observer := range a.observers {
		if observer == obs {
			a.observers = slices.Delete(a.observers, i, i+1)
			break
		}
	}
}

func (a *Agency) SetNews(new string) {
	a.latesNews = new
}

func (a *Agency) Notify() {
	for _, observer := range a.observers {
		observer.Update(a.latesNews)
	}
}

func NewAgency() *Agency {
	return &Agency{
		observers: []Observer{},
		latesNews: "",
	}
}

type Subscriber struct {
	name string
}

func (s *Subscriber) Update(news string) {
	fmt.Printf("%s received latest new: %s\n", s.name, news)
}

func NewSubscribe(name string) *Subscriber {
	return &Subscriber{
		name: name,
	}
}
