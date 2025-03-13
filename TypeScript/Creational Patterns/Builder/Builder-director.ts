import type { ComputerBuilder } from "./Builder-basic";

export class Director {
    static createHighEndComputer(builder: ComputerBuilder ) {
        return builder
            .setCPU("Intel i9")
            .setRAM("128GB")
            .setStorage("16TB SSD")
            .setGraphics("NVIDIA 5090Ti")
            .build()
    }

    static createBudgetComputer(builder: ComputerBuilder) {
        return builder
            .setCPU("Intel i5")
            .setRAM("32GB")
            .setStorage("2TB SSD")
            .setGraphics("NVIDIA 3060")
            .build()
    }
}