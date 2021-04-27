package main

import (
	"crawlerByGo/src/crawler_distributed/config"
	"crawlerByGo/src/crawler_distributed/rpcSupport"
	workerService "crawlerByGo/src/crawler_distributed/worker"
	"log"
)

//var port = flag.Int("name", 0, "CrawlServerPort")
//TODO 命令行配置port go run **.go -port 9000
func main() {
	//flag.Parse()
	//if *port == 0 {
	//	log.Printf("must specify a port")
	//	return
	//}
	log.Fatal(rpcSupport.ServeRpc(config.CrawlWorkerPort0, workerService.CrawlService{}))
	log.Fatal(rpcSupport.ServeRpc("9001", workerService.CrawlService{}))
}
