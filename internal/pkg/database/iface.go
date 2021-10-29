package database

// Iface ..
type Iface interface {
	// AutoMigrate create table
	AutoMigrate(entities ...interface{})

	// GetStmt ..
	GetStmt() Iface

	// Table specify table name
	Table(name string) Iface

	// Model specify orm object
	Model(entity interface{}) Iface

	// Where add condition
	// sub-query can nested in query，
	// client.GetStmt().Table("orders").Where("amount > (?)", client.GetStmt().Table("orders").Select("AVG(amount)")).Find(&orders)
	// result：SELECT * FROM "orders" WHERE amount > (SELECT AVG(amount) FROM "orders");
	// Complex SQL
	// client.GetStmt().Table("users").Where(
	// 		client.GetStmt().Where("pizza = ?", "pepperoni").Where(client.GetStmt().Where("size = ?", "small").Or("size = ?", "medium")),
	// ).Or(
	// 		client.GetStmt().Where("pizza = ?", "hawaiian").Where("size = ?", "xlarge"),
	// ).Find(&User{})
	// result：SELECT * FROM `users` WHERE (
	//		pizza = "pepperoni" AND (size = "small" OR size = "medium")
	//	) OR (
	//		pizza = "hawaiian" AND size = "xlarge")
	Where(query interface{}, args ...interface{}) Iface

	// Not
	Not(query interface{}, args ...interface{}) Iface

	// client.GetStmt().Table("users").Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2", Age: 18}).Find(&users)
	// result：SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);
	// Complex SQL
	// client.GetStmt().Table("users").Where(
	// 		client.GetStmt().Where("pizza = ?", "pepperoni").Where(client.GetStmt().Where("size = ?", "small").Or("size = ?", "medium")),
	// ).Or(
	// 		client.GetStmt().Where("pizza = ?", "hawaiian").Where("size = ?", "xlarge"),
	// ).Find(&User{})
	// result：SELECT * FROM `users` WHERE (
	//		pizza = "pepperoni" AND (size = "small" OR size = "medium")
	//	) OR (
	//		pizza = "hawaiian" AND size = "xlarge")
	Or(query interface{}, args ...interface{}) Iface

	// Select select column, defuault *
	Select(query interface{}, args ...interface{}) Iface

	// Order
	//client.GetStmt().Table("users").Order("age desc").Order("name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;
	Order(value interface{}) Iface

	// Limit
	Limit(value int) Iface

	// Offset
	Offset(value int) Iface

	// Group ...
	Group(name string) Iface

	// Having ...
	Having(query interface{}, args ...interface{}) Iface

	// Distinct ...
	Distinct(args ...interface{}) Iface

	// Pluck used to query single column from a model as a map
	//     var ages []int64
	//     client.GetStmt().Find(&users).Pluck("age", &ages)
	Pluck(column string, value interface{}) error

	// Joins ...
	// client.GetStmt().Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&result{})
	// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id
	// Complex SQL
	// client.GetStmt().Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").
	//		Joins("JOIN credit_cards ON credit_cards.user_id = users.id").
	//		Where("credit_cards.number = ?", "411111111111").Find(&user)
	Joins(query string, args ...interface{}) Iface

	// Raw
	Raw(sql string, values ...interface{}) Iface

	// Exec
	Exec(sql string, values ...interface{}) Iface

	// Count
	Count(int64Ptr *int64) Iface

	// Create ..
	Create(entityPtr interface{}) error

	// Update
	// client.GetStmt().Model(&user).Where("active = ?", true).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;
	Update(column string, value interface{}) error

	// Updates update multiple column
	// client.GetStmt().Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
	// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;
	//
	// client.GetStmt().Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
	// UPDATE users SET name='hello', age=18, actived=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
	//
	// client.GetStmt().Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
	// UPDATE users SET name='hello' WHERE id=111;
	//
	// client.GetStmt().Model(&product).Update("price", gorm.Expr("price * ? + ?", 2, 100))
	// UPDATE "products" SET "price" = price * 2 + 100, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 3;
	Updates(values interface{}) error

	// Save upsert
	Save(valuePtr interface{}) error

	// First
	First(destPtr interface{}) error

	// Last
	Last(destPtr interface{}) error

	// Scan
	Scan(structPtr interface{}) error

	// Find
	Find(destPtr interface{}) error

	// Delete ..
	//
	// client.GetStmt().Delete(&email)
	// DELETE from emails where id = 10;
	//
	// client.GetStmt().Where("name = ?", "jinzhu").Delete(&email)
	// DELETE from emails where id = 10 AND name = "jinzhu";
	// client.GetStmt().Where("email LIKE ?", "%jinzhu%").Delete(Email{})
	// DELETE from emails where email LIKE "%jinzhu%";
	Delete(value interface{}) error

	DoTransaction(f func(tx *Client) error) error
}

var (
	ins Iface
)

// New ...
func New(db Iface) {
	ins = db
}

type Client struct {
	Pool Iface
}

func GetClient() *Client {
	return &Client{Pool: ins}
}

func (cl Client) GetStmt() Iface {
	return cl.Pool.GetStmt()
}
