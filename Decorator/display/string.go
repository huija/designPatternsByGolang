package display

type StringDisplay struct {
	*DisplaySuper
	str string
}

func NewStringDisplay(str string) *StringDisplay {
	display := &StringDisplay{str: str}
	super := NewDisplaySuper(display)
	display.DisplaySuper = super
	return display
}

func (S *StringDisplay) GetColumns() int {
	return len(S.str)
}

func (S *StringDisplay) GetRows() int {
	return 1
}

func (S *StringDisplay) GetRowText(row int) string {
	if row == 0 {
		return S.str
	}
	return ""
}
