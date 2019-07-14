package parser

import (
	"crawler/engine"
	"fmt"

	// "fmt"
	"regexp"
)

const cityListRe =`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`
func ParserCityList (contents []byte, _ string) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: ParseCity,
		})
		 fmt.Printf("City: %s, Url: %s\n", m[2], m[1])
		}

	return result
}