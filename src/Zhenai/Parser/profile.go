package Parser

import (
	"crawlerByGo/src/Engine"
	"regexp"
)
const weddingRegex = `<div class="m-btn purple"></div>`


func ParseProfile(contents []byte) Engine.ParseResult{
	compile := regexp.MustCompile(cityRegex)
	subMatches := compile.FindAllSubmatch(contents, -1)
	result := Engine.ParseResult{}
	for _, subMatche := range subMatches {
		result.Items = append(result.Items,"User: "+string(subMatche[2]))
		result.Requests = append(result.Requests,Engine.Request{
			Url: string(subMatche[1]),
			ParserFunc: Engine.NilParser,
		})
	}
	return result
}
