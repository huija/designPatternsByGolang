package main

import "designPatternsByGolang/Strategy/sorter"

func main() {
	array1 := []interface{}{3, 2, 1, 4, 5}
	sap1 := sorter.NewSortAndPrint(array1, &sorter.SelectSorter{})
	sap1.Execute()
	array2 := []interface{}{3, 2, 1, 4, 5}
	sap2 := sorter.NewSortAndPrint(array2, &sorter.BubbleSorter{})
	sap2.Execute()
}
