package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
SERVICES
https://fixer.io/
https://currencylayer.com/
https://exchangeratesapi.io/
https://currencyapi.com/
https://www.abstractapi.com/
https://currency.getgeoapi.com/
https://openexchangerates.org/
https://www.exchangerate-api.com/
*/

func getFixer() string {
	apiKey := "YOUR_API_KEY"
	//url := "http://data.fixer.io/api/latest?access_key=" + apiKey
	url := fmt.Sprintf("http://data.fixer.io/api/latest?access_key=%s", apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(respBody)
}

func main() {
	fmt.Println("Parser")
	cur := getFixer()
	fmt.Println(cur)
}
