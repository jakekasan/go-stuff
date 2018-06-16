package main

import "fmt"


type coffee struct {
  nameOfCoffee string
  price int
  requirements []string
  fulfilled []resource
}

type resource struct {
  nameOfResource string
  cost int
}

type order struct {
  coffees []coffee
  cost int
}


func generateOrder(name string) coffee {
  return coffee{nameOfCoffee: name,price: 10, requirements: []string{"water","ground coffee beans"},fulfilled: []resource{}}
}





func main(){
  // main stuff!

  resources := make([]resource,0)


  // make water resources
  for range [5]int{} {
    resources = append(resources,resource{nameOfResource: "water", cost: 1})
  }

  // make coffee resources
  for range [5]int{} {
    resources = append(resources,resource{nameOfResource: "ground coffee beans", cost: 5})
  }

  var espresso = coffee{nameOfCoffee: "espresso", price: 10, requirements: []string{"water","ground coffee beans"},fulfilled: []resource{}}

  fmt.Println(espresso.nameOfCoffee)

  // get orders

  pendingOrders := make([])

  for range [5]int{} {

  }




}
