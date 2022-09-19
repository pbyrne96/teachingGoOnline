package main


func main() {

	testHashTable := Init()

	list := []string{
		"ERIC",
		"BOBBY",
		"TOMMY",
		"LOVELY",
		"L",
	}


	for _, v := range list {
		testHashTable.Insert(v)
	}

}
