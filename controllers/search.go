package controllers

import (
	"github.com/astaxie/beego"
	"organ-go-api/utils"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search/query"
	"encoding/json"
	"fmt"
	"strconv"
)

// 索引搜索相关接口
type SearchController struct {
	beego.Controller
}

// URLMapping ...
func (s *SearchController) URLMapping() {
	s.Mapping("ClientSearchOrgan", s.ClientSearchOrgan)   // 客户端程序用于搜索机构 get 桌面程序
	s.Mapping("CommonSearch", s.CommonSearch)             // 搜索机构 post
	s.Mapping("ClientSearchSchool", s.ClientSearchSchool) // 客户端程序用于搜索学校 get
}

// 客户端程序用于搜索机构
// @Title ClientSearchOrgan
// @Description 客户端程序用于搜索机构
// @Param	organname	query 	string	true	"The key for search"
// @Success 200 {object} utils.SearchResult1
// @Failure 403
// @router /search [get]
func (s *SearchController) ClientSearchOrgan() {
	fmt.Println("--------- GET /organindexsearch/search ----------")
	// 获取搜索参数
	organname := s.GetString("organname") // 获取请求的query
	// 包装模拟的请求流 example 上的请求方式，本应前端传递，但旧项目接口已规范，用此种方式处理
	MockRequestBody := []byte(`{
		"size": 10,
		"from": 0,
		"explain": false,
		"highlight":{},
		"query": {
			"boost": 1.0,
			"query": "` + organname + `"
		},
		"fields": ["uuid", "nameEn","nameOrigin", "nameCn", "shortName"]
	}`)

	var searchRequest bleve.SearchRequest

	if err := json.Unmarshal(MockRequestBody, &searchRequest); err == nil {
		// 校验query
		if srqv, ok := searchRequest.Query.(query.ValidatableQuery); ok {
			err = srqv.Validate()
			if err != nil {
				fmt.Println(err.Error())
				s.Data["json"] = err.Error()
				return
			}
		}
		searchResponse, err := utils.CurrentIndex.Index.Search(&searchRequest) // 通过构建的请求条件进行搜索
		if err != nil {
			fmt.Println(err.Error())
			s.Data["json"] = err.Error()
		} else {
			// 对匹配结果集合的数据中大于平均分的取出来
			var scoreSum float64 // 取出总分
			for _, v := range searchResponse.Hits {
				scoreSum += v.Score
			}
			average := scoreSum / float64(len(searchResponse.Hits)) // 求出平均值
			matchList := make([]int, 0)                         // 初始化匹配数组
			// 循环出分数大于平均值的, 遍历得分高的结果, 进行输出
			for i, v := range searchResponse.Hits {
				if v.Score >= average {
					matchList = append(matchList, i)
				}
			}

			// 对结果集进行包装
			var ResultList []utils.List1

			for _, v := range matchList {
				ResultList = append(ResultList, utils.List1{
					searchResponse.Hits[v].Fields["uuid"].(string),
					searchResponse.Hits[v].Fields["nameEn"].(string),
				})
			}

			// 最终结果集
			Result := utils.SearchResult1{
				TotalCount: int64(len(matchList)),
				TotalPage:  1,
				Message:    "搜索完成",
				Success:    searchResponse.Status.Successful,
				SearchTime: strconv.Itoa(int(searchResponse.Took)/1000000) + "ms",
				List: ResultList,
			}

			s.Data["json"] = Result

		}
	} else {
		fmt.Println(err.Error())
		s.Data["json"] = err.Error()
	}
	s.ServeJSON()
}

// 搜索机构
// @Title CommonSearch
// @Description 搜索机构
// @Param	body		body 	bleve.SearchRequest 	true		"body for query"
// @Success 200 {object} bleve.SearchRequest
// @Failure 403
// @router /search [post]
func (s *SearchController) CommonSearch() {
	fmt.Println("--------- POST /organindexsearch/search ----------")
	s.Data["json"] = "目前此接口暂无用到"
	s.ServeJSON()
}

// 客户端程序用于搜索学校 get
// @Title ClientSearchSchool
// @Description 客户端程序用于搜索学校
// @Param	schoolname	query 	string 	true		"schoolname for query"
// @Success 200 {object} []utils.List2
// @Failure 403
// @router /searchschool [get]
func (s *SearchController) ClientSearchSchool() {
	fmt.Println("--------- GET /organindexsearch/searchschool ----------")

	// 获取搜索参数
	schoolname := s.GetString("schoolname") // 获取请求的query
	// 包装模拟的请求流 example 上的请求方式，本应前端传递，但旧项目接口已规范，用此种方式处理
	MockRequestBody := []byte(`{
		"size": 10,
		"from": 0,
		"explain": false,
		"highlight":{},
		"query": {
			"boost": 1.0,
			"query": "` + schoolname + `"
		},
		"fields": ["uuid", "nameEn","nameOrigin", "nameCn", "shortName"]
	}`)

	var searchRequest bleve.SearchRequest

	if err := json.Unmarshal(MockRequestBody, &searchRequest); err == nil {
		// 校验query
		if srqv, ok := searchRequest.Query.(query.ValidatableQuery); ok {
			err = srqv.Validate()
			if err != nil {
				fmt.Println(err.Error())
				s.Data["json"] = err.Error()
				return
			}
		}
		searchResponse, err := utils.CurrentIndex.Index.Search(&searchRequest) // 通过构建的请求条件进行搜索
		if err != nil {
			fmt.Println(err.Error())
			s.Data["json"] = err.Error()
		} else {
			// 对匹配结果集合的数据中大于平均分的取出来
			var scoreSum float64 // 取出总分
			for _, v := range searchResponse.Hits {
				scoreSum += v.Score
			}
			average := scoreSum / float64(len(searchResponse.Hits)) // 求出平均值
			matchList := make([]int, 0)                         // 初始化匹配数组
			// 循环出分数大于平均值的, 遍历得分高的结果, 进行输出
			for i, v := range searchResponse.Hits {
				if v.Score >= average {
					matchList = append(matchList, i)
				}
			}

			// 对结果集进行包装
			var ResultList []utils.List2

			for _, v := range matchList {
				ResultList = append(ResultList, utils.List2{
					searchResponse.Hits[v].Fields["uuid"].(string),
					searchResponse.Hits[v].Fields["nameEn"].(string),
					searchResponse.Hits[v].Fields["countryCode"].(string),
				})
			}

			s.Data["json"] = ResultList

		}
	} else {
		fmt.Println(err.Error())
		s.Data["json"] = err.Error()
	}
	s.ServeJSON()
}
