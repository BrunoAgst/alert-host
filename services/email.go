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
	smtHost := os.Getenv("SMTPHOST")
	smtPort := os.Getenv("SMTPPORT")
	content := []byte("From: " + from + "\r\n" + "To: " + to + "\r\n" + "Subject: Alert Host Down!!!\r\n\r\n" + message + "\r\n")

	auth := smtp.PlainAuth("", from, password, smtHost)

	err := smtp.SendMail(smtHost+":"+smtPort, auth, from, []string{to}, content)

	if err != nil {
		log.Fatal(err)
	}
}
