package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

// 单机版相亲网站爬虫,获取并打印所有城市第一页用户的详细信息
func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}

