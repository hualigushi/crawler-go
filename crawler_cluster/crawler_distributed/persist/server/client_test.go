package main

import (
	"crawler/crawler_distributed/config"
	"crawler/crawler_distributed/rpcsupport"
	"crawler/engine"
	"crawler/model"
	"testing"
	"time"
)

func TestItemSaver (t *testing.T) {
	const host = ":1234"
	//1.start ItemSaverServer
	go serveRpc(host, "dating_profile")
	time.Sleep(time.Second)
	//2.start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	//3.Call save
	item := engine.Item{
		Url: "http://album.zhenai.com/u/1491957333",
		Type: "zhenai",
		Id: "1491957333",
		Payload: model.Profile{
			Name: "111",
			Gender: "男",
			IntroduceContent: "111111" ,
		},
	}

	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok"{
		t.Errorf("result: %s; err: %s", result, err)
	}
}