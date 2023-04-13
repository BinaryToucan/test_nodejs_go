package main

import (
	"fmt"
	//"github.com/gin-gonic/gin"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"randomuser/db"
	u "randomuser/user"
)

const url = "https://randomuser.me/api/"

func main() {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	db.ConnectToDB()

	var response u.Response
	json.Unmarshal(body, &response)
	for _, p := range response.Results {
		db.InsertAll(p)
	}

	db.CloseConnect()
}
