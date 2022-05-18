package main

import (
	"fmt"
	"ivers_api/pkg"
	"log"
)

//"bytes"

func main() {
	token, err := pkg.Login() //логинимся, получаем token
	if err != nil {
		log.Fatal(err)
	}
	c, err := pkg.Parse(token) //парсим сайт, получаем список последних заказов
	if err != nil {
		log.Fatal(err)
	}

	db, err := pkg.SqlLogin() //Открываем базу
	if err != nil {
		log.Fatal(err)
	}
	err = pkg.AddNew(c, db) //Загружаем в базу новые и изменённые
	if err != nil {
		log.Fatal(err)
	}
	unshipped, err := pkg.Unshipped(db) //получаем из базы все неотгруженные в []models.Sql
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Unshipped %v", unshipped)
	shippedunshipped, err := pkg.Shippedunshipped(c, unshipped, token) //Получаем отгруженные на сайте, но неотгруженные в базе
	if err != nil {
		log.Fatal(err)
	}
	if len(shippedunshipped) != 0 {
		fmt.Printf("Пауза 10 мин\n")
		//time.Sleep(10 * time.Minute)
		shippedunshipped, err := pkg.Shippedunshipped(c, unshipped, token)
		if err != nil {
			log.Fatal(err)
		}
		pkg.Exel(shippedunshipped)             //выгружаем отгруженные в exel
		pkg.ChangeStatus(db, shippedunshipped) //меняем статус у вновь отгруженных в базе

	}
	//fmt.Printf("рез=%v\n", shippedunshipped)
	err = pkg.ClearOld(db)
	if err != nil {
		log.Fatal(err)
	}

	pkg.SqlLogout(db)
	err = pkg.Logout(token) //отлогиниваемся
	if err != nil {
		log.Fatal(err)
	}

}

/*запускать каждый час
читаем по апи заказы за два последних дня
проверяем каждый из них, если заказа нет в нашей базе - парсим сайт чтобы найти примечание и добавляем
заказ в базу со статусом не отгружено

Для каждого неотгруженного из базы проверяем по апи, не изменился ли с татус на выполнен
Если изменился - ждём 10 минут
Для каждого неотгруженного из базы проверяем по апи, не изменился ли статус на выполнен
Если изменился - закидываем его в накладную и меняем его в базе на отгружено
отправляем накладную на почту. Удаляем в базе заказы старше 2 месяца

пробегаем по всем Неотгруженным (разделить могут только их)заказам в базе, сверяем количество.
Если изменилось - то заказ разделили, меняем в базе func separate




Количество брать из Статуса*/
