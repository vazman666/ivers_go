package pkg

import (
	"database/sql"
	"fmt"
)

func ClearOld(db *sql.DB) error {

	_, err := db.Exec("DELETE FROM ivers WHERE `date` < DATE_SUB(NOW(), INTERVAL 1 MONTH)")
	if err != nil {
		return fmt.Errorf("Ошибка запроса DELETE %v", err)
	}
	return nil

}
