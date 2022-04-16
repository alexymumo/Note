package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://nba-stats4.p.rapidapi.com/players/?page=1&per_page=1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Host", "nba-stats4.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "275bf84d57mshb22f17d763ad964p18544fjsn2aa0008d7dc5")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
