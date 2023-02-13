package connect

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	user     = "root"
	password = "Jose123_+#"
	host     = "127.0.0.1"
	port     = "3306"
	database = "loja_feliz"
)
var StringConnect = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)

func Connect() *sql.DB {

	db, err := sql.Open("mysql", StringConnect)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
