package main

import "fmt"

const ArraySize int = 64;

type HashTable struct {
	Array [] *bucket
}


func filterNil(data []*bucket, predicate func(value any) bool) ([]*bucket){
	result := []*bucket{}
   for _, value := range data {
	   if (predicate(value)) {
		   result = append(result, value)
	   }
   }

   return result
}

func notNil(value any) bool{
   return value != nil
}

func checkResize(data []*bucket) bool{
   return len(filterNil(data, notNil)) == len(data)
}

func (h *HashTable) concatArrays(existingArray []*bucket) []*bucket{
	newSize := (ArraySize+ArraySize)
	newArray := make([]*bucket, newSize, newSize)

	runtime := (len(existingArray) -1)
	i := 0
	for i <= runtime  {
		newArray[i] = existingArray[i]
		i++
	}

	h.Array = newArray
	return newArray
}

func Hash(key string) uint64{
	Hash := uint64(5381)
	for _ , b :=  range key{
		Hash += uint64(b) + Hash + Hash<<5
	}
	return Hash % uint64(ArraySize)
}


func (h *HashTable) getKeys() []string {
	nonNilArray := filterNil(h.Array, notNil);
	emptyArr := make([]string, 0, len(nonNilArray) * 2);
	for _, value := range nonNilArray {
		currentNode := value.head;
		currentDepthArr := make([]string, 0 , 64);
		lastInsert := 0;
		for currentNode != nil {
			currentDepthArr[lastInsert] = currentNode.key;
			currentNode = currentNode.next;
		}
		emptyArr = append(emptyArr, currentDepthArr...);
	}
	return emptyArr;
}

func (h *HashTable) getValue() []any{
	nonNilArray := filterNil(h.Array, notNil);
	emptyArr := make([]any, 0, len(nonNilArray) * 2);
	for _, value := range nonNilArray {
		currentNode := value.head;
		currentDepthArr := make([]any, 0 , 64);
		lastInsert := 0;
		for currentNode != nil {
			currentDepthArr[lastInsert] = currentNode.value;
			currentNode = currentNode.next;
		}
		emptyArr = append(emptyArr, currentDepthArr...);
	}
	return emptyArr;
}

func (h *HashTable) getItems() [][]any{
	nonNilArray := filterNil(h.Array, notNil);
	emptyArr := make([][]any, 0, len(nonNilArray) * 2);
	for _, value := range nonNilArray {
		currentNode := value.head;
		currentDepthArr := make([][]any, 0 , 64);
		lastInsert := 0;
		for currentNode != nil {
			arrAtPoint := make([] any, 2, 2);
			arrAtPoint[0] = currentNode.key
			arrAtPoint[1] = currentNode.value
			currentDepthArr[lastInsert] = arrAtPoint;
			currentNode = currentNode.next;
		}
		emptyArr = append(emptyArr, currentDepthArr...);
	}
	return emptyArr;
}

func (h *HashTable) Insert(key string, value int)  bool{
	// need to check size and weather to increase the size should also be done here
	resize := checkResize(h.Array[:])
	if (resize) {
		// do some resizing
		newArray := h.concatArrays(h.Array[:])
		fmt.Printf("newArray = %T\n",newArray)

	}
	index := Hash(key)
	return h.Array[index].insert(key, value)
}

func (h *HashTable) Search(key string) bool {
	index := Hash(key)
	return h.Array[index].search(key)
}


func (h *HashTable)	Delete(key string) bool {
	index := Hash(key)
	return h.Array[index].delete(key)
}


func Init() *HashTable {
	result := &HashTable{}
	i:=0
	for i <= ArraySize {
		result.Array[i] = &bucket{}
		i++
	}
	return result
}
