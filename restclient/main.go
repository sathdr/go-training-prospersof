package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	url := "https://jsonplaceholder.typicode.com/users"
	method := "GET"

	client := new(http.Client)
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(body))

	var m []map[string]interface{}

	if err := json.Unmarshal(body, &m); err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%#v\n", m)

	for _, user := range m {
		fmt.Println(user["name"])
		// var address = user["address"]
		//fmt.Printf("%T\n", address)
	}
}
