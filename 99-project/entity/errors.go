package entity

import "errors"

var (
	ErrNotFound              = errors.New("Not found")
	ErrInvalidEntity         = errors.New("Invalid entity")
	ErrCannotBeDeleted       = errors.New("Cannot Be Deleted")
	ErrDocumentNumberInvalid = errors.New("Document Number invalid")
)
