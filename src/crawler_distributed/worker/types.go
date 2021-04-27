package workerService

import (
	"crawlerByGo/src/Engine"
	"crawlerByGo/src/crawler_distributed/config/funcConfig"
	"errors"
	"fmt"
	"log"
)

type SerializeParser struct {
	FunctionName string
	Args         interface{}
}

type Request struct {
	Url    string
	Parser SerializeParser
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func SerializeRequest(r Engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializeParser{
			FunctionName: name,
			Args:         args,
		},
	}
}

func SerializeParseResult(pr Engine.ParseResult) ParseResult {
	var r []Request
	for _, request := range pr.Requests {
		r = append(r, SerializeRequest(request))
	}
	return ParseResult{
		Items:    pr.Items,
		Requests: r,
	}
}

func DeserializeRequest(r Request) (Engine.Request, error) {
	parserFunc, ok := funcConfig.FUNC_MAP[r.Parser.FunctionName]
	if ok {
		return Engine.Request{
			Url:    r.Url,
			Parser: Engine.NewFuncParser(parserFunc, r.Parser.FunctionName),
		}, nil
	}
	return Engine.Request{}, errors.New(fmt.Sprintf("undefind parserFunc#:%s", r.Parser.FunctionName))
}

func DeserializeParseResult(ps ParseResult) Engine.ParseResult {
	var r []Engine.Request
	for _, request := range ps.Requests {
		deserializeRequest, err := DeserializeRequest(request)
		if err != nil {
			log.Printf("DeserializeParseResult error:#%v", err)
			continue
		}
		r = append(r, deserializeRequest)
	}

	return Engine.ParseResult{
		Items:    ps.Items,
		Requests: r,
	}
}
