package request_parser

type BaseParser interface{
	parse([]byte)

	getRequestUrl()
}

