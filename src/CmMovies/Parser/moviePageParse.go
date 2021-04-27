package CmParser

import (
	"crawlerByGo/src/Engine"
)

// byCategory
const byCategoryPageRegex = `<a target="_self" href="(/[^m]+m=vod-list-id-[0-9]+-pg-[0-9]+-order--by--class-0-year-[0-9]*-letter--area-[^-]*-lang-.html)" class="pagelink_a">下一页</a>`

func ParsePage(contents []byte) Engine.ParseResult {

	result := Engine.ParseResult{}

	//compile := regexp.MustCompile(byCategoryPageRegex)
	//subMatch := compile.FindSubmatch(contents)
	//if len(subMatch) > 0 {
	//	//result.Items = append(result.Items, "moviePage: https://www.cmcm5.com"+string(subMatch[1]))
	//	result.Requests = append(result.Requests, Engine.Request{
	//		Url:        "https://www.cmcm5.com" + string(subMatch[1]),
	//		ParserFunc: ParsePage,
	//	})
	//	result.Requests = append(result.Requests, Engine.Request{
	//		Url:        "https://www.cmcm5.com" + string(subMatch[1]),
	//		ParserFunc: ParseMovieList,
	//	})
	//} else {
	//	ParseMovieList(contents)
	//}

	return result
}
