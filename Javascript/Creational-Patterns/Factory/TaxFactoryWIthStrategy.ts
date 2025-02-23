// Since TaxFactory.ts may violate SOLID pricipal and lack of testability
// We apply Strategy pattern on top of TaxFactory.ts

abstract class TaxCalculator {
    abstract calculateTax(amount: number) : number
}

class USTaxCalculator extends TaxCalculator {
    calculateTax(amount: number): number {
        return amount * 0.1
    }
}

class EUTaxCalculator extends TaxCalculator {
    calculateTax(amount: number): number {
        return amount * 0.2
    }
}

class JPCalculator extends TaxCalculator {
    calculateTax(amount: number): number {
        return amount * 0.08
    }
}

const calculatorStrategy: { [key: string]: new () => TaxCalculator } = {
    "US": USTaxCalculator,
    "EU": EUTaxCalculator,
    "JP": JPCalculator
}

export class TaxCalculatorFactoryWithStrategy {
    static createTaxCalculator(country: string) {
        const countryClass = calculatorStrategy[country]
        if (!country) {
            throw new Error("Unsupported country")
        }
        return new countryClass()
    }
}