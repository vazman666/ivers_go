package models

var Login = "29771"
var Pass = "23y1sm"

type LoginType struct {
	Success bool   `json:"success"`
	Id      string `json:"id"`
	Errors  string `json:"errors"`
	Result  struct {
		Content string `json:"content"`
		Params  struct {
			Success bool   `json:"success"`
			Token   string `json:"token"`
		} `json:"params"`
	} `json:"result"`
}

type Logout struct {
	Success bool `json:"success"`
}
type SearchProducts struct {
	Success bool `json:"success"`
	Result  struct {
		Params struct {
			Success         bool `json="success"`
			HistoryProducts []struct {
				Uuid         string  `json="uuid"`
				OrderUuid    string  `json=orderUuid"`
				CartItemId   string  `json="cart_item_id"`
				Brand        string  `json="brand"`
				BaseNumber   string  `json="baseNumber"`
				Name         string  `json="name"`
				Info         string  `json="info"`
				ExtraNumber  string  `json="extraNumber"`
				Price        float32 `json="price"`
				PricesString string  `json="pricesString"`
				Quantity     int     `json="quantity"`
				Statuses     []struct {
					Name         string `json="name"`
					Description  string `json="description"`
					TimeDelivery string `json="timeDelivery"`
					Quantity     int    `json="quantity"`
					Date         string `json="date"`
				} `json="statuses"`
			} `json'"historyProducts"`
		} `json="params"`
	} `json="result"`
}
type Parse struct {
	Result struct {
		Result struct {
			Products []struct {
				//Sku        string `json="sku"`
				Name       string `json="name"`
				BaseNumber string `json="baseNumber"`
				Quantity   string `json="quantity"`
				Price      string `json="price"`
				Brand      string `json="brand"`
				Statuses   []struct {
					Name     string `json="name"`
					Quantity int    `json="quantity"`
				} `json="statuses"`
				Entity struct {
					Identificator string `json="identificator"`
				} `json='"entity"`
				ClientComment string `json="clientComment"`
			} `json="products"`
		} `json="result"`
	} `json="result"`
}
type Sql struct {
	Id       string
	SubId    string
	Number   string
	Name     string
	Firm     string
	Price    string
	Quantity string
	Remark   string
	Status   bool
	Date     string
}

//MySQL > DELETE FROM `mytable` WHERE `date` < DATE_SUB(NOW(), INTERVAL 1 MONTH);  удалить все записи старше одного месяца

//INSERT INTO table(data) VALUES(NOW(. ;   добавить текущую дату в таблицу
//Тип столбца DATE

//Структура таблицы japautozap/ivers
// id string длиной 36(37 с запасом)
// subid int
// number string (15) номер детали
// name string(20) название детали
// firm string(15) фирма
// price string (10) цена
// quantity string (10) количество
// remark string (20) примечание
// status bool   отгружено/нет
// date DATE время записи заказа

// создаём таблицу
// mysql -u vazman -p
// create database japautozap;  - если ещё не создана
// use japautozap
/* create table ivers (id varchar (37),
subid TINYINT UNSIGNED,
number varchar(15),
name varchar(20),
firm varchar(15),
price varchar(10),
quantity varchar(10),
remark varchar(20),
status bool,
date DATE);*/
//ALTER TABLE ivers modify column name varchar(50);   изменяем столбец
//ALTER TABLE ivers modify column number varchar(30);
//mysqldump -u vazman -p japautozap ivers > japautozap_ivers.sql
//mysql -uvazman -prbhgbxb1 japautozap < japautozap_ivers.sql
// update ivers set status=false;
//ALTER TABLE ivers ADD COLUMN quanity2  VARCHAR(10) AFTER status;
//ALTER TABLE ivers ADD COLUMN status2  bool  AFTER quanity2;
//ALTER TABLE ivers DROP COLUMN status2;
//ALTER TABLE ivers ADD COLUMN subid  TINYINT UNSIGNED AFTER id;
//update ivers set status=false where number="C-110";
