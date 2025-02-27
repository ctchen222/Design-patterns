interface PaymentStrategy {
    pay(amount: number): void
}

export class CreditCartPayment implements PaymentStrategy {
    private cardNumber: string

    constructor(cardNumber: string) {
        this.cardNumber = cardNumber
    }

    pay(amount: number): void {
        console.log(`Pay ${amount} with card ${this.cardNumber}`)
    }
}

export class LinePayPayment implements PaymentStrategy {
    private mobile: string

    constructor(mobile: string) {
        this.mobile = mobile
    }

    pay(amount: number): void {
        console.log(`Pay ${amount} with LinePay ${this.mobile}`)
    }
}

export class PaypalPayment implements PaymentStrategy {
    private email: string

    constructor(email: string) {
       this.email = email
    }

    pay(amount: number): void {
        console.log(`Pay ${amount} with PayPal ${this.email}`)
    }
}

export class PaymentContext {
    private strategy: PaymentStrategy
    
    constructor(strategy: PaymentStrategy) {
       this.strategy = strategy 
    }

    setStrategy(strategy: PaymentStrategy) {
        this.strategy = strategy
    }

    executePayment(amount: number): void {
        this.strategy.pay(amount)
    }
}