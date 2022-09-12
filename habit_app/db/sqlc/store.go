package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Contain all functions to execute db queries and transaction
type Store struct {
	*Queries
	db *sql.DB
}

// Store Constructor
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// executes queries function in db transaction & support rollback
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rollback err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// TODO: impl insert new habit with new skill
