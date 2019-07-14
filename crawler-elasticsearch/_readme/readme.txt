elasticsearch go 客户端
F:\Go\src>go get -v github.com/olivere/elastic

postman

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
