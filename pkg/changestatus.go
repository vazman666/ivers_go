package pkg

import (
	"database/sql"
	"fmt"
	"ivers_api/models"
)

func ChangeStatus(db *sql.DB, f []models.Sql) error {

	for _, i := range f {
		//fmt.Printf("Update status %v\n", i)
		_, err := db.Exec("UPDATE ivers SET status=true WHERE id=? and subid=?", i.Id, i.SubId)
		if err != nil {
			return fmt.Errorf("Ошибка запроса UPDATE %v", err)
		}

	}
	return nil
}
