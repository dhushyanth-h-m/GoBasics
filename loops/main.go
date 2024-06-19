package main

import (
	"fmt"
	"slices"
)

func main() {
	animals := [2]string{} //array of strings

	animals[0] = "Dog"
	animals[1] = "Cat"

	fmt.Println(animals)

	flexAnimals := []string{ //slice of strings
		"Dog",
		"Cat",
		"Fish",
	}

	flexAnimals = append(flexAnimals, "Bird")
	fmt.Println(flexAnimals)

	flexAnimals = slices.Delete(flexAnimals, 0, 1)

	fmt.Println(flexAnimals)

	for i := 0; i < len(flexAnimals); i++ {
		fmt.Printf("Animal at %d is %s\n", i, flexAnimals[i])
	}

	for index, value := range flexAnimals {
		fmt.Printf("This is my index %d and this is my animal %s\n", index, value)
	}
}
