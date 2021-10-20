package entity

type Status int
type TransactionType int

const (
	ACTIVE Status = iota
	INACTIVE
	DELETED
)

const (
	CREDIT TransactionType = iota
	DEBIT
)
