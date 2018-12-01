package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

var ch chan string = make(chan string)
var arr []interface{}

type City struct {
	Id   string
	Name string
}

var a []City

func httpGet(c *City) {

	resp, err := http.Get("https://www.tianqiapi.com/api/?version=v1&city=" + c.Name + "&cityid=" + c.Id)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	var js interface{}
	if err := json.Unmarshal(data, &js); err != nil {
		fmt.Println(err)
	}
	arr = append(arr, js)
	ch <- c.Id
}
func Getwea() {
	r := gin.Default()
	a = append(a, City{
		Id:   "101190807",
		Name: "新沂",
	})
	a = append(a, City{
		Id:   "101190101",
		Name: "南京",
	})
	a = append(a, City{
		Id:   "101191101",
		Name: "常州",
	})
	a = append(a, City{
		Id:   "101190301",
		Name: "镇江",
	})
	a = append(a, City{
		Id:   "101190901",
		Name: "淮安",
	})
	a = append(a, City{
		Id:   "101191001",
		Name: "连云港",
	})
	a = append(a, City{
		Id:   "101121501",
		Name: "日照",
	})
	a = append(a, City{
		Id:   "101120901",
		Name: "临沂",
	})
	a = append(a, City{
		Id:   "101120602",
		Name: "青州",
	})
	a = append(a, City{
		Id:   "101090713",
		Name: "黄骅",
	})
	a = append(a, City{
		Id:   "101030100",
		Name: "天津",
	})
	a = append(a, City{
		Id:   "101090501",
		Name: "唐山",
	})
	a = append(a, City{
		Id:   "101091101",
		Name: "秦皇岛",
	})
	b := a[:]
	go func() {
		for i := 0; i < len(b); i++ {
			httpGet(&b[i])
		}
	}()

	for i := 0; i < len(b); i++ {
		<-ch
	}
	r.GET("/weather", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": "200",
			"data": arr,
		})
	})
	r.Run(":8090")

}
