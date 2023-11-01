package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"example.com/internal/app/example/echoargs/echoargsctl"
	"example.com/internal/app/example/useraccount/useraccountctl"
	"example.com/internal/pkg/httpserver/httpserverimp"
	"example.com/internal/pkg/myconfig/viperimp"
	"example.com/internal/pkg/mylog/zapimp"

	_ "github.com/mattn/go-sqlite3"
)

var (
	branchName string
	commitId   string
	buildTime  string
)

func main() {
	showVer := flag.Bool("v", false, "show version")
	profile := flag.String("profile", "example", "Environment profile, something similar to spring profiles")
	flag.Parse()
	if *showVer {
		fmt.Printf("%s: %s\t%s\n", branchName, commitId, buildTime)
		os.Exit(0)
	}
	fmt.Printf("using profile: %s\n", *profile)

	config, err := viperimp.New(*profile)
	if err != nil {
		log.Fatalf("读取配置信息失败, err: %s", err)
	}

	logger, err := zapimp.New(zapimp.Configuration{
		EnableConsole:     config.GetBoolOrDefault("log.enableConsole", true),
		ConsoleJSONFormat: config.GetBoolOrDefault("log.consoleJSONFormat", true),
		ConsoleLevel:      config.GetStringOrDefault("log.consoleLevel", "debug"),
		EnableFile:        config.GetBoolOrDefault("log.enableFile", true),
		FileJSONFormat:    config.GetBoolOrDefault("log.fileJSONFormat", true),
		FileLevel:         config.GetStringOrDefault("log.fileLevel", "info"),
		FileLocation:      config.GetStringOrDefault("log.fileLocation", "/tmp/goapp/example_info.log"),
		ErrFileLevel:      config.GetStringOrDefault("log.errFileLevel", "error"),
		ErrFileLocation:   config.GetStringOrDefault("log.errFileLocation", "/tmp/goapp/example_err.log"),
	})
	if err != nil {
		logger.Fatalf("初始化日志模块失败, err: %s", err)
	}

	// mysql
	// cfg := mysql.Config{
	// 	User:   config.GetStringOrDefault("database.user", ""), // os.Getenv("DBUSER"),
	// 	Passwd: config.GetStringOrDefault("database.user", ""), // os.Getenv("DBPASS"),
	// 	Net:    "tcp",
	// 	Addr:   fmt.Sprintf("%s:%s", config.GetStringOrDefault("database.host", ""), config.GetStringOrDefault("database.port", "")),
	// 	DBName: config.GetStringOrDefault("database.db", ""),
	// }
	// db, err := sql.Open("mysql", cfg.FormatDSN())
	// if err != nil {
	// 	logger.Fatalf("初始化数据库连接失败, err: %s", err)
	// }
	// pingErr := db.Ping()
	// if pingErr != nil {
	// 	logger.Fatalf("%s", pingErr)
	// }
	// logger.Infof("数据库已连接!")

	// sqlite
	dbPath := "/tmp/exam.sqlitedb"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		logger.Fatalf("%s", err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)
	defer func() {
		os.Remove(dbPath)
	}()
	pingErr := db.Ping()
	if pingErr != nil {
		logger.Fatalf("%s", pingErr)
	}
	logger.Infof("数据库已连接!")

	httpSrv := httpserverimp.New(&httpserverimp.Options{
		Port:   config.GetIntOrDefault("server.port", 0),
		Logger: logger,
	})
	httpSrv.AddRouter(useraccountctl.New(config, logger, db))
	httpSrv.AddRouter(echoargsctl.New(config, logger, db))
	logger.Infof("http server 端口: %d", config.GetIntOrDefault("server.port", 0))

	var wg sync.WaitGroup
	// 启动服务
	wg.Add(1)
	go func() {
		defer wg.Done()
		httpSrv.Listen()
	}()

	// 等候关闭信号
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	// 关闭服务
	httpSrv.Shutdown()
	// 等待服务完全关闭
	wg.Wait()
}
