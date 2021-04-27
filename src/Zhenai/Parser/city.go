package Parser

import (
	"crawlerByGo/src/Engine"
	"regexp"
)
const cityRegex = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
func ParseCity(contents []byte) Engine.ParseResult{
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
