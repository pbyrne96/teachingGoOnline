package main

func getMin(a []string, b []int) (int){
	lenA:= len(a)
	lenB:= len(b)
	if (lenA < lenB) {
		return lenA
	}
	return lenB
}


func main() {

	testHashTable := Init()

	list := []string{
		"ERIC",
		"BOBBY",
		"TOMMY",
		"LOVELY",
		"L",
	}

	values := []int{
		1,
		2,
		3,
		4,
		5,
		6,
	}
	runtime:= getMin(list, values)
	iteration:=0
	for iteration <= runtime  {
		k:= list[iteration]
		v:= values[iteration]
		testHashTable.Insert(k, v)
		iteration++;
	}

}
