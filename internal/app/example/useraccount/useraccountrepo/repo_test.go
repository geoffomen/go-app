package useraccountrepo_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"example.com/internal/app/common/base/entity"
	"example.com/internal/app/common/base/vo"
	"example.com/internal/app/example/useraccount/useraccountrepo"
	"example.com/internal/app/example/useraccount/useraccountsrv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

type mylogger struct{}

func (ml *mylogger) Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (ml *mylogger) Errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (ml *mylogger) Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func openDb() *sql.DB {
	db, err := sql.Open("sqlite3", "/tmp/exam.sqlitedb")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func TestCreateTable(t *testing.T) {
	db := openDb()
	_, err := db.Exec(`CREATE table 'user_account' (
			'id' INTEGER PRIMARY KEY AUTOINCREMENT ,
			'account' varchar(50) NULL,
			'password' varchar(100) NULL, 
			'salt' varchar(100) NULL,
			'status' tinyint(3) NULL,
			'name' varchar(50) NULL, 
			'avatar' varchar(500) NULL,
			'phone' varchar(50) NULL,
			'created_time' DATE NULL,
			'created_by' INTEGER  NULL,
			'updated_time' DATE NULL,
			'updated_by' INTEGER,
			'deleted_time' DATE NULL,
			'deleted_by' INTEGER NULL
		)`,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func TestRemoveDb(t *testing.T) {
	os.Remove("/tmp/exam.sqlitedb")
}

func TestCreate(t *testing.T) {
	db := openDb()
	repo := useraccountrepo.New(db, &mylogger{})
	id, err := repo.Create(vo.SessionInfo{Ctx: context.Background()},
		useraccountsrv.UseraccountEntity{
			Account:  "account1",
			Password: "123456",
			Salt:     "",
			Status:   1,
			Name:     "user1",
			Avatar:   "avatar1",
			Phone:    "phone1",
			BaseEntity: entity.BaseEntity{
				CreatedTime: time.Now(),
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	assert.NotEqual(t, id, 0)
}

func TestSelectById(t *testing.T) {
	db := openDb()
	repo := useraccountrepo.New(db, &mylogger{})
	entity, err := repo.SelectById(vo.SessionInfo{Ctx: context.Background()}, 1)
	assert.Nil(t, err)
	assert.NotEqual(t, entity.Id, 0)
}

func TestUpdateById(t *testing.T) {
	db := openDb()
	repo := useraccountrepo.New(db, &mylogger{})
	entity, _ := repo.SelectById(vo.SessionInfo{Ctx: context.Background()}, 1)
	entity.Name = "updatedName"
	err := repo.UpdateById(vo.SessionInfo{Uid: 100000000, Ctx: context.Background()}, *entity)
	assert.Nil(t, err)
	ne, _ := repo.SelectById(vo.SessionInfo{Ctx: context.Background()}, 1)
	assert.Equal(t, "updatedName", ne.Name)
	assert.Equal(t, 100000000, ne.UpdatedBy)
	assert.Equal(t, ne.Id, 1)
}

func TestLogicalDeleteById(t *testing.T) {
	db := openDb()
	repo := useraccountrepo.New(db, &mylogger{})
	err := repo.LogicalDeleteById(vo.SessionInfo{Uid: 10000000, Ctx: context.Background()}, 1)
	assert.Nil(t, err)
	ne, _ := repo.SelectById(vo.SessionInfo{Ctx: context.Background()}, 1)
	assert.Equal(t, 10000000, ne.DeletedBy)
	assert.Equal(t, ne.Id, 1)
}

func TestSelectPage(t *testing.T) {
	db := openDb()
	repo := useraccountrepo.New(db, &mylogger{})
	es, total, err := repo.SelectPage(vo.SessionInfo{Ctx: context.Background()}, []string{"1=1", fmt.Sprintf("deleted_time < '%s'", time.UnixMilli(0))}, "", 0, 100)
	assert.NotEqual(t, 0, total)
	assert.Nil(t, err)
	assert.Less(t, 0, len(es))
}
