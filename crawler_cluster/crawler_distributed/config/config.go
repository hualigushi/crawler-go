package config

const (
	// Parser Name
	ParseCity = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile = "ParseProfile"
	NilParser = "NilParser"

	// ElasticSearch
	Elasticindex = "dating_profile"

	// RPC Endpoints
	ItemSaverRpc = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	// Rate limiting
	Qps = 20
)