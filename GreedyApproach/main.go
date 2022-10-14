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
	return weight / profit
}


type MaxAndIndex struct {
	value int
	index int
}

func (g* Knapsack) maxInArray() {

}

func (g* Knapsack) BinarySearch(arr []int, low int, high int ) MaxAndIndex {
	returnValue := MaxAndIndex{value: 0 , index: -1}

	for (low <= high) {
		m := low + (high - low) / 2;
		if(low==high) {
			return MaxAndIndex{value: arr[low], index: low}
		}
		if ((high == low + 1) && ( arr[low] >= arr[high])) {
			return MaxAndIndex{value: arr[low], index: low}
		}
		if ((high == low + 1) && (arr[low] < arr[high])) {
			return MaxAndIndex{value: arr[high], index: high}
		}
		if ((arr[m] > arr[m+1]) && (arr[m] > arr[m-1])) {
			return MaxAndIndex{value: arr[m], index: m}
		}
		if ((arr[m] > arr[m+1]) && (arr[m] < arr[m-1])) {
			high = m - 1
		} else {
			low= m + 1
		}
	}
	return returnValue
}

func (g *Knapsack) MaxAndIndex(ratio []int) MaxAndIndex {
	return g.BinarySearch(ratio, 0, len(ratio)-1)
}

func (g *Knapsack) LinearMaxAndIndex(ratios []int)  MaxAndIndex{
	// currently linear search -> could chunk into right and left and compare both results
	returnValue := MaxAndIndex{value: 0, index: 0};

	for index, value := range ratios {
		if (value > returnValue.value) {
			returnValue.value = value;
			returnValue.index = index
		}
	}

	return returnValue;
}

func (g *Knapsack) Greedy() []int {
	// n marks the size or the capacity of the desired optimal output
	weightAndIndex := []int{}
	for index, value := range *g.elements {
		currentRatio := g.ProfitbyWeight(value.profit, value.weight);
		weightAndIndex[index] = currentRatio;
	}
	currentProfit := 0;
	currentWeight := 0
	for (currentProfit <= g.upperLim) {
		currentMax := g.MaxAndIndex(weightAndIndex)
		originalObject := (*g.elements)[currentMax.index]
		currentWeight+=originalObject.weight;
		currentProfit+=originalObject.profit;
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

	fmt.Println(initialKnapsack.Greedy());
}
