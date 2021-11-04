package grpc_service

import (
	"github.com/mayerkv/go-notifications/domain"
	"time"
)

func mapOrderDirection(direction OrderDirection) domain.Predicate {
	if direction == OrderDirection_DESC {
		return domain.PredicateDesc
	}

	return domain.PredicateAsc
}

func mapTemplateOrderBy(orderBy SearchTemplatesRequest_OrderBy) string {
	if orderBy == SearchTemplatesRequest_NAME {
		return "name"
	}

	return ""
}

func mapTemplates(items []domain.Template) []*Template {
	var res = make([]*Template, len(items))
	for _, item := range items {
		res = append(res, mapTemplate(item))
	}

	return res
}

func mapTemplate(item domain.Template) *Template {
	return &Template{
		Id:       item.Id,
		Name:     item.Name,
		Template: item.Template,
	}
}

func mapNotifications(items []domain.Notification) []*Notification {
	var res = make([]*Notification, len(items))
	for _, item := range items {
		res = append(res, mapNotification(item))
	}

	return res
}

func mapNotification(item domain.Notification) *Notification {
	return &Notification{
		Id:        item.Id,
		Addressee: item.Addressee,
		Message:   item.Message,
		Type:      mapNotificationType(item.Type),
		CreatedAt: item.CreatedAt.Format(time.RFC3339),
	}
}

func mapNotificationType(notificationType domain.NotificationType) Notification_Type {
	if notificationType == domain.NotificationTypeEmail {
		return Notification_EMAIL
	}

	return 0
}

func mapNotificationOrderBy(by SearchNotificationsRequest_OrderBy) string {
	if by == SearchNotificationsRequest_CREATED_AT {
		return "createdAt"
	}

	return ""
}
