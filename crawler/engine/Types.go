package engine

// url 和 URL 解析器
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

//
type ParseResult struct {
	Requests []Request     // 城市 url 和对应的解析方法
	Items    []interface{} // 城市名称列表
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
