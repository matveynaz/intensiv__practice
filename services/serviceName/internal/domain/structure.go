package domain

import (
	"errors"
	"regexp"
)

type Contact struct {
	Identifier  string
	FirstName   string
	LastName    string
	Patronymic  string
	PhoneNumber string
}

func (c *Contact) FullName() string {
	return c.LastName + " " + c.FirstName + " " + c.Patronymic
}

func NewContact(identifier, firstName, lastName, patronymic, phoneNumber string) (*Contact, error) {
	// Check if the phone number contains only numbers
	if !isValidPhoneNumber(phoneNumber) {
		return nil, errors.New("phone number must contain numbers only")
	}

	return &Contact{
		Identifier:  identifier,
		FirstName:   firstName,
		LastName:    lastName,
		Patronymic:  patronymic,
		PhoneNumber: phoneNumber,
	}, nil
}

func isValidPhoneNumber(phoneNumber string) bool {
	// Use regular expression to match numbers only
	match, _ := regexp.MatchString("^[0-9]+$", phoneNumber)
	return match
}
