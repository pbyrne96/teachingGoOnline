package main

import "fmt"


type object struct {
	profit int;
	weight int;
}

type Knapsack struct {
	elements *[]object
	upperLim int
}

func (g *Knapsack) init(elements []object ) *Knapsack {
	return &Knapsack{elements: &elements};
}


func (g *Knapsack) Acceptable(check int) bool {
	return true
}

func (g *Knapsack) ProfitbyWeight(profit int, weight int) int {
	return profit ^ weight
}

func (g *Knapsack) Greedy(n int) []int {
	// n marks the size or the capacity of the desired optimal output

	cnt := 0;
	for _, index := range *g.elements {
		fmt.Println(index)
		cnt++;
	}

	return []int{0};
}

func main () {
	initialKnapsack := new(Knapsack)
	allObject := make([]object, 0);
	initialData := map[int]int{
		10:2,
		5:3,
		15:5,
		7:7,
		6:1,
		18:4,
		3:2,
	}

	for key, value := range initialData {
		allObject = append(allObject, object{profit: key, weight: value } )
	}

	initialKnapsack.elements = &allObject;
	initialKnapsack.upperLim = 15

	fmt.Println(initialKnapsack);
}
