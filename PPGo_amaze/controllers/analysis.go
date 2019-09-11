package controllers


type AnalysisController struct {
	BaseController
}

// TODO xsrf
func (self *AnalysisController) Analyze() {
	self.Ctx.WriteString("敬请期待")
}

