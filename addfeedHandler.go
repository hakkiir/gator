package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hakkiir/gator/internal/database"
)

func handlerAddfeed(s *state, cmd command, user database.User) error {

	if len(cmd.Args) < 2 {
		return errors.New("not enough arguments: need name and url")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})

	s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}
