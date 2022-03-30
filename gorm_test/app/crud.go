package app

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Result struct {
	Name  string
	Total int
}

func Crud(db *gorm.DB) {

	db.AutoMigrate(&User{})

	user := User{Name: "hannah", Age: 30, Birthday: time.Now()}

	//result := db.Create(&user)

	//fmt.Println(user.ID)
	//fmt.Println(result.Error)
	//fmt.Println(result.RowsAffected)

	var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	//db.CreateInBatches(users, 100)

	//SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	var result1 *User
	db.Model(&User{}).First(&result1)

	db.Find(&users, []int{1, 2, 3})

	db.Find(&users)

	user = User{}
	db.Where("name = ?", "jinzhu1").Find(&user)
	fmt.Println("user : ", user)

	users = []User{}
	db.Where("name In ?", []string{"jinzhu", "jinzhu2"}).Find(&users)
	fmt.Println("users : ", users)

	user = User{}
	db.First(&user, "id = ?", "1")

	users = []User{}
	db.Find(&users, User{Age: 20})

	users = []User{}
	db.Find(&users, map[string]interface{}{"age": 20})

	//SELECT * FROM `users` WHERE (`users`.`name` <> "jinzhu1" AND `users`.`name` <> "jinzhu2")
	users = []User{}
	db.Not([]User{{Name: "jinzhu1"}, {Name: "jinzhu2"}}).Find(&users)

	//SELECT * FROM `users` WHERE `name` NOT IN ("jinzhu1","jinzhu2")
	users = []User{}
	db.Not(map[string]interface{}{"name": []string{"jinzhu1", "jinzhu2"}}).Find(&users)

	//SELECT * FROM `users` WHERE name = 'jinzhu1' OR (`users`.`name` = "jinzhu2" AND `users`.`age` = 18)
	db.Where("name = 'jinzhu1'").Or(User{Name: "jinzhu2", Age: 18}).Find(&users)

	// SELECT `name`,`age` FROM `users`
	db.Select("name", "age").Find(&users)
	fmt.Println("users : ", &users)

	//SELECT COALESCE(age,42) FROM `users`
	rows, _ := db.Table("users").Select("MAX(age,?)", 42).Rows()
	defer rows.Close()
	for rows.Next() {
		fmt.Println("rows : ", rows)
	}
	// SELECT * FROM `users` ORDER BY age desc,name
	users = []User{}
	db.Order("age desc").Order("name").Find(&users)

	//SELECT name, sum(age) as total FROM `users` GROUP BY `name` HAVING sum(age) > -1
	var results []Result
	db.Model(&User{}).Select("name, sum(age) as total").
		Group("name").
		Having("sum(age) > -1").
		Find(&results)
	fmt.Println("results", results)

	//SELECT DISTINCT `name`,`age` FROM `users` ORDER BY name, age desc
	db.Model(&user).
		Distinct("name", "age").
		Order("name, age desc").
		Find(&results)
	fmt.Println(results)

	//SELECT users.name, products.code FROM `users` left join products on products.id = users.id
	var result2 map[string]interface{}
	db.Model(&User{}).
		Select("users.name, products.code").
		Joins("left join products on products.id = users.id").
		Scan(&result2)
	fmt.Println(result2)
}
