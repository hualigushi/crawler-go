package parser

import (
	"crawler/engine"
	"crawler/model"
	"log"
	"regexp"
	"strconv"
)

var basicInfoRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)</div>`)
var detailInfoRe = regexp.MustCompile(`<div class="m-btn pink" [^>]*>([^<]+)</div>`)
var objectInfoRe = regexp.MustCompile(`<div class="m-btn" [^>]*>([^<]+)</div>`)
var introduceContentRe = regexp.MustCompile(`<div class="m-content-box m-des" [^>]*><span [^>]*>([^<]+)</span><!----></div>`)
var genderRe = regexp.MustCompile(`<div class="m-title" [^>]*>([^<]{1})的动态</div>`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
func ParseProfile (contents []byte, url string, name string) engine.ParserResult{
	profile := model.Profile{}
	profile.Name = name

	matches := basicInfoRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		profile.BasicInfo = append(profile.BasicInfo, string(m[1]))
	}
	age := profile.BasicInfo[1][0 : 2]
	intAge, err := strconv.Atoi(age)
	if err != nil {
		log.Printf("get age error")
	}
	profile.Age = intAge

	matches1 := detailInfoRe.FindAllSubmatch(contents, -1)
	for _, m := range matches1 {
		profile.DetailInfo = append(profile.DetailInfo, string(m[1]))
	}

	matches2 := objectInfoRe.FindAllSubmatch(contents, -1)
	for _, m := range matches2 {
		profile.ObjectInfo = append(profile.ObjectInfo, string(m[1]))
	}

	matches3 := introduceContentRe.FindAllSubmatch(contents, -1)
	for _, m := range matches3 {
		profile.IntroduceContent = string(m[1])
	}

	matches4 := genderRe.FindSubmatch(contents)

	if  len(matches4)>0 && string(matches4[1]) == "他" {
		profile.Gender = "男"
	} else {
		profile.Gender = "女"
	}

	matches5 := idUrlRe.FindSubmatch([]byte(url))

	result := engine.ParserResult{
		Items: []engine.Item{
			{
				Url: url,
				Type: "zhenai",
				Id: string(matches5[1]),
				Payload: profile,
			},
		},
	}

	return result
}


func ProfileParser (name string) engine.ParserFunc{
	return func (c []byte, url string) engine.ParserResult{
		return ParseProfile(c, url, name)
	}
}
