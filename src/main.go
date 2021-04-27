package main

import (
	CmParser "crawlerByGo/src/CmMovies/Parser"
	"crawlerByGo/src/Engine"
	"crawlerByGo/src/Persist"
	"crawlerByGo/src/Scheduler"
	"crawlerByGo/src/Worker"
	"crawlerByGo/src/crawler_distributed/config"
)

func main() {
	//Engine.SimpleEngine{}.Run(Engine.Request{
	//	Url:        "https://www.cmcm5.com/dianying.html",
	//	ParserFunc: CmParser.ParseCategoryList,
	//})

	engine := Engine.ConcurrentEngine{
		Scheduler:        &Scheduler.SimpleScheduler{},
		WorkerCount:      100,
		ItemChan:         Persist.ItemSaver("movie_db"),
		RequestProcessor: Worker.Worker,
	}

	engine.Run(Engine.Request{
		Url:    "https://www.cmcm5.com/dianying.html",
		Parser: Engine.NewFuncParser(CmParser.ParseCategoryList, config.ParseCategoryList),
	})
	//movie := Moudle.Info{
	//	Url:  "www.bilibili.com",
	//	Name: "bilibili",
	//}
	//Persist.Save(movie)
	//Persist.Save([]string{"1"})

	//contents, _ := Fetcher.Fetch("https://www.cmcm5.com/index.php?m=vod-list-id-5-pg-1-order--by--class--year--letter--area--lang-.html")
	//CmParser.ParsePage(contents)

}
