elasticsearch go �ͻ���
F:\Go\src>go get -v github.com/olivere/elastic

postman

ɾ�������ĵ�
POST 47.97.163.47:9200/dating_profile/zhenai/_delete_by_query
body : 
{
    "query": {
        "match_all": {}
    }
}

��ѯ
GET 47.97.163.47:9200/dating_profile/zhenai/_search?size=100&q=Age:(<35)