package client

import (
	"crawlerByGo/src/crawler_distributed/config"
	"crawlerByGo/src/crawler_distributed/rpcSupport"
	"log"
)

func ItemSaver(host string) (chan interface{}, error) {
	client, err := rpcSupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	in := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			itemCount++
			item := <-in
			log.Printf(" item saver got itrm#%d %v", itemCount, item)
			// call rpcto save item
			result := ""
			err = client.Call(config.ItemSaverRPC, item, &result)
			if err != nil {
				log.Printf("save item#%#v error#:%v", item, err)
			}
		}
	}()
	return in, nil
}
