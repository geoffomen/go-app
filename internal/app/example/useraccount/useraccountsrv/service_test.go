package useraccountsrv

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"ibingli.com/internal/app/example/useraccount/useraccountdm"
	"ibingli.com/internal/pkg/myConfig/mapImp"
	"ibingli.com/internal/pkg/myHttpServer"
	"ibingli.com/internal/pkg/myLog/buildinImp"
)

func setup() *Service {
	config, _ := mapImp.New()
	logger := buildinImp.New()
	db, _ := sql.Open("sqlite3", fmt.Sprintf("/tmp/exam-%d.sqlitedb", time.Now().Unix()))
	_, _ = db.Exec(`CREATE table 'user_account' (
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
	srv := New(config, logger, db, nil)
	return srv
}

func TestService(t *testing.T) {
	srv := setup()
	args := useraccountdm.CreateRequestDto{
		Account:  "account1",
		Password: "123456",
	}
	id, err := srv.Register(&myHttpServer.SessionInfo{}, &args)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, 0, id)
}
