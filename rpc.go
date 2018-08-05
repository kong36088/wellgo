/**
 * @author wellsjiang
 */

package wellgo

type RPC interface {
	RPCHandler(Request) (Request, error)

	EncodeResponse(Request, Response) ([]byte, error)

	EncodeErrResponse(Request, Response, error) ([]byte, error)
}
