package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hakkiir/gator/internal/config"
	"github.com/hakkiir/gator/internal/database"
)

func handlerRegister(s *state, cmd command) error {

	if len(cmd.Args) == 0 {
		return errors.New("invalid arguments: name missing")
	}

	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		return err
	}
	config.SetUser(name, *s.cfg)
	fmt.Println("User was created")
	fmt.Println(user)
	return nil
}
