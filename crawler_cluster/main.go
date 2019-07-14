package main

import (
	"crawler/crawler_distributed/config"
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan: itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParserCityList, config.ParseCityList),
	})

}


