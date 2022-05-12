package services

import (
	"log"
	"net/smtp"
	"os"
)

func sendNotification(message string) {
	from := os.Getenv("FROM")
	password := os.Getenv("PASSWORD")
	to := os.Getenv("TO")
	smtpHost := os.Getenv("SMTPHOST")
	smtpPort := os.Getenv("SMTPPORT")
	content := []byte("From: " + from + "\r\n" + "To: " + to + "\r\n" + "Subject: Alert Host Down!!!\r\n\r\n" + message + "\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, content)

	if err != nil {
		log.Fatal(err)
	}
}
