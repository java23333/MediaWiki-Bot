package Plugin

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"log"
	"strings"
	"xyz.nyan/MediaWiki-Bot/MediaWikiAPI"
)

func Error(WikiLink string, title string) string {
	return "在[" + WikiLink + "]中找不到[" + title + "]哦，请检查输入是否正确"
}

//搜索wiki
func SearchWiki(WikiName string, title string) string {
	SearchInfo, _ := MediaWikiAPI.Opensearch(WikiName, title)
	if len(SearchInfo) != 0 {
		SearchList := SearchInfo[1].([]interface{})
		if len(SearchList) != 0 {
			var SearchPages strings.Builder
			for _, value := range SearchList {
				PagseName := "[" + value.(string) + "]"
				SearchPages.WriteString(PagseName)
				SearchPages.WriteString("\n")
			}
			return SearchPages.String()
		}
		return ""
	} else {
		return ""
	}
}

//为空处理
func NilProcessing(WikiName string, title string) string {
	SearchInfo := SearchWiki(WikiName, title)
	if SearchInfo != "" {
		Info := "没有找到页面，但是找到了以下结果\n" + SearchInfo + "请使用 " + WikiName + ":页面标题 重新查询"
		return Info
	} else {
		WikiLink := MediaWikiAPI.GetWikiLink(WikiName)
		return Error(WikiLink, title)
	}
}

//获取Wiki页面标题，过滤后缀
func GetUrlTitle(WikiName string, PageName string) string {
	WikiLink := MediaWikiAPI.GetWikiLink(WikiName)
	doc, err := htmlquery.LoadURL(WikiLink + "/" + PageName)
	if err != nil {
		fmt.Println(err)
	}
	for _, n := range htmlquery.Find(doc, "/html/head/title") {
		PageTitle := htmlquery.OutputHTML(n, false)
		countSplit := strings.SplitN(PageTitle, " - ", 2)
		Title := countSplit[0]
		return Title
	}
	return ""
}

//获取Wiki页面信息
func QueryWikiInfo(WikiName string, title string) (interface{}, error) {
	info, err := MediaWikiAPI.QueryInfoUrl(WikiName, title)
	pagesIdInfo := info["query"].(map[string]interface{})["pages"]
	var PageId string
	for one := range pagesIdInfo.(map[string]interface{}) {
		PageId = one
	}

	return info["query"].(map[string]interface{})["pages"].(map[string]interface{})[PageId], err
}

//查询页面是否存在重定向
func QueryRedirects(WikiName string, title string) (whether bool, to string, from string, err error) {
	info, err := MediaWikiAPI.QueryRedirects(WikiName, title)

	_, ok := info["query"]
	if ok {
		normalized, ok := info["query"].(map[string]interface{})["normalized"]
		if ok {
			return true, normalized.([]interface{})[0].(map[string]interface{})["to"].(string), normalized.([]interface{})[0].(map[string]interface{})["from"].(string), err
		} else {
			PageTitleInfo := GetUrlTitle(WikiName, title)
			if PageTitleInfo != title {
				ToTitle := PageTitleInfo
				return true, ToTitle, title, err
			}
		}
		return false, "", "", err
	}
	return false, "", "", err
}

//获取Wiki页面信息
func GetWikiInfo(WikiName string, title string) (string, error) {
	var err error
	RedirectsState, ToTitle, FromTitle, _ := QueryRedirects(WikiName, title)
	var info map[string]interface{}
	if RedirectsState {
		info, err = MediaWikiAPI.QueryExtracts(WikiName, 100, ToTitle)
	} else {
		info, err = MediaWikiAPI.QueryExtracts(WikiName, 100, title)
	}

	_, ok := info["query"]
	if !ok {
		return NilProcessing(WikiName, title), err
	}

	pagesIdInfo, ok := info["query"].(map[string]interface{})["pages"]
	if !ok {
		return NilProcessing(WikiName, title), err
	}

	var PageId string
	for one := range pagesIdInfo.(map[string]interface{}) {
		PageId = one
	}

	if PageId != "-1" {
		PagesExtract := info["query"].(map[string]interface{})["pages"].(map[string]interface{})[PageId].(map[string]interface{})["extract"]
		var returnText string
		if RedirectsState {
			WikiPageInfo, err := QueryWikiInfo(WikiName, ToTitle)
			if err != nil {
				log.Println(err)
			}
			WikiPageLink := WikiPageInfo.(map[string]interface{})["fullurl"].(string)
			returnText = WikiPageLink + "\n(重定向[" + FromTitle + "]->[" + ToTitle + "])\n" + PagesExtract.(string)
		} else {
			WikiPageInfo, err := QueryWikiInfo(WikiName, title)
			if err != nil {
				log.Println(err)
			}
			WikiPageLink := WikiPageInfo.(map[string]interface{})["fullurl"].(string)
			returnText = WikiPageLink + "\n[" + title + "]\n" + PagesExtract.(string)
		}
		return returnText, err
	} else {
		return NilProcessing(WikiName, title), err
	}
}
