package repository

import (
	"github.com/mayerkv/go-notifications/domain"
	"sync"
)

type InMemoryNotificationRepository struct {
	sync.Mutex
	items map[string]domain.Notification
}

func NewInMemoryNotificationRepository() domain.NotificationRepository {
	return &InMemoryNotificationRepository{
		items: map[string]domain.Notification{},
	}
}

func (r *InMemoryNotificationRepository) Save(notification *domain.Notification) error {
	r.Lock()
	defer r.Unlock()

	r.items[notification.Id] = *notification

	return nil
}

func (r *InMemoryNotificationRepository) FindById(id string) (*domain.Notification, error) {
	r.Lock()
	defer r.Unlock()

	if n, ok := r.items[id]; ok {
		return &n, nil
	}

	return nil, nil
}

func (r *InMemoryNotificationRepository) FindAll(pageable domain.Pageable) (domain.NotificationPage, error) {
	r.Lock()
	defer r.Unlock()

	var notifications = make([]domain.Notification, 0)
	for _, item := range r.items {
		notifications = append(notifications, item)
	}

	return domain.NotificationPage{
		Items:      r.skip(notifications, pageable),
		TotalCount: len(notifications),
		Page:       pageable.Page,
		Size:       pageable.Size,
	}, nil
}

func (r *InMemoryNotificationRepository) skip(notifications []domain.Notification, pageable domain.Pageable) []domain.Notification {
	start := (pageable.Page - 1) * pageable.Size
	if start > len(notifications)-1 {
		return []domain.Notification{}
	}

	end := start + pageable.Size
	if end > len(notifications)-1 {
		return notifications[start:]
	}

	return notifications[start:end]
}
