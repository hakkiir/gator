package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/hakkiir/gator/internal/config"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.Args) == 0 {
		return errors.New("invalid arguments: username missing")
	}
	name := cmd.Args[0]
	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}

	err = config.SetUser(user.Name, *s.cfg)

	if err != nil {
		return err
	}

	fmt.Println("user has been set")
	return nil
}
