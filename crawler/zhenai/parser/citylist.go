package parser

import (
	"go-web-test-gin/crawler/engine"
	"regexp"
)

var CityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

// 获取城市数据
func ParseCityList(contents []byte) engine.ParseResult {
	matches := CityListRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		var i int = 1
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})

		i--
		if i == 0 {
			break
		}
	}
	return result
}
