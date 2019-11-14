package sorter

import (
	"github.com/gokangaroo/common/utils"
)

type Sorter interface {
	Sort(array []interface{})
}

// 选择排序
type SelectSorter struct {
}

func (S *SelectSorter) Sort(array []interface{}) {
	for i := 0; i < len(array)-1; i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if utils.Compare(array[i], array[j]) > 0 {
				min = j
			}
			tmp := array[min]
			array[min] = array[i]
			array[i] = tmp
		}
	}
}

// 冒泡排序
type BubbleSorter struct {
}

func (B *BubbleSorter) Sort(array []interface{}) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if utils.Compare(array[j], array[j+1]) > 0 {
				tmp := array[j]
				array[j] = array[j+1]
				array[j+1] = tmp
			}
		}
	}
}
