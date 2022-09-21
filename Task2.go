package main
import "fmt"

type employee struct {
  name string
  salary int
  position string

}
type company struct {
  companyName string
  employees []employee
}

func main(){

  Emp1 := employee{
    name: "Arsal",
    salary: 1000,
    position: "CEO",
  }
  Emp2 := employee{
    name: "Ahmad",
    salary: 500,
    position: "CTO",
  }
  Emp3 := employee{
    name: "Jawad",
    salary: 250,
    position: "CFO",
  }
  emp := []employee{Emp1, Emp2, Emp3}

  comp := company{
    companyName: "Google",
    employees: emp,
  }

  fmt.Printf("%v has %v", comp.companyName, comp.employees)
}
