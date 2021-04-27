package CmParser

import (
	"crawlerByGo/src/Engine"
	"crawlerByGo/src/crawler_distributed/config"
	"regexp"
)

// byCategory
const byCategoryRegex = `<a href="(/[^m]+m=vod-list-id-[0-9]+-pg-1-order--by--class--year--letter--area--lang-.html)" >([^<]+)</a>`

// byYear
const byYearRegex = `href="(/[^m]+m=vod-list-id-1-pg-1-order--by--class--year-[0-9]+-letter--area--lang-.html)">([^<]+)`

// byArea
const byAreaRegex = `href="(/[^m]+m=vod-list-id-1-pg-1-order--by--class--year--letter--area-[^-]+-lang-.html)">([^<]+)`

// ParseCategoryList <a target="_self" href=("/[^m]+m=vod-list-id-[0-9]+-pg-[0-9]+-order--by--class-0-year-0-letter--area--lang-.html" class="pagelink_a">下一页</a>
func ParseCategoryList(contents []byte) Engine.ParseResult {

	regexList := []string{byCategoryRegex, byYearRegex, byAreaRegex}
	result := Engine.ParseResult{}
	for _, regex := range regexList {
		compile := regexp.MustCompile(regex)
		subMatches := compile.FindAllSubmatch(contents, -1)
		for _, subMatch := range subMatches {
			//result.Items = append(result.Items, "movieCategory: "+string(subMatch[2]))
			result.Requests = append(result.Requests, Engine.Request{
				Url:    "https://www.cmcm5.com" + string(subMatch[1]),
				Parser: Engine.NewFuncParser(ParseMovieList, config.ParseMovieList),
			})
		}
	}

	return result
}
