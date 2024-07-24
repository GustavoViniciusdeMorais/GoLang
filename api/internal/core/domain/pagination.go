package domain

import "strconv"

type Pagination struct {
	Page     string
	Limit    string
	LimitInt int
	Offset   int
}

func NewPagination(page string, limit string) *Pagination {
	return &Pagination{
		Page:  page,
		Limit: limit,
	}
}

func (p *Pagination) CalculatePagination() *Pagination {
	page, err := strconv.Atoi(p.Page)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(p.Limit)
	if err != nil || limit < 1 {
		limit = 25
	}
	p.LimitInt = limit

	p.Offset = (page - 1) * limit

	return p
}
