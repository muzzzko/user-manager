package respository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github/user-manager/tools/logger"
)

func PerformTransaction(ctx context.Context, pool *pgxpool.Pool, perform func(tx pgx.Tx) error) error {
	var err error
	ctxLogger := logger.GetFromContext(ctx)

	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("fail to begin txn: %w", err)
	}

	defer func() {
		if err == nil {
			if txerr := tx.Commit(ctx); txerr != nil {
				ctxLogger.WithError(txerr).Error("fail commit txn")
			}

			return
		}

		if txerr := tx.Rollback(ctx); txerr != nil {
			ctxLogger.WithError(txerr).Error("fail rollback txn")
		}
	}()

	return perform(tx)
}
