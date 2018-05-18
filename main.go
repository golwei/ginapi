package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/onLogin", func(c *gin.Context) {
		code := c.Query("code")
		//fmt.Println(code)
		//get open id
		appid := "wx5032f0d783147d67"
		secret := "bc9fc2d1a632ab230259dc98ea5ec116"
		rsp, _ := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appid, secret, code))
		defer rsp.Body.Close()
		body, _ := ioutil.ReadAll(rsp.Body)

		fmt.Println("onLogin->", code)
		fmt.Println("==========")
		fmt.Println(string(body))
		fmt.Println("==========")

	})
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello web")
	})

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("wcqt.bid"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(r, &m))
}
