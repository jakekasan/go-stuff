package main

import (
  "fmt"
  "time"
)


type order struct {
  coffee string
  cost int
}

type machine struct {
  name string
}

func (m machine) use() {
  fmt.Println("Using ",m.name)
  time.Sleep(time.Second)
}

type resources struct {
  name string
}

type resourcePool struct {
  grinders chan machines
  pressers chan machines
  steamers chan machines
  coffeeBeans chan resources
  milk chan resources
}

// make order channel

func makeOrderChannel(n int) chan order {
  return make(chan order,n)
}

func makeNewMachine(name string) machine {
  return machine{name: name}
}

func makeNewMachineChannel(n int,name string) chan machine {
  machines := make(chan machine,n)
  for i := 0; i < n; i++ {
    newMachine := makeNewMachine(name)
    machines<-newMachine
  }
  return machines
}

func grindCoffeeBeans(grinders chan machine) string {
  grinder := <-grinders
  grinder.use()
  grinders<-grinder
  return "groundCoffee"
}

func makeEspresso(groundCoffee string,pressers chan machine) string {
  presser := <-pressers
  presser.use()
  pressers<-presser
  return "espresso"
}

func steamMilk(steamers chan machine) string {
  steamer := <-steamers
  steamer.use()
  steamers<-steamer
  return "steamedMilk"
}

func makeLatte(steamedMilk string,espresso string) string {
  time.Sleep(time.Second * 2)
  return "latte"
}

func initBarista(orders chan order, resources resourcePool) {
  for ord := range orders {
    switch {
    case ord.coffee == "espresso":
      groundCoffee := grindCoffeeBeans(resources.grinders)
      espresso := makeEspresso(groundCoffee,resources.pressers)
      fmt.Println("Finished",espresso)
    case ord.coffee == "latte":
      groundCoffee := grindCoffeeBeans(resources.grinders)
      espresso := makeEspresso(groundCoffee,resources.pressers)
      steamedMilk := steamMilk(resources.steamers)
      latte := makeLatte(steamedMilk,espresso)
      fmt.Println("Finished",latte)
    }
  }
}

func runCoffeeShop(iterations int) {
  finished := false
  go initBarista()
  go initBarista()
  for !finished {

  }
}


func main(){
  // main stuff!

  orders := makeOrderChannel(10)

  grinders := makeNewMachineChannel(2,"grinder")
  pressers := makeNewMachineChannel(2,"presser")
  steamers := makeNewMachineChannel(2,"steamer")
  coffeeBeans :=

  for range [5]int{}{
    orders<-order{coffee: "latte"}
  }

  for range [5]int{}{
    orders<-order{coffee: "espresso"}
  }

  go initBarista(orders,grinders,pressers,steamers)
  go initBarista(orders,grinders,pressers,steamers)
  go initBarista(orders,grinders,pressers,steamers)
  go initBarista(orders,grinders,pressers,steamers)



  time.Sleep(time.Second * 60)


}
