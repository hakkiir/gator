package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/hakkiir/gator/internal/config"
	"github.com/hakkiir/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {

	//read config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	//open Db connection
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	//save program state
	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}
	//register commands
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("register", handlerRegister)
	cmds.register("login", handlerLogin)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddfeed)
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", handlerFollow)
	cmds.register("following", handlerFollowing)
	//read cl arguments
	if len(os.Args) < 2 {
		log.Fatal("not enought arguments")
		return
	}

	cmdName, args := os.Args[1], os.Args[2:]

	//run command
	cmd := command{
		Name: cmdName,
		Args: args,
	}
	err = cmds.run(programState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
