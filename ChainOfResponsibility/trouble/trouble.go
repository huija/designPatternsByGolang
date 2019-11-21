package trouble

import "fmt"

type Trouble struct {
	number int
}

func NewTrouble(number int) *Trouble {
	return &Trouble{number: number}
}

func (T *Trouble) GetNumber() int {
	return T.number
}

func (T *Trouble) String() string {
	return "[Trouble " + fmt.Sprintf("%d", T.number) + "]"
}
