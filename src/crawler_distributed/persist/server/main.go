package main

import (
	"crawlerByGo/src/crawler_distributed/config"
	"crawlerByGo/src/crawler_distributed/persist"
	"crawlerByGo/src/crawler_distributed/rpcSupport"
	"github.com/elastic/go-elasticsearch"
	"log"
)

func main() {
	log.Fatal(serverRpc(config.ItemSaverPort, config.ElasticIndex))
}
func serverRpc(host, indexName string) error {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return err
	}
	return rpcSupport.ServeRpc(host, &persist.ItemSaverService{
		Client:    client,
		IndexName: indexName,
	})
}
