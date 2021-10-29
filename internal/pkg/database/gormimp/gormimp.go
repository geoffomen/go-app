package gormimp

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/geoffomen/go-app/internal/pkg/database"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormDb struct {
	db              *gorm.DB
	isInTransaction bool
}

// Config ..
type GormConfig struct {
	Dialect            string
	Host               string
	Port               int
	Db                 string
	UserName           string
	Password           string
	OtherParams        string
	MaxIdleConns       int
	MaxOpenConns       int
	ConnMaxLifetimeSec int
}

func NewGorm(dbConfig GormConfig) (*GormDb, error) {
	switch dbConfig.Dialect {
	case "mysql":
		return newMysql(dbConfig)
	case "sqlite":
		return newSqlite(dbConfig)
	default:
		return nil, fmt.Errorf("unknow dialect")
	}
}

func newSqlite(dbConfig GormConfig) (*GormDb, error) {
	db, err := gorm.Open(sqlite.Open(dbConfig.Db), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // 慢 SQL 阈值
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // 彩色打印
			},
		),
	})
	if err != nil {
		return nil, fmt.Errorf("初始化数据库失败：%s", err)
	}
	return &GormDb{db: db}, nil
}

func newMysql(dbConfig GormConfig) (*GormDb, error) {
	var sb strings.Builder
	sb.WriteString(dbConfig.UserName)
	sb.WriteString(":")
	sb.WriteString(dbConfig.Password)
	sb.WriteString("@tcp(")
	sb.WriteString(dbConfig.Host)
	sb.WriteString(":")
	sb.WriteString(strconv.Itoa(dbConfig.Port))
	sb.WriteString(")/")
	sb.WriteString(dbConfig.Db)
	sb.WriteString("?")
	sb.WriteString(dbConfig.OtherParams)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       sb.String(),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  false,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: newGormLogger(),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to init database: %s", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to init database: %s", err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetimeSec))

	return &GormDb{db: db}, nil
}

func (srv *GormDb) GetStmt() database.Iface {
	return &GormDb{db: srv.db, isInTransaction: srv.isInTransaction}
}

func (srv *GormDb) AutoMigrate(dst ...interface{}) {
	srv.db.AutoMigrate(dst...)
}

// Table
func (srv *GormDb) Table(name string) database.Iface {
	tx := srv.db.Table(name)
	*srv = GormDb{db: tx}
	return srv
}

// Model
func (srv *GormDb) Model(value interface{}) database.Iface {
	ndb := srv.db.Model(value)
	*srv = GormDb{db: ndb}
	return srv
}

func (srv *GormDb) Where(query interface{}, args ...interface{}) database.Iface {
	vars := make([]interface{}, 0)
	for _, item := range args {
		switch item := item.(type) {
		case *GormDb:
			vars = append(vars, item.db)
		default:
			vars = append(vars, item)
		}
	}
	if v, ok := query.(*GormDb); ok {
		tx := srv.db.Where(v.db, vars...)
		*srv = GormDb{db: tx}
	} else {
		tx := srv.db.Where(query, vars...)
		*srv = GormDb{db: tx}
	}
	return srv
}

func (srv *GormDb) Not(query interface{}, args ...interface{}) database.Iface {
	if v, ok := query.(*GormDb); ok {
		tx := srv.db.Not(v.db, args...)
		*srv = GormDb{db: tx}
	} else {
		tx := srv.db.Not(query, args...)
		*srv = GormDb{db: tx}
	}
	return srv
}

func (srv *GormDb) Or(query interface{}, args ...interface{}) database.Iface {
	if v, ok := query.(*GormDb); ok {
		tx := srv.db.Or(v.db, args...)
		*srv = GormDb{db: tx}
	} else {
		tx := srv.db.Or(query, args...)
		*srv = GormDb{db: tx}
	}
	return srv
}
func (srv *GormDb) Select(query interface{}, args ...interface{}) database.Iface {
	tx := srv.db.Select(query, args...)
	*srv = GormDb{db: tx}
	return srv
}

func (srv *GormDb) Order(value interface{}) database.Iface {
	tx := srv.db.Order(value)
	*srv = GormDb{db: tx}
	return srv
}

func (srv *GormDb) Limit(value int) database.Iface {
	tx := srv.db.Limit(value)
	*srv = GormDb{db: tx}
	return srv
}

func (srv *GormDb) Offset(value int) database.Iface {
	tx := srv.db.Offset(value)
	*srv = GormDb{db: tx}
	return srv
}

// Group ...
func (srv *GormDb) Group(name string) database.Iface {
	tx := srv.db.Group(name)
	*srv = GormDb{db: tx}
	return srv
}

// Having ...
func (srv *GormDb) Having(query interface{}, args ...interface{}) database.Iface {
	tx := srv.db.Having(query, args...)
	*srv = GormDb{db: tx}
	return srv
}

func (srv *GormDb) Pluck(column string, value interface{}) error {
	tx := srv.db.Pluck(column, value)
	if tx.Error != nil {
		return fmt.Errorf("%s", tx.Error)
	}
	*srv = GormDb{db: tx}
	return nil
}

// Distinct ...
func (srv *GormDb) Distinct(args ...interface{}) database.Iface {
	tx := srv.db.Distinct(args...)
	*srv = GormDb{db: tx}
	return srv
}

func (srv *GormDb) Joins(query string, args ...interface{}) database.Iface {
	tx := srv.db.Joins(query, args...)
	*srv = GormDb{db: tx}
	return srv
}

func (srv *GormDb) Raw(sql string, values ...interface{}) database.Iface {
	tx := srv.db.Raw(sql, values...)
	*srv = GormDb{db: tx}
	return srv
}

func (srv *GormDb) Exec(sql string, values ...interface{}) database.Iface {
	tx := srv.db.Exec(sql, values...)
	*srv = GormDb{db: tx}
	return srv
}

func (srv *GormDb) Count(int64Ptr *int64) database.Iface {
	ndb := srv.db.Session(&gorm.Session{WithConditions: true})
	ndb.Offset(-1).Limit(-1).Count(int64Ptr)
	return srv
}

// Create ..
func (srv *GormDb) Create(entityPtr interface{}) error {
	rt := srv.db.Create(entityPtr)
	if rt.Error != nil {
		return fmt.Errorf("%s", rt.Error)
	}
	// *srv = gormClient{db: rt}
	return nil
}

func (srv *GormDb) Update(column string, value interface{}) error {
	rt := srv.db.Update(column, value)
	if rt.Error != nil {
		return fmt.Errorf("%s", rt.Error)
	}
	if rt.RowsAffected == 0 {
		return fmt.Errorf("update failed")
	}
	*srv = GormDb{db: rt}
	return nil
}

func (srv *GormDb) Updates(values interface{}) error {
	rt := srv.db.Updates(values)
	if rt.Error != nil {
		return fmt.Errorf("%s", rt.Error)
	}
	if rt.RowsAffected == 0 {
		return fmt.Errorf("update failed")
	}
	*srv = GormDb{db: rt}
	return nil
}

func (srv *GormDb) Save(value interface{}) error {
	rt := srv.db.Save(value)
	if rt.Error != nil {
		return fmt.Errorf("%s", rt.Error)
	}
	*srv = GormDb{db: rt}
	return nil
}

func (srv *GormDb) First(dest interface{}) error {
	rt := srv.db.First(dest)

	if rt.Error != nil {
		return fmt.Errorf("%s", rt.Error)
	}

	*srv = GormDb{db: rt}
	return nil
}

func (srv *GormDb) Last(dest interface{}) error {
	rt := srv.db.Last(dest)
	if rt.Error != nil {
		return fmt.Errorf("%s", rt.Error)
	}
	*srv = GormDb{db: rt}
	return nil
}

func (srv *GormDb) Scan(structPtr interface{}) error {
	rt := srv.db.Scan(structPtr)
	if rt.Error != nil {
		return fmt.Errorf("%s", rt.Error)
	}
	*srv = GormDb{db: rt}
	return nil
}

func (srv *GormDb) Find(dest interface{}) error {
	rt := srv.db.Find(dest)

	if rt.Error != nil {
		return fmt.Errorf("%s", rt.Error)
	}
	*srv = GormDb{db: rt}
	return nil

}

func (srv *GormDb) Delete(value interface{}) error {
	rt := srv.db.Delete(value)
	if rt.Error != nil {
		return fmt.Errorf("%s", rt.Error)
	}
	*srv = GormDb{db: rt}
	return nil
}

// DoTransaction ..
func (srv *GormDb) DoTransaction(f func(tx *database.Client) error) error {
	var err error
	if srv.isInTransaction {
		// nested transaction
		savePoint := fmt.Sprintf("sp_%p_%d", f, time.Now().UnixNano())
		srv.db.SavePoint(savePoint)
		defer func() {
			// Make sure to rollback when panic, Block error or Commit error
			if err != nil {
				srv.db.RollbackTo(savePoint)
			}
		}()

		err = f(&database.Client{Pool: &GormDb{db: srv.db.Session(&gorm.Session{WithConditions: true}), isInTransaction: true}})
		return err
	}
	tx := srv.db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	err = f(&database.Client{Pool: &GormDb{db: tx, isInTransaction: true}})
	if err == nil {
		err = tx.Commit().Error
	}
	return err

}
