package client

import (
	"crawler/crawler_distributed/config"
	"crawler/crawler_distributed/worker"
	"crawler/engine"
	"net/rpc"
)

func CreateProcessor (clientsChan chan *rpc.Client) engine.Processor {
	return func (req engine.Request) (engine.ParserResult, error) {
		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult
		c := <- clientsChan

		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)

		if err != nil {
			return engine.ParserResult{}, nil
		}
		return worker.DeserializeResult(sResult), nil
	}
}