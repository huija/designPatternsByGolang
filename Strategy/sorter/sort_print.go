package sorter

import "fmt"

type SortAndPrint struct {
	array  []interface{}
	sorter Sorter
}

func NewSortAndPrint(array []interface{}, sorter Sorter) *SortAndPrint {
	return &SortAndPrint{array: array, sorter: sorter}
}

func (S *SortAndPrint) Execute() {
	S.Print()
	S.sorter.Sort(S.array)
	S.Print()
}

func (S *SortAndPrint) Print() {
	for i := 0; i < len(S.array); i++ {
		fmt.Printf("%+v, ", S.array[i])
	}
	fmt.Println()
}
