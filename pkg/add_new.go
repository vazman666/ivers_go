package pkg

import (
	"database/sql"
	"fmt"
	"ivers_api/models"
	"strconv"
)

/*Функция получает данные,спарсенные с сайта(новые)
проверяет, есть ли они в базе и не было ли разделения. Заносит в базу все новые с состоянием неотгруженно и все разделения
с состоянием неотгруженно	*/

func AddNew(parse models.Parse, db *sql.DB) error {
	for _, i := range parse.Result.Result.Products {
		for k, j := range i.Statuses {
			//fmt.Printf("Деталь %v id=%v Количество %v подномер %v статус %v\n", i.BaseNumber, i.Entity.Identificator, j.Quantity, k, j.Name)
			rows, err := db.Query("select status,quantity from ivers where id = ? and subid = ?", i.Entity.Identificator, k)
			if err != nil {
				return fmt.Errorf("Ошибка Select при работе с базой %v", err)
			}
			var stat, finds bool
			var quantity string
			for rows.Next() {
				err := rows.Scan(&stat, &quantity)
				if err != nil {
					fmt.Println(err)
					continue
				}
				//fmt.Printf("Нашлось в базе id %v subid %v quantity %v\n", i.Entity.Identificator, k, quantity)
				finds = true
			}
			defer rows.Close()
			//fmt.Printf("%v в базе %v\n", i.BaseNumber, finds)
			tmp := strconv.Itoa(j.Quantity)
			if finds {

				if tmp != quantity {
					//fmt.Printf("Нужно сменить количество\n")
					_, err = db.Exec("UPDATE ivers SET quantity=? WHERE id=?", tmp, i.Entity.Identificator)
					if err != nil {
						return fmt.Errorf("Ошибка UPDATE при работе с базой %v", err)
					}

				}

			} else {
				//fmt.Printf("Нужна добавить в базу \n")
				_, err = db.Exec("INSERT INTO ivers VALUES(?, ?, ?, ?, ?, ?, ?, ?, ? ,CURRENT_DATE())",
					i.Entity.Identificator, k, i.BaseNumber, i.Name, i.Brand, i.Price, tmp, i.ClientComment, false)
				if err != nil {
					return fmt.Errorf("Ошибка INSERT при работе с базой %v", err)
				}
			}
		}
	}
	return nil

}
