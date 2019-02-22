package main

import (
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	emp := &Employee{
		Name: "Gangadhar",
		Age:  35,
	}

	ba, err := proto.Marshal(emp)
	if err != nil {
		log.Fatal("Error during proto Marshal ", err)
	}
	log.Println(string(ba))
	var employee *Employee
	employee = &Employee{}
	err = proto.Unmarshal(ba, employee)
	if err != nil {
		log.Println("Error during Unmarshalling...", err)
	}
	log.Println(employee.Age)
	log.Println(employee.Name)
}
