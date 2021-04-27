package CmParser

import (
	"crawlerByGo/src/Engine"
	"crawlerByGo/src/crawler_distributed/config"
	"regexp"
)

var (
	movieListRegex = regexp.MustCompile(`<a class="link-hover" href="(/[a-z]+/[a-z0-9]+.html)" title="([^"]+)" [^>]*>`)
	pageRegex      = regexp.MustCompile(`<a target="_self" href="(/[^m]+m=vod-list-id-[0-9]+-pg-[0-9]+-order--by--class-0-year-[0-9]*-letter--area-[^-]*-lang-.html)" class="pagelink_a">下一页</a>`)
)

func ParseMovieList(contents []byte) Engine.ParseResult {

	result := Engine.ParseResult{}

	subMatches := movieListRegex.FindAllSubmatch(contents, -1)
	for _, subMatch := range subMatches {
		//result.Items = append(result.Items, "movieList: "+string(subMatch[2]))
		result.Requests = append(result.Requests, Engine.Request{
			Url:    "https://www.cmcm5.com" + string(subMatch[1]),
			Parser: Engine.NewFuncParser(ParseMovie, config.ParseMovie),
			//ParserFunc: ParseMovie,
		})
	}

	pageSubMatch := pageRegex.FindSubmatch(contents)
	if len(pageSubMatch) > 0 {
		//result.Items = append(result.Items, "moviePage: https://www.cmcm5.com"+string(pageSubMatch[1]))
		result.Requests = append(result.Requests, Engine.Request{
			Url:    "https://www.cmcm5.com" + string(pageSubMatch[1]),
			Parser: Engine.NewFuncParser(ParseMovieList, config.ParseMovieList),
		})
	}

	return result
}
