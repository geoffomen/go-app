package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/geoffomen/go-app/internal/pkg/config"
	"github.com/geoffomen/go-app/internal/pkg/database"
	"github.com/geoffomen/go-app/internal/pkg/database/gormimp"
	"github.com/geoffomen/go-app/internal/pkg/mylog"
	"github.com/geoffomen/go-app/internal/pkg/webfw"
)

var (
	branchName string
	commitId   string
	buildTime  string

	showVer = flag.Bool("v", false, "show version")
)

func main() {
	profile := flag.String("profile", "dev", "Environment profile, something similar to spring profile")
	flag.Parse()
	cf := config.New(*profile)

	if *showVer {
		fmt.Printf("%s: %s\t%s\n", branchName, commitId, buildTime)
		os.Exit(0)
	}

	mylog.New(cf)

	db, err := gormimp.NewGorm(gormimp.GormConfig{
		Dialect:     cf.GetStringOrDefault("database.dialect", ""),
		UserName:    cf.GetStringOrDefault("database.userName", ""),
		Password:    cf.GetStringOrDefault("database.password", ""),
		Host:        cf.GetStringOrDefault("database.host", "localhost"),
		Port:        cf.GetIntOrDefault("database.port", 3306),
		Db:          cf.GetStringOrDefault("database.db", "test"),
		OtherParams: cf.GetStringOrDefault("database.otherParams", ""),
	})
	if err != nil {
		mylog.Panicf("failed to initrialize config component, err: %v", err)
	}
	database.New(db)

	ws := webfw.New(cf)
	// ws.RegisterHandler(accountctl.Controller())
	ws.Start()
}
