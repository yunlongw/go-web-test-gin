package parser

import (
	"go-web-test-gin/crawler/engine"
	"go-web-test-gin/crawler/model"
	"regexp"
)

//var name = regexp.MustCompile(`<h1 data-v-5b109fc3 class="nickName">([a-z][A-z][0-9]+[\u4e00-\u9fa5]+)</h1>`)
//var name = regexp.MustCompile(`([\u4e00-\u9fa5]+)|([0-9]+岁)|([\u4e00-\u9fa5]+)|([\u4e00-\u9fa5]+)|([0-9][a-z]+)|([0-9]+\-[0-9]+元)`)
var name = regexp.MustCompile(`<div class="m-btn purple">([^<]+)</div>`)

func ParseProfile(contents []byte, Name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = Name

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	extractString(contents, name)

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
