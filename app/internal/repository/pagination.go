package repository

type Pagination struct {
	Page  int
	Limit int
}

func (p *Pagination) IfDefaultPage() {
	if p.Limit <= 0 {
		p.Limit = 10
	}

	if p.Page <= 0 {
		p.Page = 1
	}
}

func (p *Pagination) GetLimit() int {
	return p.Limit
}

func (p *Pagination) GetOffset() int {
	return (p.Page - 1) * p.Limit
}
