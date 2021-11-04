package repository

import (
	"github.com/mayerkv/go-notifications/domain"
	"sync"
)

type InMemoryTemplateRepository struct {
	sync.Mutex
	items map[string]domain.Template
}

func NewInMemoryTemplateRepository() domain.TemplateRepository {
	return &InMemoryTemplateRepository{
		items: map[string]domain.Template{},
	}
}

func (r *InMemoryTemplateRepository) Save(template *domain.Template) error {
	r.Lock()
	defer r.Unlock()

	r.items[template.Id] = *template

	return nil
}

func (r *InMemoryTemplateRepository) FindById(id string) (*domain.Template, error) {
	r.Lock()
	defer r.Unlock()

	if template, ok := r.items[id]; ok {
		return &template, nil
	}

	return nil, nil
}

func (r *InMemoryTemplateRepository) FindAll(pageable domain.Pageable) (domain.TemplatePage, error) {
	r.Lock()
	defer r.Unlock()

	var templates = make([]domain.Template, len(r.items))
	for _, item := range r.items {
		templates = append(templates, item)
	}

	return domain.TemplatePage{
		Items:      r.skip(templates, pageable),
		TotalCount: len(templates),
		Page:       pageable.Page,
		Size:       pageable.Size,
	}, nil
}

func (r *InMemoryTemplateRepository) skip(templates []domain.Template, pageable domain.Pageable) []domain.Template {
	start := (pageable.Page - 1) * pageable.Size
	if start > len(templates)-1 {
		return []domain.Template{}
	}

	end := start + pageable.Size
	if end > len(templates)-1 {
		return templates[start:]
	}

	return templates[start:end]
}
