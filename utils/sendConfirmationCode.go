package utils

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

// Sending code by email
func SendEmail(email string, code string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := 465
	senderEmail := os.Getenv("EMAIL_USER")
	senderPassword := "EMAIL_PASSWORD"

	message := gomail.NewMessage()
	message.SetHeader("From", senderEmail)
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Email Verification Code")
	message.SetBody("text/plain", fmt.Sprintf("Your email confirmation code: %s", code))

	dialer := gomail.NewDialer(smtpHost, smtpPort, senderEmail, senderPassword)
	dialer.SSL = true

	log.Println("Starting to send email...")
	if err := dialer.DialAndSend(message); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	log.Println("Email sent successfully")
	return nil
}
