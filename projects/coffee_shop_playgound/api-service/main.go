package main

import (
	"net/http"
	"strconv"

	// mdtos "dev.azure.com/mdtchrf/CRHFMLifeInfrastructure/clctl.git/os"
	"github.com/gin-gonic/gin"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// pkg.go.dev/github.com/gin-gonic/gin#readme-api-examples
	r := gin.Default()
	r.GET("/clctl/:cmd/*action", func(c *gin.Context) {
		cmd := c.Param("cmd")
		action := c.Param("action")
		message := cmd + " called to " + action
		c.String(http.StatusOK, message)
	})

	/*r.GET("/clctl/test", func(c *gin.Context) {
		mdtos.LinuxInstallSessionManagerPlugin()
	})*/

	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	r.GET("/register", func(c *gin.Context) {
		name := c.Query("name")
		age, err := strconv.Atoi(c.Query("age"))
		if err != nil {
			panic(err)
		}

		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run()
}
