package internal

import (
	"services/contact/internal/delivery"
	"services/contact/internal/repository"
	"services/contact/internal/usecase"
)

type ContactService struct {
	ContactRepo repository.ContactRepository
	GroupRepo   repository.GroupRepository
	ContactUC   usecase.ContactUseCase
	GroupUC     usecase.GroupUseCase
	HTTPHandler delivery.HTTPHandler
}

func NewContactService(contactRepo repository.ContactRepository, groupRepo repository.GroupRepository) *ContactService {
	contactUC := usecase.NewContactUseCase(contactRepo)
	groupUC := usecase.NewGroupUseCase(groupRepo)
	httpHandler := delivery.NewHTTPHandler(contactUC, groupUC)

	return &ContactService{
		ContactRepo: contactRepo,
		GroupRepo:   groupRepo,
		ContactUC:   contactUC,
		GroupUC:     groupUC,
		HTTPHandler: httpHandler,
	}
}
