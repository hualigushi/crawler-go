# crawler-go
go语言实现的爬虫程序

慕课网课程-Google资深工程师深度讲解Go语言的项目

爬取珍爱网数据

## crawler-single 单机版爬虫程序

运行： go run main.go

## crawler-goroutine 并发版爬虫程序

运行：go run main.go

## crawler-elasticsearch 数据存储到ealsticsearch中

1.阿里云(47.97.163.47)服务器安装docker: yum install docker

2.启动docker: service docker install

3.安装ealsticsearch 镜像: docker pull docker.elastic.co/elasticsearch/elasticsearch:7.2.0

4.启动elasticsearch: docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.2.0

postman查看 47.97.163.47:9200

5.运行: go run main.go

6.运行前端: go run frontend/start.go

7.查看前端界面: localhost: 8888
