/*
// Copyright 2018 Sendhil Vel. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

base_website.go
Date 		: 19/06/2018
Comment 	: This is seed file for creating any go website.
Version 	: 1.0.9
by Sendhil Vel
*/

package main

/*
	imports necessary pages
*/
import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
)

/*
	Defining necessary variables
*/
var tpl *template.Template
var err error
var r *gin.Engine
var rpcurl string

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}

/*
	jsonresponse - This function return the response in a json format
*/
func jsonresponse(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "process executed successfully",
	})
}

/*
	performLogin - gets the posted values for varibles username and password and check if the username/password combination is valid
*/
func performLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username, " : ", password)
}

/*
	initializeRoutes - This will defines various routes and relavant information
*/
func initializeRoutes(port string) {
	r.GET("/test", showHomePage)
	r.GET("/", showHomePage)
	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/", showHomePage)
		userRoutes.GET("/login", showLoginPage)
		userRoutes.POST("/login", performLogin)
		userRoutes.GET("/jsonresponse", jsonresponse)
	}
	fmt.Println("-------Starting server-------------")
}

/*
	Main - main function of the file
*/
func main() {

	gotenv.Load()
	port := os.Getenv("WEBSITE_PORT")

	/*Starting Server*/
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	r = gin.Default()
	r.LoadHTMLGlob("templetes/html/*")
	r.Static("/css", "templetes/css")
	r.Static("/js", "templetes/js")
	r.Static("/img", "templetes/img")
	r.Static("/fonts", "templetes/fonts")
	go initializeRoutes(port)
	fmt.Println("Web Portal is running on " + port)
	r.Run(port)
	fmt.Println("-------Started server-------------")
}

/*
	showHomePage - this will display status of website
*/
func showHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"Server": "Cool you are ready to start website in goLang",
	})
}

/*
	showLoginPage - This will load and show login page with necessary parameters
*/
func showLoginPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Login",
	}, "login.html")
}
