package main

import (
	"context"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	config "github.com/rewinfrey/go-example/config"
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
	if err := tx.QueryRowContext(context.Background(), "INSERT INTO users (name, namey, age) OUTPUT inserted.id VALUES (@p1, @p1, @p2)", "Rick", 37).Scan(&id); err != nil {
		panic(err)
	}
	fmt.Println(id)
	tx.Commit()

	// limit := 1

	// rows, err := db.DB().QueryContext(context.Background(), "SELECT id FROM users WHERE name = @p1 ORDER BY id ASC OFFSET 0 ROWS FETCH NEXT "+strconv.Itoa(limit)+" ROWS ONLY", "Rick")
	rows, err := db.DB().QueryContext(context.Background(), "SELECT id FROM users WHERE name = @p1 AND namey = @p1 ORDER BY id ASC OFFSET 0 ROWS FETCH FIRST @p2 ROWS ONLY", "Rick", 1)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			panic(err)
		}
		fmt.Println(id)
	}

	return

}
