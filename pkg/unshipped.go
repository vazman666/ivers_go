package pkg

import (
	"database/sql"
	"fmt"
	"ivers_api/models"
)

func Unshipped(db *sql.DB) ([]models.Sql, error) {

	var rez []models.Sql
	rows, err := db.Query("select * from ivers")
	if err != nil {
		return nil, fmt.Errorf("Ошибка Select при работе с базой %v", err)
	}

	var tmp models.Sql
	for rows.Next() {
		err := rows.Scan(&tmp.Id, &tmp.SubId, &tmp.Number, &tmp.Name, &tmp.Firm, &tmp.Price,
			&tmp.Quantity, &tmp.Remark, &tmp.Status, &tmp.Date)
		if err != nil {
			return nil, fmt.Errorf("Ошибка Scan при работе с базой %v", err)
		}
		//fmt.Printf("Скан с базы %v\n", tmp)
		if !tmp.Status {
			rez = append(rez, tmp)
		}

	}
	defer rows.Close()
	return rez, nil
}
