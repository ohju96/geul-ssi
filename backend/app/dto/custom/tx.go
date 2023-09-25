package custom

import (
	"context"
	"geulSsi/config/db"

	"geulSsi/ent"
)

func TxGenerate(ctx context.Context) (tx *ent.Tx, customError *Error) {
	tx, err := db.MySQL.Tx(ctx)
	if err != nil {
		return nil, TxGenerateFail
	}
	return tx, nil
}

func TxRollback(tx *ent.Tx) *Error {
	if err := tx.Rollback(); err != nil {
		return TxRollbackFail
	}
	return nil
}

func TxCommit(tx *ent.Tx) *Error {
	if err := tx.Commit(); err != nil {
		TxRollback(tx)
		return TxCommitFail
	}
	return nil
}
