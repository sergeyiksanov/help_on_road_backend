package services

type TransactionContext interface{}

type ITransactionManager interface {
	WithTransaction(fn func(ctx TransactionContext) error) error
}
