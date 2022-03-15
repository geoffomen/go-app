package main

import (
	"flag"
	"fmt"
	"testing"

	"github.com/geoffomen/go-app/pkg/config/viperimp"
	"github.com/geoffomen/go-app/pkg/database"
	"github.com/geoffomen/go-app/pkg/database/gormimp"
	"github.com/geoffomen/go-app/pkg/httpclient"
	"github.com/geoffomen/go-app/pkg/mylog"
	"github.com/geoffomen/go-app/pkg/mylog/zapimp"

	"github.com/geoffomen/go-app/examples/user/userimp"
	"github.com/geoffomen/go-app/examples/useracc/useraccimp"
)

func TestApi(t *testing.T) {

	profile := flag.String("profile", "example", "Environment profile, something similar to spring profile")
	flag.Parse()
	cf, _ := viperimp.New(*profile)

	zapimp.New(zapimp.Configuration{
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
	database.SetInstance(db)

	database.GetClient().GetStmt().AutoMigrate(useraccimp.AccountEntity{})
	database.GetClient().GetStmt().AutoMigrate(useraccimp.LoginTokenEntity{})
	database.GetClient().GetStmt().AutoMigrate(userimp.UserEntity{})

	type TestStruct struct {
		Methon  string
		Url     string
		Headers map[string]string
		Content interface{}
	}
	tests := []TestStruct{
		{"POST", "http://localhost:8000/exam/v1/useracc/register", nil, map[string]string{"account": "account1", "password": "123456"}},
		{"POST", "http://localhost:8000/exam/v1/useracc/login", nil, map[string]string{"account": "account1", "password": "123456"}},
		{"GET", "http://localhost:8000/exam/v1/user/info", map[string]string{"Authorization": ""}, map[string]string{"id": "1"}},
	}

	for _, testCase := range tests {
		switch testCase.Methon {
		case "GET":
			res, err := httpclient.Get(testCase.Url, testCase.Content.(map[string]string), testCase.Headers)
			if err != nil {
				t.Error(err) //Something is wrong while sending request
			}
			fmt.Printf("%s", res)
		case "POST":
			res, err := httpclient.PostJson(testCase.Url, testCase.Headers, testCase.Content.(map[string]string))
			if err != nil {
				t.Error(err) //Something is wrong while sending request
				fmt.Printf("%s", err)
			}
			fmt.Printf("%s", res)
		}
	}
}
