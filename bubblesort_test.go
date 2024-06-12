package pyroscopeaction

import (
	"math/rand"
	"testing"
)

func BenchmarkBubbleSort1000(b *testing.B) {
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i
	}

	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	arr2 := make([]int, len(arr))

	for n := 0; n < b.N; n++ {
		copy(arr2, arr)
		BubbleSort(arr2)
	}
}
