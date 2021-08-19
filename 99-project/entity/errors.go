package entity

import "errors"

var (
	ErrNotFound = errors.New("Not found")
	ErrInvalidEntity = errors.New("Invalid entity")
	ErrCannotBeDeleted = errors.New("Cannot Be Deleted")
	ErrNotEnoughBooks = errors.New("Not enough books")
	ErrBookAlreadyBorrowed = errors.New("Book already borrowed")
	ErrBookNotBorrowed = errors.New("Book not borrowed")
	ErrDocumentNumberInvalid = errors.New("Document Number invalid")
)