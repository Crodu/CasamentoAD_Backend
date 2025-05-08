package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/Crodu/CasamentoBackend/internal/config"
	"github.com/Crodu/CasamentoBackend/internal/models"
	"github.com/Crodu/CasamentoBackend/internal/payments"
	"github.com/gin-gonic/gin"
)

// Assuming db is a global variable for database connection

func GetAllUsers(c *gin.Context) {
	var users []models.User
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func Login(c *gin.Context) {
	// Implement login logic here
	token := "example_token"
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.UserInput
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// Ensure the ID is not set to allow auto-increment
	newUser := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}
	if err := db.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func GetAllGuests(c *gin.Context) {
	var guests []models.Guest
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Find(&guests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch guests"})
		return
	}
	c.JSON(http.StatusOK, guests)
}

func GetGuestByID(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	var guest models.Guest
	if err := db.First(&guest, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guest not found"})
		return
	}
	c.JSON(http.StatusOK, guest)
}

func CreateGuest(c *gin.Context) {
	var guest models.GuestInput
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newGuest := models.Guest{
		FirstName: guest.FirstName,
		LastName:  guest.LastName,
		Email:     guest.Email,
	}

	if err := db.Create(&newGuest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create guest"})
		return
	}
	c.JSON(http.StatusCreated, guest)
}

func GetAllGifts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Use a join to fetch gifts along with their bought status
	var giftResponses []models.GiftResponse
	if err := db.Table("gifts").
		Select(`gifts.id, gifts.name, gifts.description, gifts.price, gifts.link, guests.first_name AS bought_by`).
		Joins("LEFT JOIN bought_gifts ON gifts.id = bought_gifts.gift_id").
		Joins("LEFT JOIN guests ON bought_gifts.guest_id = guests.id").
		Scan(&giftResponses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch gifts"})
		return
	}

	c.JSON(http.StatusOK, giftResponses)
}

func GetGiftByID(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	var gift models.Gift
	if err := db.First(&gift, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gift not found"})
		return
	}
	c.JSON(http.StatusOK, gift)
}

func CreateGift(c *gin.Context) {
	var gift models.GiftInput
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&gift); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newGift := models.Gift{
		Name:        gift.Name,
		Description: gift.Description,
		Price:       gift.Price,
		Link:        gift.Link,
	}

	if err := db.Create(&newGift).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create gift"})
		return
	}
	c.JSON(http.StatusCreated, gift)
}

func GetAllBoughtGifts(c *gin.Context) {
	var boughtGifts []models.BoughtGift
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Find(&boughtGifts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bought gifts"})
		return
	}
	c.JSON(http.StatusOK, boughtGifts)
}

func GetBoughtGiftByID(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)
	var boughtGift models.BoughtGift
	if err := db.First(&boughtGift, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bought gift not found"})
		return
	}
	c.JSON(http.StatusOK, boughtGift)
}

func CreateBoughtGift(c *gin.Context) {
	var boughtGift models.BoughtGift
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&boughtGift); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := db.Create(&boughtGift).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bought gift"})
		return
	}
	c.JSON(http.StatusCreated, boughtGift)
}

type BuyGiftInput struct {
	GiftID        uint   `json:"gift_id"`
	GuestName     string `json:"guest_name"`
	GuestLastName string `json:"guest_last_name"`
	Email         string `json:"email"`
}

func GenerateGiftPayment(c *gin.Context) {
	var input BuyGiftInput
	mercadoPagoKey := c.MustGet("config").(config.Config).MercadoPagoKey
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Fetch the gift and guest from the database
	var gift models.Gift
	if err := db.First(&gift, input.GiftID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gift not found"})

		return
	}

	var guest models.Guest
	if err := db.Select("first_name, last_name, email").Where("email = ?", input.Email).First(&guest).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create a new guest if not found
			newGuest := models.Guest{
				FirstName: input.GuestName,
				LastName:  input.GuestLastName,
				Email:     input.Email,
			}
			if err := db.Create(&newGuest).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create guest"})
				return
			}
			// Fetch the newly created guest to ensure it's properly assigned
			if err := db.Where("email = ?", input.Email).First(&guest).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch newly created guest"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch guest"})
			return
		}
	}

	// Generate payment using Mercado Pago (or any other payment service)
	// Check for existing pending payment
	var existingPayment models.Payment
	err := db.Where("gift_id = ? AND guest_id = ? AND status = ?", gift.ID, guest.ID, "pending").First(&existingPayment).Error
	if err == nil {
		// Pending payment found, return its details
		c.JSON(http.StatusOK, gin.H{
			"message": "Existing pending payment found",
			"qrcode":  existingPayment.QRCode,
			"payment": gin.H{ // Reconstruct a minimal payment response
				"id": existingPayment.PaymentID,
				"transaction_details": gin.H{
					"external_resource_url": existingPayment.Link,
				},
			},
		})
		return
	} else if err != gorm.ErrRecordNotFound {
		// An error occurred other than not finding a record
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check for existing payments"})
		return
	}

	// No pending payment found, proceed to generate a new one
	response, err := payments.GeneratePayment(gift.Price, input.Email, guest.FirstName, guest.LastName, gift.Name, mercadoPagoKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate payment"})
		return
	}

	qrcode, err := payments.GetQRCode(response)
	if err := db.Create(&models.Payment{
		GuestID:   guest.ID,
		GiftID:    gift.ID,
		PaymentID: strconv.Itoa(response.ID),
		QRCode:    qrcode,
		Status:    "pending",
		Link:      response.TransactionDetails.ExternalResourceURL,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
		return
	}

	jsonPayment, err := json.Marshal(response)

	c.JSON(http.StatusOK, gin.H{
		"message": "Payment generated successfully",
		"qrcode":  qrcode,
		"payment": jsonPayment,
	})
}

// paymentdata = {
//   action: "payment.updated",
//   api_version: "v1",
//   data: {"id":"123456"},
//   date_created: "2021-11-01T02:02:02Z",
//   id: "123456",
//   live_mode: false,
//   type: "payment",
//   user_id: 278927631
// }

type PaymentWebhook struct {
	Action     string `json:"action"`
	APIVersion string `json:"api_version"`
	Data       struct {
		ID string `json:"id"`
	} `json:"data"`
	DateCreated string `json:"date_created"`
	ID          string `json:"id"`
	LiveMode    bool   `json:"live_mode"`
	Type        string `json:"type"`
	UserID      int    `json:"user_id"`
}

func ConfirmPayment(c *gin.Context) {
	var input PaymentWebhook
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if the payment ID exists in the database
	var payment models.Payment
	if err := db.First(&payment, "payment_id = ?", input.Data.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	// Update the payment status to confirmed
	if err := db.Model(&payment).Update("status", "confirmed").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to confirm payment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment confirmed successfully"})
}

func CancelPaymentIfTimeout(c *gin.Context) {
	var timeLimit = 30 // minutes
	var input models.Payment
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// Calculate the time threshold
	thresholdTime := time.Now().Add(-time.Duration(timeLimit) * time.Minute)

	// Update payments that are pending and older than the thresholdTime
	result := db.Model(&models.Payment{}).
		Where("status = ? AND created_at < ?", "pending", thresholdTime).
		Update("status", "canceled")

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel payments"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No pending payments found older than the time limit"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment canceled successfully"})
}
