package whatever

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQsort(t *testing.T) {
	randList := make([]int, 100)
	for i := 0; i < 100; i++ {
		randList[i] = rand.Intn(100)
	}
	fmt.Println(randList)
	QuickSort(randList)
	fmt.Println(randList)
}
