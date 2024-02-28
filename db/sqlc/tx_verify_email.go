package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
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
		result.VerifyEmail, err = q.UpdateVerifyEmail(ctx, UpdateVerifyEmailParams{
			ID:         arg.EmailId,
			SecretCode: arg.SecretCode,
		})
		if err != nil {
			return err
		}
		result.User, err = q.UpdateUser(ctx, UpdateUserParams{
			Username: result.VerifyEmail.Username,
			IsEmailVerified: pgtype.Bool{
				Bool:  true,
				Valid: true,
			},
		})
		return err
	})
	return result, err
}
