package main

import (
	"github.com/RakhimovAns/CreditCardValidator/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/", CheckCard)
	r.Run()
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func CheckCard(c *gin.Context) {
	var card models.Card
	if err := c.ShouldBindJSON(&card); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	if Check(card.Number) == true {
		c.JSON(http.StatusOK, gin.H{"message": "Your card is true"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "Your card is false"})
}

// Check this function check card by Luna's algorithm
func Check(number string) bool {
	var sum int = 0
	for i := 0; i < len(number); i++ {
		var digit int
		digit = int(number[i] - '0')
		if i%2 == 0 {
			digit *= 2
		}
		if digit > 10 {
			digit = digit%10 + digit/10
		}
		sum += digit
	}
	return sum%10 == 0
}
