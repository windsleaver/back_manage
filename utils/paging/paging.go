package paging

import "fmt"

type PageResult struct {
	Page     int
	PageSize int
	Filter   string
}

func PageHandle(page, pageSize int, filter string) *PageResult {
	var _filter string
	if len(filter) > 0 {
		_filter = fmt.Sprintf("%s%s%s", "%", filter, "%")
	}
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	} else if pageSize >= 1000 {
		pageSize = 1000
	}
	page = (page - 1) * pageSize
	return &PageResult{
		Page:     page,
		PageSize: pageSize,
		Filter:   _filter,
	}
}

func PageFilter(filter string) string {
	if len(filter) > 0 {
		filter = fmt.Sprintf("%s%s%s", "%", filter, "%")
	}
	return filter
}
