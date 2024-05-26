package handlers

import (
	"github.com/RakhimovAns/CreditCardValidator/helper"
	"github.com/RakhimovAns/CreditCardValidator/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func CheckCard(c *gin.Context) {
	var card models.Card
	if err := c.ShouldBindJSON(&card); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}
	if helper.Check(card.Number) == true {
		c.JSON(http.StatusOK, gin.H{"message": "Your card is valid"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "Your card is not valid"})
}
