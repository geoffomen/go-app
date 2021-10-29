package main

import (
	"flag"
	"fmt"
	"testing"

	"github.com/geoffomen/go-app/examples/account/accountimp"
	"github.com/geoffomen/go-app/examples/user/userimp"
	"github.com/geoffomen/go-app/internal/pkg/config"
	"github.com/geoffomen/go-app/internal/pkg/database"
	"github.com/geoffomen/go-app/internal/pkg/database/gormimp"
	"github.com/geoffomen/go-app/internal/pkg/httpclient"
	"github.com/geoffomen/go-app/internal/pkg/mylog"
)

func TestApi(t *testing.T) {

	profile := flag.String("profile", "example", "Environment profile, something similar to spring profile")
	flag.Parse()
	cf := config.New(*profile)

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
	database.GetClient().GetStmt().AutoMigrate(accountimp.AccountEntity{})
	database.GetClient().GetStmt().AutoMigrate(accountimp.LoginTokenEntity{})
	database.GetClient().GetStmt().AutoMigrate(userimp.UserEntity{})

	type TestStruct struct {
		Methon  string
		Url     string
		Headers map[string]string
		Content interface{}
	}
	tests := []TestStruct{
		{"POST", "http://localhost:8080/exam/v1/account/register", nil, map[string]string{"account": "account1", "password": "123456"}},
		{"POST", "http://localhost:8080/exam/v1/account/login", nil, map[string]string{"account": "account1", "password": "123456"}},
		{"GET", "http://localhost:8080/exam/v1/user/info", map[string]string{"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJpc3N1ZUF0IjoiMjAyMS0xMC0yMVQxNDoxNjo0MC42ODAwMzk2NzQrMDg6MDAiLCJ1aWQiOjF9.jR92dCc6EMGp4vmgZFjTXydKsKX2ykrYMQ6n8YBD_7I"}, map[string]string{"id": "1"}},
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
