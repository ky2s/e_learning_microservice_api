package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type FormatUser struct {
	Status string `json:"status"`
	Data   User   `json:"data"`
}

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Profession string `json:"profession"`
	Avatar     string `json:"avatar"`
	Role       string `json:"role"`
}

func GetJson(url string, target interface{}) error {

	resp, err := http.Get(url)
	fmt.Println(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func GetUserByID(id int) (FormatUser, error) {

	urlUserService := os.Getenv("URL_SERVICE_USERS")
	url := urlUserService + "/users/" + strconv.Itoa(id) + "/detail"

	var user FormatUser

	err := GetJson(url, &user)

	if err != nil {

		return user, err

	}

	return user, nil

}
