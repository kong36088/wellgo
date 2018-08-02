/**
  * @author wellsjiang
  * @date 2018/8/2
  */

package wellgo

type ControllerInterface interface {
	Init(*WContext)
	Run()
}

type Controller struct {
	Ctx *WContext

	Args map[string]interface{}
}

func (c *Controller) Init(ctx *WContext) {
	c.Ctx = ctx
	c.Args = ctx.req.GetArgs()
}

func (c *Controller) Run() {}
