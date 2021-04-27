package funcConfig

import (
	CmParser "crawlerByGo/src/CmMovies/Parser"
	"crawlerByGo/src/Engine"
	"crawlerByGo/src/crawler_distributed/config"
)

var FUNC_MAP = map[string]Engine.ParserFunc{
	config.ParseCategoryList: CmParser.ParseCategoryList,
	config.ParseMovieList:    CmParser.ParseMovieList,
	config.ParseMovie:        CmParser.ParseMovie,
}
