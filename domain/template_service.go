package domain

import (
	"bytes"
	"errors"
	"text/template"
)

var (
	ErrTemplateNotExists = errors.New("template not exists")
)

type TemplateService struct {
	templateRepository TemplateRepository
}

func NewTemplateService(templateRepository TemplateRepository) *TemplateService {
	return &TemplateService{templateRepository: templateRepository}
}

func (s *TemplateService) CreateTemplate(name, tpl string) (*Template, error) {
	t := CreateTemplate(name, tpl)

	if err := s.templateRepository.Save(t); err != nil {
		return nil, err
	}

	return t, nil
}

func (s *TemplateService) SearchTemplates(pageable Pageable) (TemplatePage, error) {
	return s.templateRepository.FindAll(pageable)
}

func (s *TemplateService) Render(templateId string, ctx map[string]string) (string, error) {
	t, err := s.getTemplate(templateId)
	if err != nil {
		return "", err
	}

	tpl, err := template.New(t.Name).Parse(t.Template)
	if err != nil {
		return "", err
	}

	buffer := bytes.NewBuffer([]byte{})
	if err = tpl.Execute(buffer, ctx); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func (s *TemplateService) getTemplate(id string) (*Template, error) {
	tpl, err := s.templateRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	if tpl == nil {
		return nil, ErrTemplateNotExists
	}

	return tpl, nil
}
