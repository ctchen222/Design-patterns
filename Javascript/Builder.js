class Burger {
  constructor() {
    this.buns = null
    this.patty = null
    this.cheese = null
  }

  setBuns(buns) {
    this.buns = buns
    return this
  }

  setPatty(patty) {
    this.patty = patty
    return this
  }

  setCheese(cheese) {
    this.cheese = cheese
    return this
  }
}

class BurgerBuilder {
  constructor() {
    this.burger = new Burger()
  }

  addBuns(buns) {
    this.burger.setBuns(buns)
    return this
  }

  addPatty(patty) {
    this.burger.setPatty(patty)
    return this
  }

  addCheese(cheese) {
    this.burger.setCheese(cheese)
    return this
  }

  build() {
    return this.burger
  }
}

const burger = new BurgerBuilder()
  .addBuns('sesame')
  .addPatty('beef')
  .addCheese('cheddar')
  .build()
console.log(burger)
