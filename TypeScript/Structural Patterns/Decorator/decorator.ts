// abstract class Coffee {
//     abstract cost(): number
//     abstract description(): string
// }

interface Coffee {
    cost(): number
    description(): string
}

export class SimpleCoffee implements Coffee {
    cost(): number {
        return 60
    }

    description(): string {
        return "Coffee"
    }
}

// 封裝原始的物件 -> 當作decorator
class CoffeeDecorator implements Coffee {
    protected coffee: Coffee

    constructor(coffee: Coffee) {
        this.coffee = coffee
    }

    cost(): number {
       return this.coffee.cost() 
    }

    description(): string {
        return this.coffee.description()
    }
}

export class MilkDecorator extends CoffeeDecorator {    
    constructor(coffee: Coffee) {
        super(coffee)
    }

    cost(): number {
        return super.cost() + 10
    }
    
    description(): string {
        return super.description() + " + milk"
    }
}

export class SugarDecorator extends CoffeeDecorator {
    constructor(coffee: Coffee) {
        super(coffee)
    }

    cost(): number {
        return super.cost() + 5
    }

    description(): string {
        return super.description() + " + sugar"
    }
}