package decorator

import "designPatternsByGolang/Decorator/display"

type FullBorder struct {
	d display.Display
	*display.DisplaySuper
}

func NewFullBorder(d display.Display) *FullBorder {
	border := &FullBorder{}
	super := display.NewDisplaySuper(border)
	border.DisplaySuper = super
	border.d = d
	return border
}

func (F *FullBorder) GetColumns() int {
	return 1 + F.d.GetColumns() + 1
}

func (F *FullBorder) GetRows() int {
	return 1 + F.d.GetRows() + 1
}

func (F *FullBorder) GetRowText(row int) string {
	if row == 0 { // 上边框
		return "+" + F.makeLine("-", F.d.GetColumns()) + "+"
	} else if row == F.d.GetRows()+1 { // 下边框
		return "+" + F.makeLine("-", F.d.GetColumns()) + "+"
	} else { // 其他边框
		return "|" + F.d.GetRowText(row-1) + "|"
	}
}

func (F *FullBorder) makeLine(ch string, count int) string {
	res := ""
	for i := 0; i < count; i++ {
		res += ch
	}
	return res
}
