package director

import "designPatternsByGolang/Builder/builder"

type Director struct {
	builder.Builder
}

func (D *Director) Construct() {
	D.MakeTitle("快餐店开门啦")
	D.MakeString("上午菜单")
	D.MakeItems([]string{"包子!", "豆浆!"})
	D.MakeString("下午茶")
	D.MakeItems([]string{"奶茶!", "巧克力!", "蛋糕!"})
	D.Close()
}
