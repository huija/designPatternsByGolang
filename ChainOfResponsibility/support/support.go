package support

import (
	"designPatternsByGolang/ChainOfResponsibility/trouble"
	"fmt"
	"github.com/gokangaroo/common/utils"
)

type SP interface {
	Resolve(trouble *trouble.Trouble) bool
}

type Support struct {
	SP
	name string
	next *Support
}

func NewSupport(name string) *Support {
	return &Support{name: name}
}

func (S *Support) SetNext(SN *Support) {
	S.next = SN
}

func (S *Support) String() string {
	return "[" + S.name + "]"
}

//SupportTrouble 链式处理
func (S *Support) SupportTrouble(trouble *trouble.Trouble) {
	for obj := S; true; obj = S.next {
		if obj.Resolve(trouble) {
			obj.Done(trouble)
			break
		} else if obj.next == nil {
			obj.Fail(trouble)
			break
		}
	}
}

func (S *Support) Done(trouble *trouble.Trouble) {
	fmt.Println(utils.ToString(trouble) + " is resolved by " + utils.ToString(S) + ".")
}

func (S *Support) Fail(trouble *trouble.Trouble) {
	fmt.Println(utils.ToString(trouble) + " cannot be resolved.")
}
