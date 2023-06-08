package usecase

import (
	"services/contact/internal/domain"
	"services/contact/internal/repository"
)

type ContactUseCase interface {
	GetContact(identifier string) (*domain.Contact, error)
	CreateContact(contact *domain.Contact) error
	UpdateContact(identifier string, contact *domain.Contact) error
	DeleteContact(identifier string) error
}

type GroupUseCase interface {
	GetGroup(identifier string) (*domain.Group, error)
	CreateGroup(group *domain.Group) error
	AddContactToGroup(contactID, groupID string) error
}

func NewContactUseCase(contactRepo repository.ContactRepository, groupRepo repository.GroupRepository) *ContactUseCase {
	return &ContactUseCase{
		contactRepo: contactRepo,
		groupRepo:   groupRepo,
	}
}

func (uc *ContactUseCase) GetContact(identifier string) (*domain.Contact, error) {
	return uc.contactRepo.GetByID(identifier)
}

func (uc *ContactUseCase) CreateContact(contact *domain.Contact) error {
	return uc.contactRepo.Create(contact)
}

func (uc *ContactUseCase) UpdateContact(identifier string, contact *domain.Contact) error {
	existingContact, err := uc.contactRepo.GetByID(identifier)
	if err != nil {
		return err
	}

	existingContact.FirstName = contact.FirstName
	existingContact.LastName = contact.LastName
	existingContact.Patronymic = contact.Patronymic
	existingContact.PhoneNumber = contact.PhoneNumber

	return uc.contactRepo.Update(existingContact)
}

func (uc *ContactUseCase) DeleteContact(identifier string) error {
	return uc.contactRepo.Delete(identifier)
}

func (uc *ContactUseCase) CreateGroup(group *domain.Group) error {
	return uc.groupRepo.Create(group)
}

func (uc *ContactUseCase) GetGroup(identifier string) (*domain.Group, error) {
	return uc.groupRepo.GetByID(identifier)
}

func (uc *ContactUseCase) AddContactToGroup(contactID, groupID string) error {
	return uc.groupRepo.AddContactToGroup(contactID, groupID)
}
