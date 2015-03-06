package paysimple

import (
	"net/url"
	"strconv"
)

type Filters interface {
	Values() url.Values
}

type PaginationFilters struct {
	Page, PageSize int
}

var _ Filters = PaginationFilters{}

// Values generates GET parameters from the given filters
func (f PaginationFilters) Values() url.Values {
	values := url.Values{}
	if f.Page != 0 {
		values.Set("page", strconv.Itoa(f.Page))
	}
	if f.PageSize != 0 {
		values.Set("pagesize", strconv.Itoa(f.PageSize))
	}
	return values
}
