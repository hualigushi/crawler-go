package main

import (
	"crawler/crawler_distributed/config"
	itemsaver "crawler/crawler_distributed/persist/client"
	"crawler/crawler_distributed/rpcsupport"
	worker "crawler/crawler_distributed/worker/client"
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver")

	workerHosts = flag.String("worker_hosts", "","worker hosts (comm separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}
    pool := createClientPoor(strings.Split(*workerHosts, ","))
	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan: itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParserCityList, config.ParseCityList),
	})

}

func createClientPoor(hosts []string) chan *rpc.Client{
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("error connneting to %s: %v", h ,err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
