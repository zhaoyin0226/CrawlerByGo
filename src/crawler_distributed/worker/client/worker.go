package client

import (
	"crawlerByGo/src/Engine"
	"crawlerByGo/src/crawler_distributed/config"
	"crawlerByGo/src/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clients chan *rpc.Client) Engine.Processor {

	return func(request Engine.Request) (Engine.ParseResult, error) {
		serviceRequest := workerService.SerializeRequest(request)
		var serviceResult workerService.ParseResult

		c := <-clients
		//client, err := rpcSupport.NewClient("9000")
		//if err != nil {
		//	panic(err)
		//}
		//log.Printf("Calling #%serviceRequest by #%v", serviceRequest, c)
		clientError := c.Call(config.CrawlWorkerPRC, serviceRequest, &serviceResult)
		if clientError != nil {
			return Engine.ParseResult{}, clientError
		}
		engineParseResult := workerService.DeserializeParseResult(serviceResult)
		return engineParseResult, nil
	}
}
