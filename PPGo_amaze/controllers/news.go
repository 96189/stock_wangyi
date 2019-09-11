package controllers


type NewsController struct {
	BaseController
}

func (self *NewsController) News() {
	self.Ctx.WriteString("敬请期待")
}