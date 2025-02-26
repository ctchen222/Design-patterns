// Either interface or abstract class can do the same thing
// interface TaxCalculator {
//     calculateTax(amount: number): number
// }

abstract class TaxCalculator {
    abstract calculateTax(amount: number) : number
}

class USTaxcalculator extends TaxCalculator {
    calculateTax(amount: number) {
        return amount * 0.1
    }
}

class EUTaxcalculator extends TaxCalculator {
    calculateTax(amount: number) {
        return amount * 0.2
    }
}

class JPTaxcalculator extends TaxCalculator {
    calculateTax(amount: number) {
        return amount * 0.08
    }
}

export class TaxCalculatorFactory {
    static createTaxCalculator(country: string) {
        switch (country) {
            case "US":
                return new USTaxcalculator()
            case "EU":
                return new EUTaxcalculator()
            case "JP":
                return new JPTaxcalculator()
            default:
                throw new Error("Unsupported country")
        }
    }
}