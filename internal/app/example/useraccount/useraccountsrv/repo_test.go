package useraccountsrv

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"ibingli.com/internal/app/example/useraccount/useraccountdm"
	"ibingli.com/internal/pkg/myDatabase"
	"ibingli.com/internal/pkg/myHttpServer"
	"ibingli.com/internal/pkg/myLog/buildinImp"
)

func openDb() myDatabase.Iface {
	db, err := sql.Open("sqlite3", fmt.Sprintf("/tmp/exam-%d.sqlitedb", time.Now().Unix()))
	_, err1 := db.Exec(`CREATE table 'user_account' (
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
	)`)
	if err != nil || err1 != nil {
		log.Fatal(err)
	}
	return db
}

func CreateRecord(db myDatabase.Iface) (int64, error) {
	repo := Repository{logger: buildinImp.New()}
	id, err := repo.Create(&myHttpServer.SessionInfo{Ctx: context.Background()},
		db,
		&useraccountdm.UseraccountEntity{
			Account:  "account1",
			Password: "123456",
			Salt:     "",
			Status:   1,
			Name:     "user1",
			Avatar:   "avatar1",
			Phone:    "phone1",
			BaseEntity: myDatabase.BaseEntity{
				CreatedTime: time.Now(),
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return id, err
}

func TestCreate(t *testing.T) {
	db := openDb()
	id, err := CreateRecord(db)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, id)
}

func TestSelectById(t *testing.T) {
	db := openDb()
	id, err := CreateRecord(db)
	assert.Nil(t, err)
	assert.NotEqual(t, id, 0)
	repo := Repository{logger: buildinImp.New()}
	entity, err := repo.SelectById(&myHttpServer.SessionInfo{Ctx: context.Background()}, db, 1)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, entity.Id)
}

func TestUpdateById(t *testing.T) {
	db := openDb()
	id, err := CreateRecord(db)
	assert.Nil(t, err)
	assert.NotEqual(t, id, 0)
	repo := Repository{logger: buildinImp.New()}
	entity, _ := repo.SelectById(&myHttpServer.SessionInfo{}, db, 1)
	entity.Name = "updatedName"
	entity.UpdatedBy = 100000000
	err = repo.UpdateById(&myHttpServer.SessionInfo{}, db, entity)
	assert.Nil(t, err)
	ne, _ := repo.SelectById(&myHttpServer.SessionInfo{Ctx: context.Background()}, db, 1)
	assert.Equal(t, "updatedName", ne.Name)
	assert.Equal(t, int64(100000000), ne.UpdatedBy)
	assert.Equal(t, int64(1), ne.Id)
}

func TestLogicalDeleteById(t *testing.T) {
	db := openDb()
	CreateRecord(db)
	repo := Repository{logger: buildinImp.New()}
	err := repo.LogicalDeleteById(&myHttpServer.SessionInfo{Uid: 10000000, Ctx: context.Background()}, db, 1)
	assert.Nil(t, err)
	ne, _ := repo.SelectById(&myHttpServer.SessionInfo{Ctx: context.Background()}, db, 1)
	assert.Equal(t, int64(10000000), ne.DeletedBy)
	assert.Equal(t, int64(1), ne.Id)
}

func TestSelectPage(t *testing.T) {
	db := openDb()
	CreateRecord(db)
	CreateRecord(db)
	CreateRecord(db)
	repo := Repository{logger: buildinImp.New()}
	filter := myDatabase.NewFilter().
		Select("created_time, id").
		Where("deleted_time <= ?", time.UnixMilli(0)).
		Where("1=1").
		Where("1=? OR (1!=?)", 1, 0).
		Where("id in (?)", []int{1, 2}).
		Where("status in (?, ?)", 0, 1).
		Where("account like ?", "%").
		Offset(0).
		Limit(10).
		IsTotal(true)
	es, total, err := repo.SelectPage(&myHttpServer.SessionInfo{Ctx: context.Background()}, db, filter)
	assert.NotEqual(t, 0, total)
	assert.Nil(t, err)
	assert.Less(t, 0, len(es))
}
