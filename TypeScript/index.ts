import { BurgerFactory } from "./Creational Patterns/Factory/Factory-basic";
import { TaxCalculatorFactory } from "./Creational Patterns/Factory/TaxFactory";
import { TaxCalculatorFactoryWithStrategy } from "./Creational Patterns/Factory/TaxFactoryWIthStrategy";
import {
  CarFactory,
  MotorcycleFactory,
} from "./Creational Patterns/AbstractFactory/absctract-factory";
import { Client } from "./Creational Patterns/AbstractFactory/absctract-factory";
import { ComputerBuilder } from "./Creational Patterns/Builder/Builder-basic";
import { Director } from "./Creational Patterns/Builder/Builder-director";
import { MediaFacade } from "./Structural Patterns/Facade/facade";
import { OldPaymentSystem, PaymentSystemAdaptor } from './Structural Patterns/Adaptor/adaptor';
import { HttpClientAdaptor, OldHttpClient } from './Structural Patterns/Adaptor/adaptor-exI';
import { MilkDecorator, SimpleCoffee, SugarDecorator } from "./Structural Patterns/Decorator/decorator";
import { APIProxy, type APIService } from "./Structural Patterns/Proxy/proxy-exI";
import { NewsAgency, Subscriber } from "./Behavioral Patterns/Observer/observer";
import { LinePayPayment, CreditCartPayment, PaymentContext, PaypalPayment } from "./Behavioral Patterns/Strategy/strategy";
import { ProxyService, type Service } from "./Structural Patterns/Proxy/proxy-exII";

// Factory basic
console.log("[Factory Start]")
const burgurFactory = new BurgerFactory();
const cheeseBurgur = burgurFactory.createBeefBurger();
cheeseBurgur.print();
console.log("[Factory Done]")

// Tax factory
console.log("[Factory Start]")
const TaxCalculator = TaxCalculatorFactory.createTaxCalculator("EU");
console.log(`Tax: ${TaxCalculator.calculateTax(100)}`);
console.log("[Factory Done]")

// Tax factory with strategy pattern
console.log("[Factory Start]")
const StrategiedTaxCalculator =
  TaxCalculatorFactoryWithStrategy.createTaxCalculator("JP");
console.log(`Tax: ${StrategiedTaxCalculator.calculateTax(100)}`);
console.log("[Factory Done]")

// Abstract factory
console.log("[Abstract Factory Start]")
const carFactory = new CarFactory();
const motorcycleFactory = new MotorcycleFactory();
const client_one = new Client(carFactory);
client_one.testDrive();
const client_two = new Client(motorcycleFactory);
client_two.testDrive();
console.log("[Abstract Factory Done]")

// Builder
console.log("[Builder Start]")
const builder = new ComputerBuilder();
const computer = builder
  .setCPU("Intel i9")
  .setRAM("32GB")
  .setStorage("2TB SSD")
  .setGraphics("NVIDIA 4090Ti")
  .build();
console.log(computer.toString());
console.log("[Builder Done]")

// Builder director
console.log("[Builder Director Start]")
const highendBuilder = new ComputerBuilder()
const highendComputer = Director.createHighEndComputer(highendBuilder)
console.log(highendComputer.toString())
const budgetBuilder = new ComputerBuilder()
const budgetComputer = Director.createBudgetComputer(budgetBuilder)
console.log(budgetComputer.toString())
console.log("[Builder Director Done]")

// Facade
console.log("[Facade Start]")
const mediaFacade = new MediaFacade() 
mediaFacade.play()
mediaFacade.turnOffNotification()
console.log("[Facade Done]")

// Adaptor
console.log("[Adaptor Start]")
const oldPaymentSystem = new OldPaymentSystem()
const newPaymentSystem = new PaymentSystemAdaptor(oldPaymentSystem)
newPaymentSystem.processPayment(100)
console.log("[Adaptor Done]")

// Adaptor ex
console.log("[Adaptor Start]")
const oldHttpClient = new OldHttpClient()
const HttpClient = new HttpClientAdaptor(oldHttpClient)
HttpClient.get("TEST URL")
console.log("[Adaptor Done]")

// Decorator
console.log("[Decorator Start]")
const coffee = new SimpleCoffee()
console.log(`${coffee.description()} costs ${coffee.cost()}`)
const milkCoffee = new MilkDecorator(coffee)
console.log(`${milkCoffee.description()} costs ${milkCoffee.cost()}`)
const sugarCoffee = new SugarDecorator(coffee)
console.log(`${sugarCoffee.description()} costs ${sugarCoffee.cost()}`)
console.log("[Decorator Done]")

// Proxy
console.log("[Proxy Start]")
const apiService: APIService = new APIProxy()
apiService.request("/test")
apiService.request("/test")
// setTimeout(() => {
//   apiService.request("/test")
// }, 3100);
console.log("[Proxy Done]")

// Proxy
console.log("[Proxy Start]")
const proxyServer: Service = new ProxyService()
proxyServer.request("admin")
proxyServer.request("user1")
proxyServer.request("user2")
console.log("[Proxy Done]")

// Observer
console.log("[Observer Start]")
const agency = new NewsAgency()
const subscriber1 = new Subscriber("CT")
const subscriber2 = new Subscriber("TWOBAO")
const subscriber3 = new Subscriber("Pania")
agency.attach(subscriber1)
agency.attach(subscriber2)
agency.attach(subscriber3)
agency.setNews("TWOBAO LOVES U!")
// agency.notify()
console.log("[Observer Done]")

// Strategy
console.log("[Strategy Start]")
const creditCartPayment = new CreditCartPayment("1234567812345678")
const linepayPayment = new LinePayPayment("0922222222")
const paypalPayment = new PaypalPayment("twobao222@test.com")
const paymentContext = new PaymentContext(creditCartPayment)
paymentContext.executePayment(10000)
paymentContext.setStrategy(linepayPayment)
paymentContext.executePayment(50)
paymentContext.setStrategy(paypalPayment)
paymentContext.executePayment(200000)
console.log("[Strategy Done]")
