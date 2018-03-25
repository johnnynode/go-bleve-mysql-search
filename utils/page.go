package utils

const (
	navigatePages = 8 // 私有
)

/* 用于分页的工具 */

type Page struct {
	List                interface{}   `json:"list"`                // 对象纪录的结果集
	Total               int64         `json:"total"`               // 总纪录数
	Limit               int64         `json:"limit"`               // 每页显示纪录数
	Pages               int64         `json:"pages"`               // 总页数
	PageNumber          int64         `json:"pageNumber"`          // 当前页
	NavigatePageNumbers []int64       `json:"navigatePageNumbers"` // 所有导航页码号
	FirstPage           bool          `json:"firstPage"`           // 是否为第一页
	LastPage            bool          `json:"lastPage"`            // 是否为最后一页
}

// 计算导航页
func CalcNavigatePageNumbers(pageNumber int64, pages int64) []int64 {
	var navigatePageNumbers = make([]int64, 8)
	//当总页数小于或等于导航页码数时
	if pages <= navigatePages {
		var i int64 = 0
		for i < pages {
			navigatePageNumbers[i] = i + 1
			i++
		}
		return navigatePageNumbers[:pages]
	}

	//当总页数大于导航页码数时
	startNum := pageNumber - navigatePages/2
	endNum := pageNumber + navigatePages/2

	if startNum < 1 {
		startNum = 1 // 最前navPageCount页
		var j int64 = 0
		for j < navigatePages {
			navigatePageNumbers[j] = startNum
			startNum++ // 自增1
			j++
		}
	} else if endNum > pages {
		endNum = pages
		//最后navPageCount页
		var k int64 = navigatePages - 1
		for k >= 0 {
			endNum--
			navigatePageNumbers[k] = endNum
			k --
		}
	} else {
		// 所有中间页
		var q int64 = 0
		for q < navigatePages {
			startNum ++
			navigatePageNumbers[q] = startNum
			q ++
		}
	}
	return navigatePageNumbers[:navigatePages]
}

// 计算pages
func GetPages(total int64, limit int64) int64 {
	return (total-1)/limit + 1
}

// 是否是首页
func IsFirstPage(currentPage int64) bool {
	return currentPage == 1
}

// 是否是尾页
func IsLastPage(currentPage int64, pages int64) bool {
	return currentPage == pages && currentPage != 1
}
