package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

func saveHistory(cur string) {
	fmt.Println(cur)
}

func backgroundUpdate() {
	for {
		cur := getFixerCurrency()
		fmt.Println(cur)
		time.Sleep(time.Second * 60)
	}
}

func getFixerCurrency() string {
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
	go backgroundUpdate()

	//cur := getFixerCurrency()
	//fmt.Println(cur)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	r.Run(":8080")
}
