package Worker

import (
	"crawlerByGo/src/Engine"
	"crawlerByGo/src/Fetcher"
	"log"
)

func Worker(request Engine.Request) (Engine.ParseResult, error) {
	body, err := Fetcher.Fetch(request.Url)
	log.Printf("Fetching %s", request.Url)
	if err != nil {
		return Engine.ParseResult{}, err
	}
	parseResult := request.Parser.Parse(body)
	return parseResult, nil
}
