package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

// 单机版相亲网站爬虫,获取并打印所有城市第一页用户的详细信息
func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		// Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	//e.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParserCityList,
	//})
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}


