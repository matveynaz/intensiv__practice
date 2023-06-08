package repository

import "services/contact/internal/domain"

type ContactRepository interface {
	GetByID(identifier string) (*domain.Contact, error)
	Create(contact *domain.Contact) error
	Update(contact *domain.Contact) error
	Delete(identifier string) error
}

type GroupRepository interface {
	GetByID(identifier string) (*domain.Group, error)
	Create(group *domain.Group) error
	AddContactToGroup(contactID, groupID string) error
}
