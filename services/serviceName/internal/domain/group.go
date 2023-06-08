package domain

import "errors"

type Group struct {
	Identifier string
	Title      string
}

func NewGroup(identifier, title string) (*Group, error) {
	if len(title) > 250 {
		return nil, errors.New("name exceeds maximum length of 250 characters")
	}

	return &Group{
		Identifier: identifier,
		Title:      title,
	}, nil
}
