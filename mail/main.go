package mail

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

func SendMail(to []string, subject string, message string) {

	// Sender data.
	from := "kulana <" + os.Getenv("SMTP_ADDRESS") + ">"
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	// smtp server configuration.
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	mail := Mail{
		Sender:  from,
		To:      to,
		Subject: subject,
		Body:    message,
	}

	// Message.
	msg := BuildMessage(mail)

	// Authentication.
	auth := smtp.PlainAuth("", username, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}

func SendTestMail() {
	to := []string{
		"thomasartmann@icloud.com",
	}

	subject := "Testmessage"
	message := "This is a test message."

	SendMail(to, subject, message)
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}
