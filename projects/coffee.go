package main

import (
  "fmt"
  "time"
  "math/rand"
)


type order struct {
  coffee string
}

type machine struct {
  name string
}

func (m machine) use(r resource) {
  fmt.Println("Using ",m.name," with ",r.name)
  time.Sleep(time.Millisecond)
}

type resource struct {
  name string
}

type resourcePool struct {
  grinders chan machine
  pressers chan machine
  steamers chan machine
  coffeeBeans chan resource
  milk chan resource
}

func (rP *resourcePool) order(toOrder string) {
  switch {
  case toOrder == "coffeeBeans":
    for i := 0; i < 10; i++ {
      r := resource{name: "coffeeBeans"}
      rP.coffeeBeans<-r
      //*rP.addBeans(resource{name: "coffeeBeans"})
    }
  case toOrder == "milk":
    for i := 0; i < 10; i++ {
      r := resource{name: "milk"}
      rP.milk<-r
    }
  default:
    return
  }
  return
}

func (rp *resourcePool) addMilk(r resource){
  rp.milk<-r
  return
}

func (rp *resourcePool) addBeans(r resource){
  rp.coffeeBeans<-r
  return
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

func makeResourcesChannel(n int, name string) chan resource {
  resChan := make(chan resource,n)
  for i := 0; i < n; i++ {
    resChan<-resource{name: name}
  }
  return resChan
}

func grindCoffeeBeans(grinders chan machine,coffeeBeans chan resource) resource {
  grinder := <-grinders
  coffeeBean := <-coffeeBeans
  grinder.use(coffeeBean)
  grinders<-grinder
  return resource{name:"groundCoffee"}
}

func makeEspresso(pressers chan machine,groundCoffee resource) string {
  presser := <-pressers
  presser.use(groundCoffee)
  pressers<-presser
  return "espresso"
}

func steamMilk(steamers chan machine, milkSource chan resource) string {
  steamer := <-steamers
  milk := <-milkSource
  steamer.use(milk)
  steamers<-steamer
  return "steamedMilk"
}

func makeLatte(steamedMilk string,espresso string) string {
  time.Sleep(time.Second * 2)
  return "latte"
}

func initBarista(orders *chan order, rP *resourcePool) {
  for ord := range *orders {
    switch {
    case ord.coffee == "espresso":
      groundCoffee := grindCoffeeBeans(rP.grinders,rP.coffeeBeans)
      espresso := makeEspresso(rP.pressers,groundCoffee)
      fmt.Println("Finished",espresso)
    case ord.coffee == "latte":
      groundCoffee := grindCoffeeBeans(rP.grinders,rP.coffeeBeans)
      espresso := makeEspresso(rP.pressers,groundCoffee)
      steamedMilk := steamMilk(rP.steamers,rP.milk)
      latte := makeLatte(steamedMilk,espresso)
      fmt.Println("Finished",latte)
    }
  }
}

// func runCoffeeShop(iterations int) {
//   finished := false
// }

func milkManager(rP *resourcePool,running *bool){
  <-rP.milk
  fmt.Println("Milk exhausted")
  name := "milk"
  rP.order(name)
  go milkManager(rP,running)
  return
}

func coffeeBeansManager(rP *resourcePool,running *bool){
  <-rP.coffeeBeans
  fmt.Println("Coffee beans exhausted")
  name := "coffeeBeans"
  rP.order(name)
  go coffeeBeansManager(rP,running)
  return
}

func generateOrders(orders *chan order,running *bool,n int){
  for i := 0; i < n; i++ {
    choices := []string{"latte","espresso"}
    newOrder := order{coffee: choices[rand.Intn(2)]}
    *orders<-newOrder
  }
}

func orderSimulator(orders *chan order,running *bool){
  for *running {
    go generateOrders(orders,running,20)
    time.Sleep(time.Second*5)
  }
}

func manager(rP *resourcePool,running *bool,orders *chan order){
  for *running {
    if len(rP.milk) < 1 {
      fmt.Println("Out of milk, ordering...")
      go rP.order("milk")
      time.Sleep(time.Second * 2)
    }
    if len(rP.coffeeBeans) < 1 {
      fmt.Println("Out of coffee beans, ordering...")
      go rP.order("coffeeBeans")
      time.Sleep(time.Second * 2)
    }
    if len(*orders) < 1 {
      fmt.Println("No orders to complete...")
      *running = false
    }
  }
}

func main(){
  // main stuff!

  var running bool = true

  orders := makeOrderChannel(10)

  grinders := makeNewMachineChannel(2,"grinder")
  pressers := makeNewMachineChannel(2,"presser")
  steamers := makeNewMachineChannel(2,"steamer")
  coffeeBeans := makeResourcesChannel(10,"coffeeBeans")
  milk := makeResourcesChannel(10,"milk")

  rP := resourcePool{
    grinders: grinders,
    pressers: pressers,
    steamers: steamers,
    coffeeBeans: coffeeBeans,
    milk: milk,
  }

  go orderSimulator(&orders,&running)
  time.Sleep(time.Second * 5)

  go manager(&rP,&running,&orders)

  go initBarista(&orders,&rP)
  go initBarista(&orders,&rP)
  go initBarista(&orders,&rP)

  for running {

  }

  fmt.Println("Orders exhausted")


}
