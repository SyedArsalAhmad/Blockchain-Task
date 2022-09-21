package main

import(
"fmt"
"strings"
//"reflect"
)
type Student struct {
  rollnumber int
  name string
  address string
}
func NewStudent (rollno int, name string, address string) *Student{
  s := new(Student)
  s.rollnumber = rollno
  s. name = name
  s.address = address
  return s
}

type StudentList struct {
  list []*Student // change *Student to student
}

func (ls *StudentList) CreateStudent (rollno int, name string, address string) *Student {
  st := NewStudent (rollno, name, address)
  ls.list = append(ls.list, st) // change here to *st from st
  return st
}

func main() {
  student := new(StudentList)
  student.CreateStudent (24, "Asim", "АAAAA")
  student.CreateStudent (25, "Naveed", "BBBBBB")
  for no:=0; no<2; no=no+1{
    fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), no, strings.Repeat("=", 25))
    fmt.Println("Student rollno: ", student.list[no].rollnumber)
    fmt.Println("Student name: ", student.list[no].name)
    fmt.Println("Student address: ", student.list[no].address)
  }

}
