package parser

import (
	"crawler/engine"
	"fmt"

	// "fmt"
	"regexp"
)

const cityListRe =`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`
func ParserCityList (contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	limit := 10
	for _, m := range matches {
		result.Items = append(result.Items, "City " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: ParseCity,
		})
		 fmt.Printf("City: %s, Url: %s\n", m[2], m[1])
	     limit--
	     if limit == 0 {
	     	break
		 }

		}

	return result
	// fmt.Printf("Matches found: %d\n", len(matches))
}