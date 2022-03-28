package util

type Pager struct {
	Page     uint64 `json:"page" form:"page"`
	PageSize uint64 `json:"page_size" form:"page_size"`
	Total    uint64 `json:"total"`
}

func (p *Pager) Offset() uint64 {
	p.Secure()
	return (p.Page - 1) * p.PageSize
}

func (p *Pager) Limit() uint64 {
	p.Secure()
	return p.PageSize
}

func (p *Pager) Secure() *Pager {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 20
	}
	if p.PageSize > 500 {
		p.PageSize = 500
	}
	return p
}

func NewRequestPager(page, pageSize uint64) *Pager {
	pager := &Pager{
		Page:     page,
		PageSize: pageSize,
	}
	pager.Secure()
	return pager
}
