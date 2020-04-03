package main

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	config "github.com/rewinfrey/go-example/config"
	models "github.com/rewinfrey/go-example/internal/models"
)

func main() {
	db, err := config.OpenDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	var id int64
	tx, err3 := db.DB().BeginTx(context.Background(), nil)
	if err3 != nil {
		panic(err3)
	}
	if err := tx.QueryRowContext(context.Background(), "INSERT INTO users (name, namey, age) OUTPUT inserted.id VALUES (@p1, @p1, @p2)", "KingJames", 37).Scan(&id); err != nil {
		panic(err)
	}
	fmt.Println(id)
	tx.Commit()

	// limit := 1

	// rows, err := db.DB().QueryContext(context.Background(), "SELECT id FROM users WHERE name = @p1 ORDER BY id ASC OFFSET 0 ROWS FETCH NEXT "+strconv.Itoa(limit)+" ROWS ONLY", "KingJames")
	rows, err := db.DB().QueryContext(context.Background(), "SELECT id FROM users WHERE name = @p1 AND namey = @p1 ORDER BY id ASC OFFSET 0 ROWS FETCH FIRST @p2 ROWS ONLY", "KingJames", 1)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			panic(err)
		}
		fmt.Println(id)
	}

	var user1 models.User
	fmt.Println("first:")
	db.First(&user1)
	fmt.Println(user1)

	var user2 models.User
	fmt.Println("last:")
	db.Last(&user2)
	fmt.Println(user2)

	var user3 models.User
	fmt.Println("where with first:")
	db.Where("id = @p1", id).First(&user3)
	fmt.Println(user3)

	var user4 models.User
	fmt.Println("where with last:")
	db.Where("id = @p1", id).Last(&user4)
	fmt.Println(user4)

	var user5 models.User
	fmt.Println("take")
	db.Take(&user5) // ERROR: mssql: Invalid usage of the option NEXT in the FETCH statement.
	fmt.Println(user5)

	var users []models.User
	fmt.Println("find:")
	db.Find(&users)
	fmt.Println(users)

	var users2 []models.User
	fmt.Println("where with find:")
	db.Where("name = @p1", "KingJames").Find(&users2)
	fmt.Println(users2)

	var users3 []models.User
	fmt.Println("limit with where and find:")
	db.Limit(2).Where("name = @p1", "KingJames").Find(&users3) // Error: mssql: Invalid usage of the option NEXT in the FETCH statement.
	fmt.Println(users3)

	var users4 []models.User
	fmt.Println("scope with where:")
	db.Scopes(nameForScope).Where("namey = @p1", "KingJames").Find(&users4)
	fmt.Println(users4)

	var users5 []models.User
	fmt.Println("scope with where and limit:")
	db.Limit(2).Scopes(nameForScope).Where("namey = @p1", "KingJames").Find(&users5) // Error: mssql: Invalid usage of the option NEXT in the FETCH statement.
	fmt.Println(users5)

	var users6 []models.User
	fmt.Println("scope with where:")
	db.Scopes(nameForScope, fakeLimit).Where("namey = @p1", "KingJames").Find(&users6) // Error: mssql: Invalid usage of the option NEXT in the FETCH statement.
	fmt.Println(users6)

	// Because we cannot user a Scope for the order by / fetch dance, we then have to resort to `Raw` as the Gorm query interface:
	var users7 []models.User
	fmt.Println("raw with order / fetch:")
	db.Raw("SELECT * FROM users WHERE name = @p1 ORDER BY id DESC OFFSET 0 ROWS FETCH NEXT 4 ROWS ONLY", "KingJames").Find(&users7)
	fmt.Println(users7)

	// Another quirk about this is errors don't seem to propagate unless a follow up query is issued, so I'm using this placeholder as the last query issued to force any errors to be reported.
	var endUser models.User
	fmt.Println("first:")
	db.First(&endUser)
	fmt.Println(endUser)

	return
}

func nameForScope(db *gorm.DB) *gorm.DB {
	return db.Where("name = @p1", "KingJames")
}

func fakeLimit(db *gorm.DB) *gorm.DB {
	return db.Where("ORDER BY id DESC OFFSET 0 ROWS FETCH NEXT 10 ROWS ONLY")
}
