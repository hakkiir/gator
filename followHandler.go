package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hakkiir/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {

	if len(cmd.Args) == 0 {
		return errors.New("not enoughr arguments: url needed")
	}
	url := cmd.Args[0]

	feedId, err := s.db.FeedByURL(context.Background(), url)
	if err != nil {
		return err
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feedId.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println(feedFollow)
	return nil
}
