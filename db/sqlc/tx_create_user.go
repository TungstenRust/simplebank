package db

import (
	"context"
)

// TransferTxParams  contains the input parameters of the transfer transaction
type CreateUserTxParams struct {
}

// TransferTxResult is the result of the transfer transaction
type CreateUserTxResult struct {
}

//CreateUserTx performs a money transfer from one account to the another account.
// It creates a transfer record, add account entries, and update accounts balance within a single database transaction

func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult
	err := store.execTX(ctx, func(q *Queries) error {
		var err error

		return err
	})
	return result, err
}
