package support

import "designPatternsByGolang/ChainOfResponsibility/trouble"

type LimitSupport struct {
	*Support
	limit int
}

func NewLimitSupport(name string, limit int) *LimitSupport {
	limitSupport := &LimitSupport{limit: limit}
	support := NewSupport(name)
	support.SP = limitSupport
	limitSupport.Support = support
	return limitSupport
}

//Resolve 只处理奇数
func (L *LimitSupport) Resolve(trouble *trouble.Trouble) bool {
	return trouble.GetNumber() < L.limit
}
