package domain

type TemplatePage struct {
	Items      []Template
	TotalCount int
	Page       int
	Size       int
}

type TemplateRepository interface {
	Save(template *Template) error
	FindById(id string) (*Template, error)
	FindAll(pageable Pageable) (TemplatePage, error)
}
