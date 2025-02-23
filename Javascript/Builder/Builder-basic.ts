class Computer {
    cpu: string
    ram: string
    storage: string
    graphics: string

    constructor() {
        this.cpu = ""
        this.ram = ""
        this.storage = ""
        this.graphics = ""
    }

    toString() {
        return `Computer with CPU: ${this.cpu}, RAM: ${this.ram}, STORAGE: ${this.storage}, GRAPHICS: ${this.graphics}`
    }
}

export class ComputerBuilder extends Computer {
    computer: any

    constructor() {
        super()
    }

    setCPU(cpu: string) {
        this.computer.cpu = cpu
        return this
    }

    setRAM(ram: string) {
        this.computer.ram = ram
        return this
    }

    setStorage(storage: string) {
        this.computer.storage = storage
        return this
    }

    setGraphics(grapics: string) {
        this.computer.graphics = grapics  
        return this
    }

    build() {
        return this.computer
    }
}