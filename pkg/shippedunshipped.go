package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ivers_api/models"
	"net/http"
	"net/url"
	"strconv"
)

func Shippedunshipped(site models.Parse, unshipped []models.Sql, token string) ([]models.Sql, error) {
	var rez []models.Sql
	for _, i := range unshipped {
		tmp, err := SearchProducts(token, i)
		if err != nil {
			return nil, err
		}
		if tmp {
			rez = append(rez, i)
		}
	}
	return rez, nil
}
func SearchProducts(token string, str models.Sql) (bool, error) {
	u, err := url.Parse("https://order.ivers.ru/api/v1/history-order/search-products")
	//u, err := url.Parse("http://localhost:8080")
	if err != nil {
		return false, fmt.Errorf("Ошибка запроса Rest API search-products %v", err)
	}
	q := u.Query()
	//for _, i := range id {
	//q.Set("uuids_products[]", i)
	q.Add("uuids_products[]", str.Id)
	//}
	u.RawQuery = q.Encode()
	//reader := bytes.NewReader(bytesRepresentation)
	//fmt.Printf("%v строка запроса  %v\n", u.RawQuery, u.String())
	client := &http.Client{}
	//fmt.Printf("==%v\n\n", u.String())
	r, _ := http.NewRequest(http.MethodPost, u.String(), nil) // URL-encoded payload
	r.Header.Add("Accept", `application/json`)
	r.Header.Add("Cookie", token)
	resp, err := client.Do(r)
	if err != nil {
		return false, fmt.Errorf("Ошибка запроса Rest API search-products %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return false, fmt.Errorf("Ошибка при авторизации.. Неверный логин/пароль? %v", err)
	}
	//fmt.Println(resp.Status)
	targets := models.SearchProducts{}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &targets)
	if err != nil {
		return false, fmt.Errorf("Ошибка Unmarshal запоса REST API %v", err)
	}
	if !targets.Success {
		return false, fmt.Errorf("Ошибка при авторизации.. Неверный логин/пароль? %v", err)
	}
	if !targets.Result.Params.Success {
		return false, fmt.Errorf("Ошибка при авторизации.. Неверный логин/пароль? %v", err)
	}
	//fmt.Printf("Статус = %v\n", targets)
	tmp, _ := strconv.Atoi(str.SubId)
	if targets.Result.Params.HistoryProducts[0].Statuses[tmp].Name == "Выполнен" {
		return true, nil
	}
	return false, nil
}
