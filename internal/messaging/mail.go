package messaging

import (
	"log"
	"net/smtp"
)

func SendMail() error {
	// Authentication for Gmail
	auth := smtp.PlainAuth("", "conta.crodu3@gmail.com", "xxxx", "smtp.gmail.com")

	// Set the recipient
	to := []string{"soratto.andre@gmail.com"}

	// Create the message
	msg := []byte("To: soratto.andre@gmail.com\r\n" +
		"Subject: Test Email\r\n" +
		"\r\n" +
		"This is a test email sent from Go using Gmail SMTP.\r\n")

	// Send the email
	err := smtp.SendMail("smtp.gmail.com:587", auth, "conta.crodu3@gmail.com", to, msg)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
