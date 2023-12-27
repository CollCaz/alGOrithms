package Sorts

// import "fmt"
import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Ordered](a []T) []T {
	for i := range a {
		if i == len(a)-1 {
			break
		}

		if a[i] > a[i+1] {
			tmp := a[i]
			a[i] = a[i+1]
			a[i+1] = tmp
			BubbleSort(a)
		}
	}

	return a
}
