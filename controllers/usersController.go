package controllers

import (
	"log"
	"net/http"

	"github.com/Hazem-Ben-Abdelhafidh/golangChat/config"
	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	client := config.ConnectDB()
	var requestBody RequestBody
	err := c.Bind(&requestBody)
	if err != nil {
		c.Error(err)
	}
	user, err := client.User.Create().SetName(requestBody.Name).SetPassword(requestBody.Password).Save(c)
	log.Println("USER IS : ", user)
	if err != nil {
		log.Println("Error! here ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failure",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"user":    user,
	})
}
