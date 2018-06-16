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
  return "latte"
}

func initBarista(orders chan order, grinders chan machine, pressers chan machine, steamers chan machine) {
  for ord := range orders {
    switch {
    case ord.coffee == "espresso":
      groundCoffee := grindCoffeeBeans(grinders)
      espresso := makeEspresso(groundCoffee,pressers)
      fmt.Println("Finished",espresso)
    case ord.coffee == "latte":
      groundCoffee := grindCoffeeBeans(grinders)
      espresso := makeEspresso(groundCoffee,pressers)
      steamedMilk := steamMilk(steamers)
      latte := makeLatte(steamedMilk,espresso)
      fmt.Println("Finished",latte)
    }
  }
}


func main(){
  // main stuff!

  orders := makeOrderChannel(10)

  grinders := makeNewMachineChannel(2,"grinder")
  pressers := makeNewMachineChannel(2,"presser")
  steamers := makeNewMachineChannel(2,"steamer")

  for range [5]int{}{
    orders<-order{coffee: "latte"}
  }

  for range [5]int{}{
    orders<-order{coffee: "espresso"}
  }

  go initBarista(orders,grinders,pressers,steamers)
  go initBarista(orders,grinders,pressers,steamers)



  time.Sleep(time.Second * 60)


}
