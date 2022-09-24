package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Store interface {
	Querier
	CreateHabitAndSkill(ctx context.Context, arg CreateHabitAndSkillTxParams) (CreateHabitAndSkillResult, error)
	CreateHabitLogTx(ctx context.Context, arg CreateHabitLogTxParams) (CreateHabitLogTxResult, error)
}

// Contain all functions to execute db queries and transaction
type SQLStore struct {
	*Queries
	db *sql.DB
}

// Store Constructor
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// executes queries function in db transaction & support rollback
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
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
	UserID                int64  `json:"user_id"`
	SkillTitle            string `json:"skill_title"`
	HabitTitle            string `json:"habit_title"`
	TargetConsecutiveDays int32  `json:"target_consecutive_days"`
}

type CreateHabitAndSkillResult struct {
	NewSkill Skill `json:"skill"`
	NewHabit Habit `json:"habit"`
}

func (store *SQLStore) CreateHabitAndSkill(ctx context.Context, arg CreateHabitAndSkillTxParams) (CreateHabitAndSkillResult, error) {
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
			Title:                 arg.HabitTitle,
			SkillID:               sql.NullInt64{Int64: result.NewSkill.SkillID, Valid: true},
			UserID:                arg.UserID,
			TargetConsecutiveDays: arg.TargetConsecutiveDays,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

type CreateHabitLogTxParams struct {
	UserID  int64 `json:"user_id"`
	HabitID int64 `json:"habit_id"`
}

type CreateHabitLogTxResult struct {
	NewHabitLog    HabitLog `json:"habit_log"`
	IsCreateFailed bool     `json:"is_create_failed"`
}

func (store *SQLStore) CreateHabitLogTx(ctx context.Context, arg CreateHabitLogTxParams) (CreateHabitLogTxResult, error) {
	var result CreateHabitLogTxResult
	now := time.Now()

	err := store.execTx(ctx, func(q *Queries) error {
		// Get latest habit log by user
		latestHabitLogs, err := q.GetLatestHabitLogByUser(ctx, GetLatestHabitLogByUserParams{
			UserID:  arg.UserID,
			HabitID: arg.HabitID,
		})
		if err != nil {
			return err
		}

		// Check can create habit logic
		var canCreateHabitLog bool
		if len(latestHabitLogs) == 0 {
			canCreateHabitLog = true
		} else {
			latestTime := latestHabitLogs[0].CreatedAt

			if now.Day() != latestTime.Day() ||
				now.Month() != latestTime.Month() ||
				now.Year() != latestTime.Year() {
				canCreateHabitLog = true
			}
		}

		// Insert new habit log if check pass
		if canCreateHabitLog {
			habitLog, err := q.CreateHabitLog(ctx, CreateHabitLogParams{
				UserID:    arg.UserID,
				HabitID:   arg.HabitID,
				CreatedAt: now,
			})
			if err != nil {
				return err
			}

			result.NewHabitLog = habitLog
		}

		result.IsCreateFailed = !canCreateHabitLog
		return nil
	})

	return result, err
}
