package utils

// list 实体
type List1 struct {
	Uuid   string            `json:"uuid"`
	NameEn string            `json:"nameEn"`
}

// SearchResult 实体
type SearchResult1 struct {
	List       []List1           `json:"list"`
	TotalCount int64             `json:"totalCount"`
	TotalPage  int64             `json:"totalPage"`
	Message    string            `json:"message"`
	Success    int               `json:"success"`
	SearchTime string            `json:"searchTime"`
}

// 客户端程序用于搜索学校 返回的 list 实体
type List2 struct {
	Uuid        string             `json:"uuid"`
	NameEn      string             `json:"nameEn"`
	CountryCode string             `json:"countryCode"`
}
