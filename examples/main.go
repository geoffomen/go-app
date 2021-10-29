package main

import (
	"flag"

	"github.com/geoffomen/go-app/examples/account/accountctl"
	"github.com/geoffomen/go-app/examples/hello/controller"
	"github.com/geoffomen/go-app/examples/user/userctl"
	"github.com/geoffomen/go-app/internal/pkg/config"
	"github.com/geoffomen/go-app/internal/pkg/database"
	"github.com/geoffomen/go-app/internal/pkg/database/gormimp"
	"github.com/geoffomen/go-app/internal/pkg/mylog"
	"github.com/geoffomen/go-app/internal/pkg/webfw"
)

func main() {
	profile := flag.String("profile", "example", "Environment profile, something similar to spring profiles")
	flag.Parse()
	cf := config.New(*profile)

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
	ws.RegisterHandler(controller.Controller())
	ws.RegisterHandler(accountctl.Controller())
	ws.RegisterHandler(userctl.Controller())
	ws.Start()
}
