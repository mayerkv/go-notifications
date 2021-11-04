package grpc_service

import (
	"context"
	"github.com/mayerkv/go-notifications/domain"
)

type NotificationsServiceImpl struct {
	templateService     *domain.TemplateService
	notificationService *domain.NotificationService
}

func NewNotificationsServiceImpl(templateService *domain.TemplateService, notificationService *domain.NotificationService) NotificationsServiceServer {
	return &NotificationsServiceImpl{templateService: templateService, notificationService: notificationService}
}

func (s *NotificationsServiceImpl) SendEmail(ctx context.Context, request *SendEmailRequest) (*Empty, error) {
	err := s.notificationService.SendEmail(request.Email, request.TemplateId, request.Context)
	if err != nil {
		return nil, err
	}

	return &Empty{}, nil
}

func (s *NotificationsServiceImpl) CreateTemplate(ctx context.Context, request *CreateTemplateRequest) (*CreateTemplateResponse, error) {
	template, err := s.templateService.CreateTemplate(request.Name, request.Template)
	if err != nil {
		return nil, err
	}

	return &CreateTemplateResponse{
		Id: template.Id,
	}, nil
}

// SearchTemplates todo implement
func (s *NotificationsServiceImpl) SearchTemplates(ctx context.Context, request *SearchTemplatesRequest) (*SearchTemplatesResponse, error) {
	templates, err := s.templateService.SearchTemplates(domain.Pageable{
		Page:           int(request.Page),
		Size:           int(request.Size),
		OrderBy:        mapTemplateOrderBy(request.OrderBy),
		OrderPredicate: mapOrderDirection(request.OrderDirection),
	})
	if err != nil {
		return nil, err
	}

	return &SearchTemplatesResponse{
		List:  mapTemplates(templates.Items),
		Count: int32(templates.TotalCount),
	}, nil
}

func (s *NotificationsServiceImpl) SearchNotifications(ctx context.Context, request *SearchNotificationsRequest) (*SearchNotificationsResponse, error) {
	notifications, err := s.notificationService.SearchNotifications(domain.Pageable{
		Page:           int(request.Page),
		Size:           int(request.Size),
		OrderBy:        mapNotificationOrderBy(request.OrderBy),
		OrderPredicate: mapOrderDirection(request.OrderDirection),
	})
	if err != nil {
		return nil, err
	}

	return &SearchNotificationsResponse{
		List:  mapNotifications(notifications.Items),
		Count: int32(notifications.TotalCount),
	}, nil
}

func (s *NotificationsServiceImpl) mustEmbedUnimplementedNotificationsServiceServer() {
	panic("implement me")
}
