import { BurgerFactory } from "./Javascript/Factory/Factory-basic";
import { TaxCalculatorFactory } from "./Javascript/Factory/TaxFactory";
import { TaxCalculatorFactoryWithStrategy } from "./Javascript/Factory/TaxFactoryWIthStrategy";
import {
  CarFactory,
  MotorcycleFactory,
} from "./Javascript/AbstractFactory/absctract-factory";
import { Client } from "./Javascript/AbstractFactory/absctract-factory";
import { ComputerBuilder } from "./Javascript/Builder/Builder-basic";
import { Director } from "./Javascript/Builder/Builder-director";

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