package entity

import "errors"

var ErrNotFound = errors.New("Not found")

var ErrInvalidEntity = errors.New("Invalid entity")

var ErrCannotBeDeleted = errors.New("Cannot Be Deleted")

var ErrNotEnoughBooks = errors.New("Not enough books")

var ErrBookAlreadyBorrowed = errors.New("Book already borrowed")

var ErrBookNotBorrowed = errors.New("Book not borrowed")