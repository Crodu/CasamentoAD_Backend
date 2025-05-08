package http

import (
	"net/http"

	"github.com/Crodu/CasamentoBackend/internal/config"
	"github.com/Crodu/CasamentoBackend/internal/database"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.Default()

	db := database.NewDBConnection()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.Use(func(c *gin.Context) {
		config, err := config.LoadConfig(".")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load config"})
			return
		}
		c.Set("config", config)
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/users", GetAllUsers)
	r.GET("/users/:id", GetUserByID)
	r.POST("/users", CreateUser)
	r.GET("/gifts", GetAllGifts)
	r.GET("/gifts/:id", GetGiftByID)
	r.POST("/gifts", CreateGift)
	r.POST("/login", Login)
	r.GET("/guests", GetAllGuests)
	r.GET("/guests/:id", GetGuestByID)
	r.POST("/guests", CreateGuest)
	r.POST("/ordergift", GenerateGiftPayment)
	r.POST("/confirmpayment", ConfirmPayment)
	// r.GET("/paymenttest", func(c *gin.Context) {
	// 	response, err := payments.GeneratePayment()
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate payment"})
	// 		return
	// 	}
	// 	var parsed map[string]interface{}
	// 	if err := json.Unmarshal(response, &parsed); err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse payment response"})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, parsed)
	// })
	// r.GET("/sendmail", func(c *gin.Context) {
	// 	err := messaging.SendMail()
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
	// })

	r.Run() // Default listens on :8080
}
