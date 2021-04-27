package workerService

import (
	"crawlerByGo/src/Worker"
)

type CrawlService struct {
}

func (service CrawlService) Process(request Request, result *ParseResult) error {
	deserializeRequest, err := DeserializeRequest(request)
	if err != nil {
		return err
	}
	parseResult, err := Worker.Worker(deserializeRequest)
	*result = SerializeParseResult(parseResult)
	if err != nil {
		return err
	}
	return nil
}
