/**
 * @author wellsjiang
 */

package wellgo

type RPC interface {
	RPCHandler(Request) (Request, error)

	EncodeResponse(*WContext, Result) ([]byte, error)

	EncodeErrResponse(*WContext, Result) ([]byte, error)
}
