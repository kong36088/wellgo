/**
  * @author wellsjiang
  * @date 2018/8/2
  */

package wellgo

type ControllerInterface interface {
	Init(*WContext)
	Run() *Result
}

type Controller struct {
	Ctx *WContext

	Args map[string]interface{}
}

func (c *Controller) Init(ctx *WContext) {
	c.Ctx = ctx
	c.Args = ctx.Req.GetArgs()
}

func (c *Controller) Run() *Result {
	return NewResult(0, OK.Error())
}
