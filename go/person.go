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

// assumping list in sorted
func isPersonTwiceOldAsOthersV2(person []Person) bool {

	if len(person) < 2 {
		fmt.Printf("Need minimum 2 person's age to continue\n")
		return false
	}

	seen := make(map[int]bool, 15)

	for index, _ := range person {
		if seen[person[index].Age] {
			return true
		}
		seen[person[index].Age*2] = true
	}
	return false

}

// with following constraints.
// Age is an integer with values that range from 0 to 50
// you may only use variables of numeric types (no array/slice/map)
// must have linear time complexity with just a single loop
// you may not use addition, subtraction, multiplication, division or modulo operators
// the function may not consist of more than 15 lines of code
// the function may not call other functions

// for 0 to 50, there will be only 25 distinct pairs where one is the double of the other:
//(0, 0), (1, 2), (2,4),(3,6) ... (25, 50).
// doubel of 0 is 0 - so this is special case
func isPersonTwiceOldAsOthersV3(person []Person) bool {
	var flagSmalls, flagHalves, numZeroes int
	if len(person) < 2 {
		fmt.Printf("Need minimum 2 person's age to continue\n")
		return false
	}
	for index, _ := range person {
		age := person[index].Age
		if age == 0 {
			numZeroes++
		}
		if age > 0 && age < 26 {
			flagSmalls |= (1 << age)
		}
		if age > 0 && (age&1) == 0 {
			flagHalves |= (1 << (age >> 1))
		}
	}
	return numZeroes > 1 || (flagSmalls&flagHalves) != 0
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
		Person{Age: 3},
		Person{Age: 6},
		Person{Age: 9},
	}
	fmt.Printf("%v\n", isPersonTwiceOldAsOthersV3(personList)) // returns true

	// test values
	personList = []Person{
		Person{Age: 2},
		Person{Age: 5},
		Person{Age: 8},
	}
	fmt.Printf("%v\n", isPersonTwiceOldAsOthersV3(personList)) // returns false

	// test values
	personList = []Person{
		Person{Age: 0},
		Person{Age: 5},
		Person{Age: 8},
		Person{Age: 0},
	}
	fmt.Printf("%v\n", isPersonTwiceOldAsOthersV3(personList)) // returns true, double of 0 is 0

	// test values
	personList = []Person{
		Person{Age: 2},
		Person{Age: 4},
	}
	fmt.Printf("%v\n", isPersonTwiceOldAsOthersV3(personList)) // returns true

	// test values
	personList = []Person{
		Person{Age: 2},
		Person{Age: 5},
		Person{Age: 3},
		Person{Age: 6},
	}
	fmt.Printf("%v\n", isPersonTwiceOldAsOthersV3(personList)) // returns true

	personList = []Person{
		Person{Age: 4},
		Person{Age: 4},
		Person{Age: 6},
		Person{Age: 6},
		Person{Age: 5},
	}
	fmt.Printf("%v\n", isPersonTwiceOldAsOthersV3(personList)) // returns false

	personList = []Person{
		Person{Age: 4},
		Person{Age: 9},
		Person{Age: 6},
		Person{Age: 18},
		Person{Age: 33},
	}
	fmt.Printf("%v\n", isPersonTwiceOldAsOthersV3(personList)) // returns true

	personList = []Person{
		Person{Age: 100},
		Person{Age: 50},
		Person{Age: 60},
		Person{Age: 100},
		Person{Age: 45},
		Person{Age: 200},
	}
	fmt.Printf("%v\n", isPersonTwiceOldAsOthersV3(personList)) // returns false

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
