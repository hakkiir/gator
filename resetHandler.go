package main

import (
	"context"
	"fmt"
	"log"
)

func handlerReset(s *state, cmd command) error {

	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("users table reseted")
	return nil
}
