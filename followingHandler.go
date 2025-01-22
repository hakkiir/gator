package main

import (
	"context"
	"fmt"

	"github.com/hakkiir/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, ff := range feedFollows {
		fmt.Println(ff.FeedName)
	}
	return nil
}
