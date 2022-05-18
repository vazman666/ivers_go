package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ivers_api/models"
	"log"
	"net/http"
	"net/url"
)

func Login() (string, error) {
	u, err := url.Parse("https://order.ivers.ru/api/v1/login")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("login", models.Login)
	q.Set("password", models.Pass)
	u.RawQuery = q.Encode()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, u.String(), nil) // URL-encoded payload
	r.Header.Add("Accept", `application/json`)
	resp, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		fmt.Printf("Ошибка при авторизации.. Неверный логин/пароль? \n")
		return "", fmt.Errorf("Ошибка при авторизации.. Неверный логин/пароль?")
	}
	//fmt.Println(resp.Status)
	targets := models.LoginType{}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &targets)
	if err != nil {
		log.Fatal(err)
	}
	if !targets.Success {
		fmt.Printf("Ошибка авторизации:Ошибка запроса\n")
		return "", fmt.Errorf("Ошибка авторизации:Ошибка запроса")
	}
	if !targets.Result.Params.Success {
		fmt.Printf("Ошибка авторизации:Неправильный логин/пароль?\n")
		return "", fmt.Errorf("Ошибка авторизации:Неправильный логин/пароль?")
	}
	return targets.Result.Params.Token, nil
}
func Logout(token string) error {
	u, err := url.Parse("https://order.ivers.ru/api/v1/logout")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	u.RawQuery = q.Encode()
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, u.String(), nil) // URL-encoded payload
	r.Header.Add("Accept", `application/json`)
	r.Header.Add("Cookie", token)
	resp, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		fmt.Printf("Ошибка при выходе.  \n")
		return fmt.Errorf("Ошибка при выходе")
	}
	//fmt.Println(resp.Status)
	targets := models.Logout{}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &targets)
	if err != nil {
		log.Fatal(err)
	}
	if !targets.Success {
		fmt.Printf("Ошибка выхода \n")
		return fmt.Errorf("Ошибка выхода")
	}
	return nil
}
