package client

import (
	"crawler/crawler_distributed/config"
	"crawler/crawler_distributed/rpcsupport"
	"crawler/engine"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)

	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func () {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			// Call RPC to save item
			result := ""
			err := client.Call(config.ItemSaverRpc,
				item, &result,
				)
			if err != nil {
				log.Printf("Item saver error %v : %v\n",item,err)
			}
		}
	}()
	return out, nil
}
