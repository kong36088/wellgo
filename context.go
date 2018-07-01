package wellgo


type Context struct{
	cfg Config

	router Router

	req Request

	resp Response

	proto string
}

type Request interface{
}

type Response interface{

}
