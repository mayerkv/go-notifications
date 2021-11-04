package domain

type NotificationService struct {
	notificationRepository NotificationRepository
	templateService        *TemplateService
}

func NewNotificationService(notificationRepository NotificationRepository, templateService *TemplateService) *NotificationService {
	return &NotificationService{notificationRepository: notificationRepository, templateService: templateService}
}

func (s *NotificationService) SendEmail(email, templateId string, ctx map[string]string) error {
	msg, err := s.templateService.Render(templateId, ctx)
	if err != nil {
		return err
	}

	notification := CreateNotification(email, msg, NotificationTypeEmail)

	return s.notificationRepository.Save(notification)
}

func (s *NotificationService) SearchNotifications(pageable Pageable) (NotificationPage, error) {
	return s.notificationRepository.FindAll(pageable)
}
