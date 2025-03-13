// abstract class NewPaymentGateway {
//     abstract processPayment(amount: number): void;
// }
interface NewPaymentGateway {
    processPayment(amount: number): void
}

export class OldPaymentSystem {
    sendPayment(amount: number) {
        console.log(`OldPayment system pay $${amount}`)
    }
}

export class PaymentSystemAdaptor implements NewPaymentGateway {
    private OldPaymentSystem: OldPaymentSystem

    constructor(OldPaymentSystem: OldPaymentSystem) {
        this.OldPaymentSystem = OldPaymentSystem
    }

    processPayment(amount: number): void {
        console.log("[Adaptor] Transfer...")
        this.OldPaymentSystem.sendPayment(amount)
    }
}