import { BurgerFactory } from "./Javascript/Creational Patterns/Factory/Factory-basic";
import { TaxCalculatorFactory } from "./Javascript/Creational Patterns/Factory/TaxFactory";
import { TaxCalculatorFactoryWithStrategy } from "./Javascript/Creational Patterns/Factory/TaxFactoryWIthStrategy";
import {
  CarFactory,
  MotorcycleFactory,
} from "./Javascript/Creational Patterns/AbstractFactory/absctract-factory";
import { Client } from "./Javascript/Creational Patterns/AbstractFactory/absctract-factory";
import { ComputerBuilder } from "./Javascript/Creational Patterns/Builder/Builder-basic";
import { Director } from "./Javascript/Creational Patterns/Builder/Builder-director";
import { MediaFacade } from "./Javascript/Structural Patterns/Facade/facade";
import { OldPaymentSystem, PaymentSystemAdaptor } from './Javascript/Structural Patterns/Adaptor/adaptor';
import { HttpClientAdaptor, OldHttpClient } from './Javascript/Structural Patterns/Adaptor/adaptor-exI';
import { MilkDecorator, SimpleCoffee, SugarDecorator } from "./Javascript/Structural Patterns/Decorator/decorator";
import { APIProxy, type APIService } from "./Javascript/Structural Patterns/Proxy/proxy";
import { NewsAgency, Subscriber } from "./Javascript/Behavioral Patterns/Observer/observer";
import { LinePayPayment, CreditCartPayment, PaymentContext, PaypalPayment } from "./Javascript/Behavioral Patterns/Strategy/strategy";

// Factory basic
const burgurFactory = new BurgerFactory();
const cheeseBurgur = burgurFactory.createBeefBurger();
cheeseBurgur.print();

// Tax factory
const TaxCalculator = TaxCalculatorFactory.createTaxCalculator("EU");
console.log(`Tax: ${TaxCalculator.calculateTax(100)}`);

// Tax factory with strategy pattern
const StrategiedTaxCalculator =
  TaxCalculatorFactoryWithStrategy.createTaxCalculator("JP");
console.log(`Tax: ${StrategiedTaxCalculator.calculateTax(100)}`);

// Abstract factory
const carFactory = new CarFactory();
const motorcycleFactory = new MotorcycleFactory();
const client_one = new Client(carFactory);
client_one.testDrive();
const client_two = new Client(motorcycleFactory);
client_two.testDrive();

// Builder
const builder = new ComputerBuilder();
const computer = builder
  .setCPU("Intel i9")
  .setRAM("32GB")
  .setStorage("2TB SSD")
  .setGraphics("NVIDIA 4090Ti")
  .build();
console.log(computer.toString());

// Builder director
const highendBuilder = new ComputerBuilder()
const highendComputer = Director.createHighEndComputer(highendBuilder)
console.log(highendComputer.toString())
const budgetBuilder = new ComputerBuilder()
const budgetComputer = Director.createBudgetComputer(budgetBuilder)
console.log(budgetComputer.toString())

// Facade
const mediaFacade = new MediaFacade() 
mediaFacade.play()
mediaFacade.turnOffNotification()

// Adaptor
const oldPaymentSystem = new OldPaymentSystem()
const newPaymentSystem = new PaymentSystemAdaptor(oldPaymentSystem)
newPaymentSystem.processPayment(100)

// Adaptor ex
const oldHttpClient = new OldHttpClient()
const HttpClient = new HttpClientAdaptor(oldHttpClient)
HttpClient.get("TEST URL")

// Decorator
const coffee = new SimpleCoffee()
console.log(`${coffee.description()} costs ${coffee.cost()}`)
const milkCoffee = new MilkDecorator(coffee)
console.log(`${milkCoffee.description()} costs ${milkCoffee.cost()}`)
const sugarCoffee = new SugarDecorator(coffee)
console.log(`${sugarCoffee.description()} costs ${sugarCoffee.cost()}`)

// Proxy
const apiService: APIService = new APIProxy()
apiService.request("/test")
apiService.request("/test")
// setTimeout(() => {
//   apiService.request("/test")
// }, 3100);

// Observer
const agency = new NewsAgency()
const subscriber1 = new Subscriber("CT")
const subscriber2 = new Subscriber("TWOBAO")
const subscriber3 = new Subscriber("Pania")
agency.attach(subscriber1)
agency.attach(subscriber2)
agency.attach(subscriber3)
agency.setNews("TWOBAO LOVES U!")
// agency.notify()

// Strategy
const creditCartPayment = new CreditCartPayment("1234567812345678")
const linepayPayment = new LinePayPayment("0922222222")
const paypalPayment = new PaypalPayment("twobao222@test.com")
const paymentContext = new PaymentContext(creditCartPayment)
paymentContext.executePayment(10000)
paymentContext.setStrategy(linepayPayment)
paymentContext.executePayment(50)
paymentContext.setStrategy(paypalPayment)
paymentContext.executePayment(200000)
