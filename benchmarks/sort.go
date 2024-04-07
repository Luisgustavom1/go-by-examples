package sort

import (
	"math/rand"
	"time"
)

func GenerateRandomSlice(size int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1000)
	}
	return slice
}

func Sort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}