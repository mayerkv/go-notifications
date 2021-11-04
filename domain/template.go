package domain

import "github.com/google/uuid"

type Template struct {
	Id       string
	Name     string
	Template string
}

func CreateTemplate(name, template string) *Template {
	return &Template{
		Id:       uuid.NewString(),
		Name:     name,
		Template: template,
	}
}
