package gormimp

import (
	"fmt"
	"os"
	"testing"

	"github.com/storm-5/go-app/pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestDatabase(t *testing.T) {
	type Product struct {
		ID    int
		Code  string
		Price int
	}
	defer func() {
		os.Remove("/tmp/sqlite.db")
	}()
	db, _ := NewGorm(GormConfig{
		Dialect: "sqlite",
		Db:      "/tmp/sqlite.db",
	}, nil)
	// create table accoding to entity
	db.GetStmt().AutoMigrate(Product{})

	// Create
	db.GetStmt().Create(&Product{Code: "D1", Price: 100})
	db.GetStmt().Create(&Product{Code: "D1", Price: 200})
	db.GetStmt().Create(&Product{Code: "D1", Price: 500})
	db.GetStmt().Create(&Product{Code: "D2", Price: 300})

	results2 := make([]Product, 0)
	s1 := db.GetStmt().Where("code = ?", "D1").Where(db.GetStmt().Where("price = ?", 100).Or("price = ?", 200))
	s2 := db.GetStmt().Where("code = ?", "D2").Where("price = ?", 300)
	s3 := db.GetStmt().Model(Product{})
	s3.Where(
		s1,
	).Or(
		s2,
	).Find(&results2)
	assert.Equal(t, 3, len(results2))
	// Read
	product := Product{}
	var total int64
	var id int
	db.GetStmt().Model(Product{}).
		Select("id").
		Where("code = ?", "D1").
		Offset(1).
		Limit(1).
		Count(&total).
		Find(&product)
	assert.Equal(t, 3, int(total))
	assert.Equal(t, "", product.Code)
	assert.Equal(t, 0, product.Price)
	assert.NotNil(t, product.ID)
	id = product.ID
	product = Product{}
	db.GetStmt().Model(Product{}).Where("id = ?", id).First(&product)
	assert.Equal(t, "D1", product.Code)
	assert.Equal(t, 200, product.Price)
	assert.NotNil(t, product.ID)

	// Update product's price to 200
	db.GetStmt().Model(Product{}).Where("id = ?", id).Update("Price", 500)
	product = Product{}
	db.GetStmt().Model(Product{}).Where("id = ?", id).First(&product)
	assert.Equal(t, "D1", product.Code)
	assert.Equal(t, 500, product.Price)
	assert.NotNil(t, product.ID)
	// Update mutiple fields
	db.GetStmt().Model(Product{}).Where("id = ?", id).Updates(Product{Price: 0, Code: "F82"}) // only update non-zero value
	product = Product{}
	db.GetStmt().Model(Product{}).Where("id = ?", id).First(&product)
	assert.Equal(t, "F82", product.Code)
	assert.Equal(t, 500, product.Price)
	assert.NotNil(t, product.ID)
	db.GetStmt().Model(Product{}).Where("code = ?", "F82").Updates(map[string]interface{}{"Price": 0, "Code": "F62"})
	product = Product{}
	db.GetStmt().Model(Product{}).Where("id = ?", id).First(&product)
	assert.Equal(t, "F62", product.Code)
	assert.Equal(t, 0, product.Price)
	assert.NotNil(t, product.ID)
	// Delete
	db.GetStmt().Model(Product{}).Where("code = ?", "F62").Delete(Product{})
	db.GetStmt().Model(Product{}).Count(&total)
	assert.Equal(t, 3, int(total))
	db.GetStmt().Model(Product{}).Where("1 = 1").Delete(Product{})
	db.GetStmt().Model(Product{}).Count(&total)
	assert.Equal(t, 0, int(total))

	// a fail transaction
	err := db.GetStmt().DoTransaction(func(txcl *database.Client) error {
		stmt1 := txcl.GetStmt()
		err := stmt1.Create(&Product{Code: "T1", Price: 1000})
		if err != nil {
			return err
		}

		results := make([]Product, 0)
		stmt2 := txcl.GetStmt()
		stmt2.Model(Product{})
		stmt2.Where("code = ?", "T1")
		stmt2.Where("price = ?", 1000)
		err = stmt2.Find(&results)
		if err != nil {
			return err
		}
		assert.Equal(t, 1, len(results))
		txcl.GetStmt().DoTransaction(func(txcl2 *database.Client) error {
			err := txcl2.GetStmt().Create(&Product{Code: "T1", Price: 1100})
			if err != nil {
				return err
			}
			return nil
		})
		txcl.GetStmt().Model(Product{}).Count(&total)
		assert.Equal(t, 2, int(total))

		txcl.GetStmt().Create(&Product{Code: "T1", Price: 2000})
		txcl.GetStmt().Model(Product{}).Count(&total)
		assert.Equal(t, 3, int(total))

		return fmt.Errorf("fake error")
	})
	if err != nil {
		results := make([]Product, 0)
		stmt2 := db.GetStmt()
		stmt2.Model(Product{})
		stmt2.Where("code = ?", "T1")
		stmt2.Find(&results)
		assert.Equal(t, 0, len(results))
		fmt.Println("事务失败", results)
	}
	// a success transaction
	db.GetStmt().DoTransaction(func(txcl *database.Client) error {
		stmt1 := txcl.GetStmt()
		err := stmt1.Create(&Product{Code: "T1", Price: 1000})
		if err != nil {
			return err
		}

		results := make([]Product, 0)
		err = txcl.GetStmt().Model(Product{}).Where("code = ?", "T1").Find(&results)
		if err != nil {
			return err
		}
		assert.Equal(t, 1, len(results))

		// inner fail
		txcl.GetStmt().DoTransaction(func(txcl2 *database.Client) error {
			err := txcl2.GetStmt().Create(&Product{Code: "T1", Price: 1100})
			if err != nil {
				return err
			}
			txcl2.GetStmt().DoTransaction(func(txcl3 *database.Client) error {
				err := txcl3.GetStmt().Create(&Product{Code: "T1", Price: 1110})
				if err != nil {
					return err
				}
				results := make([]Product, 0)
				txcl3.GetStmt().Model(Product{}).Where("code = ?", "T1").Find(&results)
				assert.Equal(t, 3, len(results))
				fmt.Println(results)
				return nil
			})
			results := make([]Product, 0)
			txcl2.GetStmt().Model(Product{}).Where("code = ?", "T1").Find(&results)
			assert.Equal(t, 3, len(results))
			fmt.Println(results)
			return fmt.Errorf("inner failed")
		})
		results = make([]Product, 0)
		txcl.GetStmt().Model(Product{}).Where("code = ?", "T1").Find(&results)
		assert.Equal(t, 1, len(results))

		txcl.GetStmt().DoTransaction(func(txcl2 *database.Client) error {
			err := txcl2.GetStmt().Create(&Product{Code: "T1", Price: 1200})
			if err != nil {
				return err
			}
			txcl2.GetStmt().DoTransaction(func(txcl3 *database.Client) error {
				err := txcl3.GetStmt().Create(&Product{Code: "T1", Price: 1210})
				if err != nil {
					return err
				}
				return nil
			})
			return nil
		})

		// more operation
		err = txcl.GetStmt().Create(&Product{Code: "T1", Price: 2000})
		if err != nil {
			return err
		}

		return nil
	})
	results := make([]Product, 0)
	stmt2 := db.GetStmt()
	stmt2.Model(Product{})
	stmt2.Where("code = ?", "T1")
	stmt2.Count(&total).Find(&results)
	assert.Equal(t, 4, len(results))
	fmt.Println("事务成功", results)
}
