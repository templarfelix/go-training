package entity

import "errors"

var (
	ErrNotFound              = errors.New("Not found")
	ErrInvalidEntity         = errors.New("Invalid entity")
	ErrValueInvalidEntity    = errors.New("Value entity")
	ErrCannotBeDeleted       = errors.New("Cannot Be Deleted")
	ErrDocumentNumberInvalid = errors.New("Document Number invalid")
)
