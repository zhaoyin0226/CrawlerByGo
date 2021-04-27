package main

import (
	CmParser "crawlerByGo/src/CmMovies/Parser"
	"crawlerByGo/src/crawler_distributed/rpcSupport"
	"log"
	"net/rpc"

	//CmParser "crawlerByGo/src/CmMovies/Parser"
	"crawlerByGo/src/Engine"
	"crawlerByGo/src/Scheduler"
	"crawlerByGo/src/crawler_distributed/config"
	itemSaverClient "crawlerByGo/src/crawler_distributed/persist/client"
	crawlWorkerClient "crawlerByGo/src/crawler_distributed/worker/client"
)

func main() {

	itemChan, err := itemSaverClient.ItemSaver(config.ItemSaverPort)
	if err != nil {
		panic(err)
	}
	clientPool := createClientPool([]string{"9000", "9001"})
	processor := crawlWorkerClient.CreateProcessor(clientPool)

	engine := Engine.ConcurrentEngine{
		Scheduler:   &Scheduler.SimpleScheduler{},
		WorkerCount: 100,
		//ItemChan:    Persist.ItemSaver("movie_db"),
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	engine.Run(Engine.Request{
		Url:    "https://www.cmcm5.com/dianying.html",
		Parser: Engine.NewFuncParser(CmParser.ParseCategoryList, config.ParseCategoryList),
	})

}

func createClientPool(hosts []string) chan *rpc.Client {

	var clients []*rpc.Client
	for _, host := range hosts {
		client, err := rpcSupport.NewClient(host)
		if err != nil {
			log.Printf("Connect #%s error: %v", host, err)
		} else {
			log.Printf("Connect to client #%s", host)
			clients = append(clients, client)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
