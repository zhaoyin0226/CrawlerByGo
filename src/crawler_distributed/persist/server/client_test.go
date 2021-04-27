package main

import (
	"crawlerByGo/src/CmMovies/Moudle"
	"crawlerByGo/src/crawler_distributed/rpcSupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = "1234"
	// startItemSaverServer
	go serverRpc(host, "test_item")
	time.Sleep(time.Second)
	// startItemSaverClient
	client, err := rpcSupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	// call save
	result := ""
	movie := Moudle.Movie{
		Id:        "99999999",
		TableName: "test_item",
		Name:      " test server",
		Status:    "update",
		Type:      "free",
	}
	err = client.Call("ItemSaverService.Save", movie, &result)
	if err != nil || result != "ok" {
		t.Errorf("save error #: %v, result is #%s", err, result)
	}
}
