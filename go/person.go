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
	arr := make(map[string]bool,15)
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

	//arr := make(map[string]bool,15)
	duplicate := findDuplicates(person)

	if duplicate == 0 {
		return false
	}

	for firstIndex, _ := range person {
		if (person[firstIndex].Age == 2 * duplicate) {
			return true
		}
	}

	return false
}

// TODO - update bruteforce algo with efficient to have O(n)
func isPersonAtleastTwiceOldAsOthers(person []Person) bool {

	if len(person) < 2 {
		fmt.Printf("Need minimum 2 person's age to continue\n")
		return false
	}

	for firstIndex, _ := range person {
		for secondIndex, _ :=  range person {
			if (person[firstIndex].Age <= person[secondIndex].Age) {
				continue
			}
			if (person[firstIndex].Age >= 2 * person[secondIndex].Age) {
				return true
			}
		}
	}
	return false
}

func main() {
	
	var personList = []Person{}
	// test values
	personList = []Person {
		Person {Age: 20},
		Person {Age: 55},
		Person {Age: 45},
		Person {Age: 65},
		Person {Age: 55},
		Person {Age: 100},
	}
	fmt.Printf("%v\n",isPersonTwiceOldAsOthers(personList))  // returns false

	personList = []Person {
		Person {Age: 20},
		Person {Age: 50},
		Person {Age: 45},
		Person {Age: 65},
		Person {Age: 45},
		Person {Age: 90},
	}
	fmt.Printf("%v\n",isPersonTwiceOldAsOthers(personList))  // returns true

	// test values
	personList = []Person {
		Person {Age: 3},
		Person {Age: 2},
		Person {Age: 1},
		Person {Age: 4},
	}
	fmt.Printf("%v\n",isPersonAtleastTwiceOldAsOthers(personList)) // return true

	// test values
	personList = []Person {
		Person {Age: 6},
		Person {Age: 7},
		Person {Age: 8},
		Person {Age: 9},
	}
	fmt.Printf("%v\n",isPersonAtleastTwiceOldAsOthers(personList)) // return false
	// test values
	personList = []Person {
		Person {Age: 1},
		Person {Age: 9},
		Person {Age: 9},
		Person {Age: 9},
	}
	fmt.Printf("%v\n",isPersonAtleastTwiceOldAsOthers(personList)) // return true
}