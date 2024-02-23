package db

import (
	"context"
)

// TransferTxParams  contains the input parameters of the transfer transaction
type VerifyEmailTxParams struct {
	EmailId    int64
	SecretCode string
}

// TransferTxResult is the result of the transfer transaction
type VerifyEmailTxResult struct {
	User        User
	VerifyEmail VerifyEmail
}

//CreateUserTx performs a money transfer from one account to the another account.
// It creates a transfer record, add account entries, and update accounts balance within a single database transaction

func (store *SQLStore) VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error) {
	var result VerifyEmailTxResult
	err := store.execTX(ctx, func(q *Queries) error {
		var err error
		return err
	})
	return result, err
}
