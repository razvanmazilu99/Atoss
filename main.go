package main

import (
	ent "atoss/pb1/entity"
	"fmt"
)

func main() {

	// location := ent.Location{Longitude: 20, Latitude: 40}

	// person := ent.Employee{Name: "Razvan", Age: 21, Location: location, Gender: "Male"}
	// person1 := ent.Employee{Name: "Andreas", Age: 26, Location: location, Gender: "Male"}
	// person2 := ent.Employee{Name: "Denisa", Age: 22, Location: location, Gender: "Female"}

	// employees := []ent.Employee{person, person1, person2}

	// company := ent.Company{Employees: employees, Location: location, Name: "Atoss"}

	// fmt.Println(location)
	// fmt.Println(person)
	// fmt.Println(company)

	companies := ent.InitCompanies()

	//fmt.Println(companies)

	var name, city string

	fmt.Print("Company name: ")
	fmt.Scanln(&name)

	fmt.Print("City: ")
	fmt.Scanln(&city)

	//name := "Atoss"
	//city := "Timisoara"

	fmt.Println("There are", CountOfEmployees(companies, name, city), "employees that work at", name, "and live in", city)
}

func CountOfEmployees(companies []ent.Company, name string, city string) int {

	sum := 0

	for _, company := range companies {
		if company.Name == name {
			for _, employee := range company.Employees {
				if employee.GetCity() == city {
					sum++
				}
			}
		}
	}

	return sum
}
