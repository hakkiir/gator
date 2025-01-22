package main

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/hakkiir/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := int32(2)
	if len(cmd.Args) > 0 {
		if numeric := regexp.MustCompile(`\d`).MatchString(cmd.Args[0]); numeric {
			lmt, err := strconv.Atoi(cmd.Args[0])
			if err == nil {
				limit = int32(lmt)
			}
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.Description)
		fmt.Println(post.PublishedAt)
	}
	return nil
}
