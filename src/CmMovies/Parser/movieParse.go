package CmParser

import (
	"crawlerByGo/src/CmMovies/Moudle"
	"crawlerByGo/src/Engine"
	"regexp"
	"time"
)

var nameRegex = regexp.MustCompile(`<dt class="name">([^<]+)<span class="bz">[^<]+</span></dt>`)

var statusRegex = regexp.MustCompile(`<dt><span>状态：</span>([^<]+)</dt>`)

var ProtagonistRegex = regexp.MustCompile(`<a target='_blank' href='(/[^m]+m=vod-search-starring-[^']+)'>([^<]+)</a>&nbsp;`)

var typeRegex = regexp.MustCompile(`<span>类型：</span><a href="/[a-z]+.html">([^<]+)</a>`)

var updateRegex = regexp.MustCompile(`<span>更新：</span>([^<]+)</dd>`)

var directorsRegex = regexp.MustCompile(`<a target='_blank' href='(/[^m]+m=vod-search-directed-[^']+)'>([^<]+)</a>&nbsp;`)

var areaRegex = regexp.MustCompile(`<span>地区：</span>([^<]+)</dd>`)

var yearRegex = regexp.MustCompile(`<span>年份：</span>([^<]+)</dd>`)

var languageRegex = regexp.MustCompile(`<dd><span>语言：</span>([^<]+)</dd>`)

var descriptionRegex = regexp.MustCompile(`<div name="ee" class="ee"><span class="js">剧情介绍：</span>([^<]+)</div>`)

var resourceRegex = regexp.MustCompile(`<a title='[^']+' href='(/play/[^']+)' rel="nofollow" target="_blank">([^<]+)</a>`)

var idRegex = regexp.MustCompile(`<a title='[^']+' href='/play/([0-9]+)[^']+' rel="nofollow" target="_blank">[^<]+</a>`)

func ParseMovie(contents []byte) Engine.ParseResult {

	movie := Moudle.Movie{}
	movie.Name = getSubMatch(nameRegex, contents)
	movie.Status = getSubMatch(statusRegex, contents)
	movie.ProtagonistList = getAllSubMatch(ProtagonistRegex, contents)
	movie.Type = getSubMatch(typeRegex, contents)
	switch movie.Type {
	case "动作片":
		movie.TableName = "action_movie"
	case "喜剧片":
		movie.TableName = "comedy"
	case "爱情片":
		movie.TableName = "romance"
	case "科幻片":
		movie.TableName = "science"
	case "恐怖片":
		movie.TableName = "horror"
	case "剧情片":
		movie.TableName = "drama"
	case "战争片":
		movie.TableName = "war"
	}
	//id, err := strconv.Atoi(getSubMatch(idRegex, contents))
	//if err != nil {
	//	log.Printf("conver string to int error #%s\n", getSubMatch(idRegex, contents))
	//	movie.Id = 99999
	//}
	movie.Id = getSubMatch(idRegex, contents)
	movie.Update = getSubMatch(updateRegex, contents)
	movie.DirectorList = getAllSubMatch(directorsRegex, contents)
	movie.Area = getSubMatch(areaRegex, contents)
	yearMatch := getSubMatch(yearRegex, contents)
	movie.Year, _ = time.Parse("2006-01-02 15:04:05", yearMatch+" 00:00:00")
	movie.Language = getSubMatch(languageRegex, contents)
	movie.Description = getSubMatch(descriptionRegex, contents)
	movie.Resource = getAllSubMatch(resourceRegex, contents)

	result := Engine.ParseResult{
		Items: []interface{}{movie},
	}
	return result
}
func getAllSubMatch(compile *regexp.Regexp, contents []byte) []Moudle.Info {
	match := compile.FindAllSubmatch(contents, -1)
	var infos []Moudle.Info
	for _, director := range match {
		infos = append(infos, Moudle.Info{
			Url:  "https://www.cmcm5.com" + string(director[1]),
			Name: string(director[2]),
		})
	}
	return infos
}
func getSubMatch(compile *regexp.Regexp, contents []byte) string {

	match := compile.FindSubmatch(contents)
	if len(match) > 1 {
		return string(match[1])
	} else {
		return ""
	}
}
