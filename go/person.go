package main

import (
	"fmt"
	"strconv"
)

// Person object
type Person struct {
	Age int
}

func findDuplicates(person []Person) int {
	arr := make(map[string]bool, 15)
	for firstIndex, _ := range person {
		strPerson := strconv.Itoa(person[firstIndex].Age)
		if arr[strPerson] {
			return person[firstIndex].Age
		}
		arr[strPerson] = true
	}
	return 0
}

func isPersonTwiceOldAsOthers(person []Person) bool {

	if len(person) < 3 {
		fmt.Printf("Need minimum 3 person's age to continue\n") // own person + two other person
		return false
	}

	duplicate := findDuplicates(person)

	if duplicate == 0 {
		return false
	}

	for firstIndex, _ := range person {
		if person[firstIndex].Age == 2*duplicate {
			return true
		}
	}

	return false
}

func isPersonAtleastTwiceOldAsOthers(person []Person) bool {

	if len(person) == 0 {
		return false
	} else if len(person) == 1 {
		return true
	}

	// O(n) approach
	// first let us find the index of max person age
	personMaxAgeIndex := 0
	for index, _ := range person {
		if person[index].Age > person[personMaxAgeIndex].Age {
			personMaxAgeIndex = index
		}
	}
	// now check if the max person age is at least as twice as others
	for index, _ := range person {
		if personMaxAgeIndex != index && person[personMaxAgeIndex].Age < 2*person[index].Age {
			return false
		}
	}
	return true
}

func main() {

	var personList = []Person{}
	// test values
	personList = []Person{
		Person{Age: 20},
		Person{Age: 55},
		Person{Age: 45},
		Person{Age: 65},
		Person{Age: 55},
		Person{Age: 100},
	}
	fmt.Printf("%v\n", isPersonTwiceOldAsOthers(personList)) // returns false

	personList = []Person{
		Person{Age: 20},
		Person{Age: 50},
		Person{Age: 45},
		Person{Age: 65},
		Person{Age: 45},
		Person{Age: 90},
	}
	fmt.Printf("%v\n", isPersonTwiceOldAsOthers(personList)) // returns true

	// test values
	personList = []Person{
		Person{Age: 3},
		Person{Age: 2},
		Person{Age: 1},
		Person{Age: 4},
	}
	fmt.Printf("%v\n", isPersonAtleastTwiceOldAsOthers(personList)) // return false

	// test values
	personList = []Person{
		Person{Age: 6},
		Person{Age: 5},
		Person{Age: 4},
		Person{Age: 12},
	}
	fmt.Printf("%v\n", isPersonAtleastTwiceOldAsOthers(personList)) // return true
	// test values
	personList = []Person{
		Person{Age: 10},
		Person{Age: 5},
		Person{Age: 9},
		Person{Age: 20},
	}
	fmt.Printf("%v\n", isPersonAtleastTwiceOldAsOthers(personList)) // return true
}
