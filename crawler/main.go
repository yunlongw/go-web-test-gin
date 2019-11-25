package main

import (
	"go-web-test-gin/crawler/engine"
	"go-web-test-gin/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList, // 城市解析器
	})
}
