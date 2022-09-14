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

type CreateHabitAndSkillTxParams struct {
	UserID             int64  `json:"user_id"`
	SkillTitle         string `json:"skill_title"`
	HabitTitle         string `json:"habit_title"`
	MaxConsecutiveDays int32  `json:"max_consecutive_days"`
}

type CreateHabitAndSkillResult struct {
	NewSkill Skill `json:"skill"`
	NewHabit Habit `json:"habit"`
}

func (store *Store) CreateHabitAndSkill(ctx context.Context, arg CreateHabitAndSkillTxParams) (CreateHabitAndSkillResult, error) {
	var result CreateHabitAndSkillResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Create skill
		result.NewSkill, err = q.CreateSkill(ctx, CreateSkillParams{
			UserID: arg.UserID,
			Title:  arg.SkillTitle,
		})
		if err != nil {
			return err
		}

		// Create Habit
		result.NewHabit, err = q.CreateHabit(ctx, CreateHabitParams{
			Title:              arg.HabitTitle,
			SkillID:            sql.NullInt64{Int64: result.NewSkill.SkillID, Valid: true},
			UserID:             arg.UserID,
			MaxConsecutiveDays: arg.MaxConsecutiveDays,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
