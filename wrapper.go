/**
  * @author wellsjiang
  * @date 2018/8/3
  */

package wellgo

type RPCHandlerFunc func(req Request) (Request, error)
