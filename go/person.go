package main

import (
	"fmt"
)

// Person object
type Person struct {
	Age int
}

// TODO - change from bruteforce approach to at least O(n) time
func isPersonTwiceOldAsOthers(person []Person) bool {

	if len(person) < 2 {
		fmt.Printf("Need minimum 2 person's age to continue\n")
		return false
	}

	for firstIndex, _ := range person {
		for secondIndex, _ := range person {
			if firstIndex != secondIndex &&
				(person[firstIndex].Age == 2*person[secondIndex].Age || person[firstIndex].Age*2 == person[secondIndex].Age) {
				return true
			}
		}
	}
	return false
}

func isPersonAtleastTwiceOldAsOthers(person []Person) bool {

	if len(person) < 2 {
		fmt.Printf("Need min 2 numbers\n")
		return false
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
		if personMaxAgeIndex != index && person[personMaxAgeIndex].Age >= 2*person[index].Age {
			return true
		}
	}
	return false
}

func main() {

	var personList = []Person{}

	// fist function: isPersonTwiceOldAsOthers

	// test values
	personList = []Person{
		Person{Age: 6},
		Person{Age: 3},
		Person{Age: 9},
	}
	fmt.Printf("%v\n", isPersonTwiceOldAsOthers(personList)) // returns true

	// test values
	personList = []Person{
		Person{Age: 2},
		Person{Age: 4},
		Person{Age: 8},
	}
	fmt.Printf("%v\n", isPersonTwiceOldAsOthers(personList)) // returns true

	// test values
	personList = []Person{
		Person{Age: 2},
		Person{Age: 4},
	}
	fmt.Printf("%v\n", isPersonTwiceOldAsOthers(personList)) // returns true

	// test values
	personList = []Person{
		Person{Age: 2},
		Person{Age: 5},
		Person{Age: 3},
		Person{Age: 6},
	}
	fmt.Printf("%v\n", isPersonTwiceOldAsOthers(personList)) // returns true

	personList = []Person{
		Person{Age: 4},
		Person{Age: 4},
		Person{Age: 6},
		Person{Age: 6},
		Person{Age: 5},
	}
	fmt.Printf("%v\n", isPersonTwiceOldAsOthers(personList)) // returns false

	personList = []Person{
		Person{Age: 100},
		Person{Age: 50},
		Person{Age: 60},
		Person{Age: 100},
		Person{Age: 45},
		Person{Age: 200},
	}
	fmt.Printf("%v\n", isPersonTwiceOldAsOthers(personList)) // returns true

	// second function: isPersonAtleastTwiceOldAsOthers

	// test values
	personList = []Person{
		Person{Age: 3},
		Person{Age: 4},
		Person{Age: 3},
		Person{Age: 4},
	}
	fmt.Printf("%v\n", isPersonAtleastTwiceOldAsOthers(personList)) // returns false

	// test values
	personList = []Person{
		Person{Age: 6},
		Person{Age: 5},
		Person{Age: 4},
		Person{Age: 12},
	}
	fmt.Printf("%v\n", isPersonAtleastTwiceOldAsOthers(personList)) // returns true
	// test values
	personList = []Person{
		Person{Age: 9},
		Person{Age: 6},
		Person{Age: 4},
	}
	fmt.Printf("%v\n", isPersonAtleastTwiceOldAsOthers(personList)) // returns true
	// test values
	personList = []Person{
		Person{Age: 9},
	}
	fmt.Printf("%v\n", isPersonAtleastTwiceOldAsOthers(personList))
}
