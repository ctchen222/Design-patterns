class Burger {
  ingredients: string[]

  constructor(ingredients: string[]) {
    this.ingredients = ingredients
  }

  print() {
    console.log(this.ingredients)
  }
}

export class BurgerFactory {
  createCheeseBurger() {
    const ingredients = ['cheese', 'beef', 'bread']
    return new Burger(ingredients)
  }

  createBeefBurger() {
    const ingredients = ['beef', 'bread']
    return new Burger(ingredients)
  }

  createChickenBurger() {
    const ingredients = ['chicken', 'bread']
    return new Burger(ingredients)
  }
}
