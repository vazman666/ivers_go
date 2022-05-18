package pkg

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SqlLogin() (*sql.DB, error) {
	db, err := sql.Open("mysql", "vazman:rbhgbxb1@unix(/var/run/mysqld/mysqld.sock)/japautozap")

	if err != nil {
		return nil, fmt.Errorf("Ошибка открытия базы %v\n", err)
	}
	return db, nil
}
func SqlLogout(db *sql.DB) {
	db.Close()
}
