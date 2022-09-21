package main

import(
"fmt"
"strings"
"crypto/sha256"
//"reflect"
)
type Student struct {
  rollnumber int
  name string
  address string
  subject []string
}
func NewStudent (rollno int, name string, address string, subject []string) *Student{
  s := new(Student)
  s.rollnumber = rollno
  s. name = name
  s.address = address
  s.subject = subject
  return s
}

type StudentList struct {
  list []*Student // change *Student to student
}

func (ls *StudentList) CreateStudent (rollno int, name string, address string, subject []string) *Student {
  st := NewStudent (rollno, name, address, subject)
  ls.list = append(ls.list, st) // change here to *st from st
  return st
}

func CalculateHash(key interface{}) [32]byte {
  byteKey := []byte(fmt.Sprintf("%v", key.(interface{})))
  return sha256.Sum256(byteKey)
}

func main() {
  student := new(StudentList)
  emp := []string{"Maths", "English"}
  emp1 := []string{"Maths", "Science", "Urdu"}
  student.CreateStudent (24, "Asim", "АAAAA", emp)
  student.CreateStudent (25, "Naveed", "BBBBBB", emp1)
  for no:=0; no<2; no=no+1{
    fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), no, strings.Repeat("=", 25))
    fmt.Println("Student rollno: ", student.list[no].rollnumber)
    fmt.Println("Student name: ", student.list[no].name)
    fmt.Println("Student address: ", student.list[no].address)
    fmt.Println("Student Subjects: ", student.list[no].subject)
    sha := CalculateHash(*student.list[no])
    fmt.Println("Hash of This List: ", sha)
  }

}
