package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
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
	err := godotenv.Load("configs/api.env")
	checkErr(err)
	apiKey := os.Getenv("FIXER_API_KEY")
	url := fmt.Sprintf("http://data.fixer.io/api/latest?access_key=%s&symbols=EUR,GBP,CAD,AUD,CHF", apiKey)
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

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	fmt.Println("Parser")
	cur := getFixer()
	fmt.Println(cur)
}
