package main

import (
	"crawler/crawler_distributed/config"
	"crawler/crawler_distributed/persist"
	"crawler/crawler_distributed/rpcsupport"
	"flag"
	"fmt"
	"github.com/olivere/elastic"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Printf("must specify a port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.Elasticindex))
}

func serveRpc(host string, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://47.97.163.47:9200"),
	)
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
