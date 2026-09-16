package achievement

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

type Store interface {
	/*
		Award returns true if the achievement was newly added.
		An already-earned achievement returns false with no error.
	*/
	Award(ctx context.Context, userID, achievementID string) (bool, error)
	List(ctx context.Context, userID string) ([]string, error)
}

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{store: store}
}

const (
	ActionWorkoutCompleted  = "WORKOUT_COMPLETED"
	AchievementFirstWorkout = "FIRST_WORKOUT"
)

var (
	ErrInvalidUserID = errors.New("user ID must not be empty")
	ErrInvalidAction = errors.New("unsupported action")
)

func (s *Service) RecordAction(ctx context.Context, userID, action string) ([]string, error) {
	if strings.TrimSpace(userID) == "" {
		return nil, ErrInvalidUserID
	}
	if action != ActionWorkoutCompleted {
		return nil, ErrInvalidAction
	}

	added, err := s.store.Award(ctx, userID, AchievementFirstWorkout)
	if err != nil {
		return nil, fmt.Errorf("award first workout: %w", err)
	}

	if !added {
		return nil, nil
	}

	return []string{AchievementFirstWorkout}, nil
}
