package support

import "designPatternsByGolang/ChainOfResponsibility/trouble"

type OddSupport struct {
	*Support
}

func NewOddSupport(name string) *OddSupport {
	oddSupport := &OddSupport{}
	support := NewSupport(name)
	support.SP = oddSupport
	oddSupport.Support = support
	return oddSupport
}

//Resolve 只处理奇数
func (O *OddSupport) Resolve(trouble *trouble.Trouble) bool {
	return trouble.GetNumber()%2 == 1
}
