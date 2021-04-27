package Parser

import (
	"crawlerByGo/src/Engine"
	"regexp"
)
const cityListRegex = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`
func ParseCityList(contents []byte) Engine.ParseResult{

	compile := regexp.MustCompile(cityListRegex)
	subMatches := compile.FindAllSubmatch(contents, -1)
	result := Engine.ParseResult{}
	for _, subMatche := range subMatches {
		result.Items = append(result.Items,"City: "+string(subMatche[2]))
		result.Requests = append(result.Requests,Engine.Request{
			Url: string(subMatche[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
