package domain

type NotificationPage struct {
	Items      []Notification
	TotalCount int
	Page       int
	Size       int
}

type NotificationRepository interface {
	Save(notification *Notification) error
	FindById(id string) (*Notification, error)
	FindAll(pageable Pageable) (NotificationPage, error)
}
