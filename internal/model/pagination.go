package model

import (
	"github.com/vshakitskiy/reddit_comments/pkg/pagination"
)

type PaginationInput struct {
	Limit int32 `json:"limit"`
	Page  int32 `json:"page"`
}

func (p *PaginationInput) ToPagination() pagination.Pagination {
	return pagination.Pagination{
		Limit: p.Limit,
		Page:  p.Page,
	}
}

func PaginationToConnectionMeta(
	p pagination.Pagination,
) *ConnectionMeta {
	return &ConnectionMeta{
		Limit:      p.GetLimit(),
		Page:       p.GetPage(),
		TotalRow:   int32(p.TotalRows),
		TotalPages: int32(p.TotalPages),
	}
}
