package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Ads  string `json:"name"`
	Pubk string `json:"pubk"`
	Prik string `json:"prik"`
}

func main() {
	r := gin.Default()
	r.Static("/statics", "./static")
	r.LoadHTMLGlob("htmls/*")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "login",
		})
	})

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "home",
		})
	})

	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{
			"title": "upload",
		})
	})

	r.GET("/download", func(c *gin.Context) {
		c.HTML(http.StatusOK, "download.html", gin.H{
			"title": "download",
		})
	})

	r.GET("/details", func(c *gin.Context) {
		c.HTML(http.StatusOK, "details.html", gin.H{
			"title": "details",
		})
	})

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sign_up.html", gin.H{
			"title": "sign_up",
		})
	})

	r.POST("/signup_1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sign_up.html", gin.H{
			"title": "sign_up",
		})

		pass := c.PostForm("password")

		u := &User{
			Ads:  pass,
			Pubk: "2",
			Prik: "3",
		}

		c.JSON(200, u)
	})

	r.POST("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "home",
		})
	})

	r.Run()

}
