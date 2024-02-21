package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users = []User{
	{ID: 1, Username: "Renan", Email: "renan@gmail.com"},
	{ID: 2, Username: "João", Email: "joao@gmail.com"},
	{ID: 3, Username: "Marcelo", Email: "marcelo@gmail.com"},
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Bem-vindo ao app monolito desenvolvido em Go!",
		})
	})

	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, user := range users {
			if id == fmt.Sprint(user.ID) {
				c.JSON(http.StatusOK, user)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado."})
	})

	router.Run("127.0.0.1:7000")
}
