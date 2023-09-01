package db

import "context"

type CreateUserTxParams struct {
	CreateUserParams
}

type CreateUserTxResult struct {
	User User
	Cart Cart
}

func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = store.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			return err
		}

		result.Cart, err = store.CreateCart(ctx, arg.CreateUserParams.Username)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}
