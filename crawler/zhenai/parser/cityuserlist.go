package parser

import (
	"go-web-test-gin/crawler/engine"
	"regexp"
)

var cityUserListRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

func ParseCity(content []byte) engine.ParseResult {
	matches := cityUserListRe.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		var i int = 1
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
		i--
		if i == 0 {
			break
		}
	}
	return result
}
