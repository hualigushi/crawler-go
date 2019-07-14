package persist

import (
	"context"
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"github.com/olivere/elastic"
	"testing"
)

func TestSaver(t *testing.T) {
	expected := engine.Item{
		Url: "http://album.zhenai.com/u/1491957333",
		Type: "zhenai",
		Id: "1491957333",
		Payload: model.Profile{
			Name: "111",
			Gender: "男",
			IntroduceContent: "111111" ,
		},
	}

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://47.97.163.47:9200"),
	)
	if err != nil {
		panic(err)
	}
	// save item
	const index = "dating_profile"
	err = save(client, index, expected)
	if err != nil {
		panic(err)
	}

    resp, err := client.Get().
    	Index(index).
    	Type(expected.Type).
    	Id(expected.Id).
    	Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)

    var actual engine.Item
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		panic(err)
	}
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	t.Logf("%+v", actual)

}