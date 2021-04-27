package Engine

import (
	"crawlerByGo/src/Fetcher"
	"log"
)

type SimpleEngine struct{}

func (se SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}
	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]
		parseResult, err := worker(request)
		if err != nil {
			log.Printf("Fetcher error :"+
				"fetchint url %s : %v", request.Url, err)
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf(" got item : %s ;", item)
		}
	}
}
func worker(request Request) (ParseResult, error) {
	body, err := Fetcher.Fetch(request.Url)
	log.Printf("Fetching %s", request.Url)
	if err != nil {
		return ParseResult{}, err
	}

	parseResult := request.Parser.Parse(body)
	return parseResult, nil
}
