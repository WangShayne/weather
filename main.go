package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/weather", func(c *gin.Context) {
		city := c.Query("city")
		cityId := c.Query("cityId")
		resp, err := http.Get("https://www.tianqiapi.com/api/?version=v1&city=" + city + "&cityid=" + cityId)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(body)
		var dat interface{}
		if err := json.Unmarshal(body, &dat); err != nil {
			log.Fatal(err)
		}
		fmt.Println(dat)
		//fmt.Printf("%v",dat)
		c.JSON(200, gin.H{
			"code":    0,
			"message": "success",
			"data":    dat,
		})

	})

	r.Run("localhost:8090")
}
