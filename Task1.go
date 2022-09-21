package main
import "fmt"

type Doctor struct {
  number int
  docName string
  patients []string
}

func main() {
  Adoc := Doctor{
    number: 500,
    docName: "Naveed",
    patients: []string{"Syed", "Arsal", "Ahmad"},
  }

  Task1(Adoc)

}

func Task1(doc Doctor) {
// code to be executed
  fmt.Printf("Doctor ID = %v, Doctor name = %v, Patients = %v", doc.number, doc.docName, doc.patients)
}
