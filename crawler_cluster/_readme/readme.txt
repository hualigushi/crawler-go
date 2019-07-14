elasticsearch go 客户端
F:\Go\src>go get -v github.com/olivere/elastic

删除整个文档
POST 47.97.163.47:9200/dating_profile/zhenai/_delete_by_query
body : 
{
    "query": {
        "match_all": {}
    }
}

查询
GET 47.97.163.47:9200/dating_profile/zhenai/_search?size=100&q=Age:(<35)

service docker start
docker images
docker ps
docker kill id

F:\Go\src\crawler>go run crawler_distributed/persist/server/itemsaver.go --port=1234

F:\Go\src\crawler>go run crawler_distributed/worker/server/worker.go --port=9000
F:\Go\src\crawler>go run crawler_distributed/worker/server/worker.go --port=9001
F:\Go\src\crawler>go run crawler_distributed/worker/server/worker.go --port=9002

F:\Go\src\crawler>go run crawler_distributed/main.go --itemsaver_host=":1234" --worker_hosts=":9000,:9001,:9002"

F:\Go\src\crawler>go run crawler/frontend/start.go