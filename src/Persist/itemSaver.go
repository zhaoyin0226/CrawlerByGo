package Persist

import (
	"context"
	"crawlerByGo/src/CmMovies/Moudle"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

func ItemSaver(indexName string) chan interface{} {
	in := make(chan interface{})
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
			//"http://localhost:9201",
		},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * 2,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MaxVersion:         tls.VersionTLS11,
				InsecureSkipVerify: true,
			},
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	//es.Index("cmmovie", strings.NewReader(string(info)))
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	go func() {
		itemCount := 0
		for {
			itemCount++
			item := <-in
			log.Printf(" item saver got itrm#%d %v", itemCount, item)
			Save(item, es, indexName)
		}
	}()
	return in
}
func Save(item interface{}, es *elasticsearch.Client, indexName string) error {
	//client, err := elasticsearch.NewDefaultClient()
	info, err := json.Marshal(item)
	if err != nil {
		log.Fatalf("Error parse the item: %s", err)
		//return errors.New(fmt.Sprintf("Error parse the item: %s", err))
	}

	res, err := es.Info()
	log.Println(res)

	req := esapi.IndexRequest{
		Index: indexName,
		//DocumentType: item.(Moudle.Movie).TableName,
		DocumentType: "movie",
		DocumentID:   item.(Moudle.Movie).Id,
		Body:         strings.NewReader(fmt.Sprintf("%s", info)),
		Refresh:      "true",
	}
	response, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return err
	}
	defer response.Body.Close()
	if response.IsError() {
		log.Printf("[%s] Error indexing Name ID=%s, tableName is %s", response.Status(), item.(Moudle.Movie).Name, item.(Moudle.Movie).TableName)
		return errors.New(fmt.Sprintf("[%s] Error indexing Name ID=%s, tableName is %s", response.Status(), item.(Moudle.Movie).Name, item.(Moudle.Movie).TableName))
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(response.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", response.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
	return nil
}
