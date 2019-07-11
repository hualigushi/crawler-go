package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
)

var basicInfoRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)</div>`)
var detailInfoRe = regexp.MustCompile(`<div class="m-btn pink" [^>]*>([^<]+)</div>`)
var objectInfoRe = regexp.MustCompile(`<div class="m-btn" [^>]*>([^<]+)</div>`)
var introduceContentRe = regexp.MustCompile(`<div class="m-content-box m-des" [^>]*><span [^>]*>([^<]+)</span></div>`)
var genderRe = regexp.MustCompile(`<div class="m-title" [^>]*>([^<]{1})的动态</div>`)
func ParseProfile (contents []byte, name string) engine.ParserResult{
	profile := model.Profile{}
	profile.Name = name

	matches := basicInfoRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		profile.BasicInfo = append(profile.BasicInfo, string(m[1]))
	}

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

	result := engine.ParserResult{
		Items: []interface{}{profile},
	}

	return result
}

