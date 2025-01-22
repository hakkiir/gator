package main

import (
	"context"
	"errors"

	"github.com/hakkiir/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {

	if len(cmd.Args) == 0 {
		return errors.New("not enoughr arguments: url needed")
	}
	url := cmd.Args[0]
	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		Name: user.Name,
		Url:  url,
	})
	if err != nil {
		return err
	}
	return nil
}
