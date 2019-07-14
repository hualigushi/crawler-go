package controler

import (
	"context"
	"crawler/engine"
	"crawler/frontend/model"
	"crawler/frontend/view"
	"fmt"
	"github.com/olivere/elastic"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL("http://47.97.163.47:9200"),
		)
	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view: view.CreateSearchResultView(
			template),
		client: client,
	}
}

// localhost:8888/search?q=已购房&from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}

	var page model.SearchResult
	page, err = h.getSearchResult(q, from)
	if err != nil {
		//http.Error(w, err.Error(),http.StatusBadRequest)
	}
	err = h.view.Render(w, page)
	if err != nil {
		//http.Error(w, err.Error(),http.StatusBadRequest)
	}
}


func (h SearchResultHandler) getSearchResult(q string, from int ) (model.SearchResult,error) {
	var result model.SearchResult
	result.Query = q
	//elastic 查询结果
	resp, err := h.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}
	result.Hits  = int(resp.TotalHits())
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result,nil
}


func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q,"Payload.$1:")
}



