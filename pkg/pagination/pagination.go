package pagination

type Pagination struct {
	Limit      int32       `json:"limit,omitempty;query:limit"`
	Page       int32       `json:"page,omitempty;query:page"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

func (p *Pagination) GetPage() int32 {
	if p.Page == 0 {
		p.Page = 1
	}

	return p.Page
}

func (p *Pagination) GetLimit() int32 {
	if p.Limit == 0 {
		p.Limit = 10
	}

	return p.Limit
}

func (p *Pagination) GetOffset() int32 {
	return (p.GetPage() - 1) * p.GetLimit()
}
