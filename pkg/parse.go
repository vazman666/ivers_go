package pkg

/*Делаем запрос к САЙТУ иверса. Получаем максимум(сколько дадут) наших заказов
возвращаем их в виде указателя на большой json*/
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"ivers_api/models"
	"net/http"
	"net/url"
)

func Parse(token string) (models.Parse, error) {
	type Par struct {
		Page int `json:"page"`
		Mode string `json:"mode"`
	}
	type m struct {
		Method string `json:"method"`
		Id     int    `json:"id"`
		Params []struct {
			Page int `json:"page"`
			Mode string `json:"mode"`
		} `json:"params"`
	}

	message := m{Method: "dataAction", Id: 1}
	message.Params = append(message.Params, Par{Page:0, Mode: "products"})
	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		return models.Parse{}, fmt.Errorf("Ошибка Marshal %v", err)
	}
	u, err := url.Parse("https://order.ivers.ru/history-order/service")
	if err != nil {
		return models.Parse{}, fmt.Errorf("Ошибка парсинга сайта %v", err)
	}
	reader := bytes.NewReader(bytesRepresentation)
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, u.String(), reader) // URL-encoded payload
	r.Header.Add("Accept", `application/json`)
	r.Header.Add("Cookie", token)
	resp, err := client.Do(r)
	if err != nil {
		return models.Parse{}, fmt.Errorf("Ошибка парсинга сайта %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		return models.Parse{}, fmt.Errorf("Ошибка авторизации парсинга сайта %v", err)
	}
	targets := models.Parse{}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &targets)
	if err != nil {
		return models.Parse{}, fmt.Errorf("Ошибка unmarshal результатов парсинга сайта %v", err)
	}

	return targets, nil
}
