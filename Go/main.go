package main

import (
	"fmt"
	"sync"

	observer "github.com/ctchen1999/design-patterns/Behavioral_Patterns/Observer"
	strategy "github.com/ctchen1999/design-patterns/Behavioral_Patterns/Strategy"
	abstractFactory "github.com/ctchen1999/design-patterns/Creational_Patterns/Abstract_Factory"
	builder "github.com/ctchen1999/design-patterns/Creational_Patterns/Builder"
	factory "github.com/ctchen1999/design-patterns/Creational_Patterns/Factory"
	singleton "github.com/ctchen1999/design-patterns/Creational_Patterns/Singleton"
	adaptor "github.com/ctchen1999/design-patterns/Structural_Patterns/Adaptor"
	decorator "github.com/ctchen1999/design-patterns/Structural_Patterns/Decorator"
	facade "github.com/ctchen1999/design-patterns/Structural_Patterns/Facade"
	proxy "github.com/ctchen1999/design-patterns/Structural_Patterns/Proxy"
)

func main() {
	fmt.Println("Design Patterns in Go")

	// Singleton
	fmt.Println("[Singleton Start]")
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		instance1 := singleton.GetInstance()
		fmt.Printf("instance1: %p\n", instance1)
	}()

	go func() {
		defer wg.Done()
		instance2 := singleton.GetInstance()
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

	fmt.Println("[Factory Start]")
	dog, _ := factory.AnimalFactory("dog")
	dogSpeak := dog.Speak()
	cat, _ := factory.AnimalFactory("cat")
	catSpeak := cat.Speak()
	fmt.Println(dogSpeak)
	fmt.Println(catSpeak)
	fmt.Println("[Factory Done]")

	fmt.Println("[Abstract Factory Start]")
	adidasFactory, _ := abstractFactory.GetSportFactory("adidas")
	nikeFactory, _ := abstractFactory.GetSportFactory("nike")
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

	fmt.Println("[Adaptor Start]")
	paypal := &adaptor.PayPal{}
	paypal.Pay(1000)
	stripe := &adaptor.Stripe{}
	stripeAdaptor := &adaptor.StripeAdaptor{
		StripeService: stripe,
	}
	stripeAdaptor.Pay(2000)
	fmt.Println("[Adaptor Done]")

	fmt.Println("[Adaptor Start]")
	oldHttpClient := &adaptor.OldHttpClient{}
	oldHttpClientAdaptor := &adaptor.OldHttpClientAdaptor{
		Client: oldHttpClient,
	}
	response := oldHttpClientAdaptor.Get("http://www.google.com")
	fmt.Println(response)
	fmt.Println("[Adaptor Done]")

	fmt.Println("[Facade Start]")
	mediaFacade := facade.NewMediaFacade()
	mediaFacade.Play()
	fmt.Println("[Facade Done]")

	fmt.Println("[Facade Start]")
	mediaFacade.TureOffNotification()
	fmt.Println("[Facade Done]")

	fmt.Println("[Decorator Start]")
	pizza := &decorator.NormalPizza{}
	pizza.GetPrice()
	tomatoPizza := &decorator.TomatoToppingPizza{
		Pizza: pizza,
	}
	tomatoPizza.GetPrice()
	cheesePizza := &decorator.CheeseToppingPizza{
		Pizza: pizza,
	}
	cheesePizza.GetPrice()
	fmt.Println("[Decorator Done]")

	fmt.Println("[Decorator Start]")
	coffee := &decorator.NormalCoffee{}
	fmt.Println(coffee.Description())
	coffee.Cost()
	latte := &decorator.MilkDecorator{
		Coffee: coffee,
	}
	fmt.Println(latte.Description())
	latte.Cost()
	sugaredCoffee := &decorator.SugarDecorator{
		Coffee: coffee,
	}
	fmt.Println(sugaredCoffee.Description())
	sugaredCoffee.Cost()
	fmt.Println("[Decorator Done]")

	fmt.Println("[Proxy] Start")
	proxyService := proxy.NewProxyService()
	proxyService.Request("admin")
	proxyService.Request("user1")
	proxyService.Request("user2")
	fmt.Println("[Proxy] Done")

	fmt.Println("[Proxy] Start")
	proxyImage := &proxy.ProxyImage{
		ImageFile: []string{"big_image_1.jpg", "big_image_2.jpg"},
	}
	proxyImage.Display()
	fmt.Println("[Proxy] Done")

	fmt.Println("[Proxy] Start")
	subscriber1 := observer.NewSubscribe("TWOBAO")
	subscriber1.Update("HI TWOBAO!")
	subscriber2 := observer.NewSubscribe("JOANNE")
	subscriber2.Update("HI JOANNE!")
	subscriber3 := observer.NewSubscribe("CTCHEN")
	subscriber3.Update("HI CTCHEN!")
	agency := observer.NewAgency()
	agency.Attach(subscriber1)
	agency.Attach(subscriber2)
	agency.Attach(subscriber3)
	agency.SetNews("HELLO SUBSCRIBER!!!")
	agency.Notify()
	fmt.Println("[Proxy] Done")

	fmt.Println("[Strategy] Done")
	linepay := strategy.NewLinePayment("0922222222")
	creditCard := strategy.NewCreditCardPayment("2222-2222-2222-2222")
	payPal := strategy.NewPaypalPayment("twobao@twobao.com")
	strategy := strategy.NewPaymentContext(linepay)
	strategy.Strategy.Pay(1000)
	strategy.SetPayment(creditCard)
	strategy.Strategy.Pay(2000)
	strategy.SetPayment(payPal)
	strategy.Strategy.Pay(3000)
	fmt.Println("[Strategy] Done")
}
