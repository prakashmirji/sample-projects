package main

import (
	"fmt"
	"strconv"
)

// Person object
type Person struct {
	Age int
}

func findDuplicatesV2(person []Person) int {
	start := 1
	end := len(person) - 1

	for start < end {
		// divide range 1 to n into lower and upper range
		// lower range : start to mid
		// upper range : mid+1 to end
		mid := start + ((end - start) / 2)
		lower_range_start, lower_range_end := start, mid
		upper_range_start, upper_range_end := mid+1, end

		// count num of items in the lower range
		items_in_lower_range := 0
		for firstIndex, _ := range person {
			if person[firstIndex].Age >= person[lower_range_start].Age && person[firstIndex].Age <= person[lower_range_end].Age {
				items_in_lower_range++
			}
		}

		distinct_integers_in_lower_range := lower_range_end - lower_range_start + 1

		if items_in_lower_range > distinct_integers_in_lower_range {
			start, end = lower_range_start, lower_range_end
		} else {
			start, end = upper_range_start, upper_range_end
		}
	}
	return person[start].Age
}

func findDuplicatesV1(person []Person) int {
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

	//duplicate := findDuplicatesV1(person)
	duplicate := findDuplicatesV2(person)

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

	// fist function: isPersonTwiceOldAsOthers

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
		Person{Age: 2},
		Person{Age: 1},
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
		Person{Age: 10},
		Person{Age: 5},
		Person{Age: 9},
		Person{Age: 20},
	}
	fmt.Printf("%v\n", isPersonAtleastTwiceOldAsOthers(personList)) // returns true
}
