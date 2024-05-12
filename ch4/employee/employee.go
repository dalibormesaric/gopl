package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Adress    string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

type Employee2 struct {
	ID           int
	Name, Adress string
	DoB          time.Time
	Position     string
	Salary       int
	ManagerID    int
}

func main() {
	var dilbert Employee
	fmt.Printf("Salary: %v\n", dilbert)
	fmt.Printf("Salary: %+v\n", dilbert)

	dilbert.Salary -= 5000 // demoted, for writing too few lines of code
	fmt.Printf("Salary: %d\n", dilbert.Salary)

	position := &dilbert.Position
	*position = "Senior " + *position // promoted, for outsourcing to Elbonia
	fmt.Printf("%s\n", dilbert.Position)

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	fmt.Printf("%s\n", dilbert.Position)

	(*employeeOfTheMonth).Position += " etc." // last statement is equal to this
	fmt.Printf("%s\n", employeeOfTheMonth.Position)

	fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"

	id := dilbert.ID
	EmployeeByID(id).Salary = 0 // fired for... no real reason
}

func EmployeeByID(id int) *Employee {
	return &Employee{ID: id}
}
