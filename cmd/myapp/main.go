package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/geoffomen/go-app/pkg/config"
	"github.com/geoffomen/go-app/pkg/database"
	"github.com/geoffomen/go-app/pkg/database/gormimp"
	"github.com/geoffomen/go-app/pkg/mylog"
	"github.com/geoffomen/go-app/pkg/webfw"
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

	mylog.New(mylog.Configuration{
		EnableConsole:     cf.GetBoolOrDefault("log.enableConsole", true),
		ConsoleJSONFormat: cf.GetBoolOrDefault("log.consoleJSONFormat", true),
		ConsoleLevel:      cf.GetStringOrDefault("log.consoleLevel", "debug"),
		EnableFile:        cf.GetBoolOrDefault("log.enableFile", true),
		FileJSONFormat:    cf.GetBoolOrDefault("log.fileJSONFormat", true),
		FileLevel:         cf.GetStringOrDefault("log.fileLevel", "info"),
		FileLocation:      cf.GetStringOrDefault("log.fileLocation", "/tmp/miis/back/info.log"),
		ErrFileLevel:      cf.GetStringOrDefault("log.errFileLevel", "error"),
		ErrFileLocation:   cf.GetStringOrDefault("log.errFileLocation", "/tmp/miis/back/err.log"),
	})

	db, err := gormimp.NewGorm(gormimp.GormConfig{
		Dialect:     cf.GetStringOrDefault("database.dialect", ""),
		UserName:    cf.GetStringOrDefault("database.userName", ""),
		Password:    cf.GetStringOrDefault("database.password", ""),
		Host:        cf.GetStringOrDefault("database.host", "localhost"),
		Port:        cf.GetIntOrDefault("database.port", 3306),
		Db:          cf.GetStringOrDefault("database.db", "test"),
		OtherParams: cf.GetStringOrDefault("database.otherParams", ""),
	}, mylog.GetInstance())
	if err != nil {
		panic(fmt.Sprintf("failed to initrialize config component, err: %v", err))
	}
	database.New(db)

	ws := webfw.New(webfw.Configuration{
		Profile: cf.GetStringOrDefault("profile", "test"),
		Port:    cf.GetStringOrDefault("server.port", "8080"),
	}, mylog.GetInstance())
	// ws.RegisterHandler(accountctl.Controller())
	ws.Start()
}
