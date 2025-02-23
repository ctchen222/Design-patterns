// 1. 抽象產品（Car 和 Motorcycle）
abstract class Car {
  abstract drive(): void
}

abstract class Motorcycle {
  abstract ride(): void
}

class SUV extends Car {
  drive() {
    console.log('Driving an SUV!')
  }
}

class Sedan extends Car {
  drive() {
    console.log('Driving a Sedan!')
  }
}

class SportBike extends Motorcycle {
  ride() {
    console.log('Riding a Sport Bike!')
  }
}

class Cruiser extends Motorcycle {
  ride() {
    console.log('Riding a Cruiser!')
  }
}

abstract class VehicleFactory {
  abstract createCar(): Car
  abstract createMotorcycle(): Motorcycle
}

export class CarFactory extends VehicleFactory {
  createCar(): Car {
    return new Sedan()
  }

  createMotorcycle(): Motorcycle {
      return new SportBike()
  }
}

export class MotorcycleFactory extends VehicleFactory {
  createCar(): Car {
    return new SUV()
  }

    createMotorcycle(): Motorcycle {
        return new Cruiser()
    }
}

export class Client {
  car: Car
  motorcycle: Motorcycle

  constructor(factory: VehicleFactory) {
    this.car = factory.createCar()
    this.motorcycle = factory.createMotorcycle()
  }

  testDrive() {
    this.car.drive()
    this.motorcycle.ride()
  }
}