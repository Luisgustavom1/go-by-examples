package sort

import (
	"math/rand"
	"time"
)

func GenerateRandomSliceUnoptimized(size int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	Dirty()
	slice := []int{}
	for i := 0; i < size; i++ {
		slice = append(slice, rand.Intn(1000))
	}
	return slice
}

func GenerateRandomSliceOptimized(size int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	Dirty()
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1000)
	}
	return slice
}

func Dirty() []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := []int{}
	for i := 0; i < 10000; i++ {
		slice = append(slice, rand.Intn(1000))
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
