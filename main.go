package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"

	"tinktur/interfaces/api/httphandler"
)

type User struct {
	ID int
}

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/tinktur")
	var user User

	e := echo.New()

	_ = e.GET("/ping", httphandler.Ping)

	err = db.QueryRow("SELECT * FROM users WHERE id = ?", 1).Scan(&user.ID)
	if err == nil {
		fmt.Println(user.ID)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
