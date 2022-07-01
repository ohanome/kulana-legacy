package email

import (
	"fmt"
	"kulana/l"
	"kulana/setup"
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

func CheckMailEnvironment(dieOnError bool) bool {
	err := false
	if os.Getenv("SMTP_HOST") == "" {
		fmt.Println("Missing SMTP_HOST config. Edit the environment file under " + setup.GetEnvFile() + " and try again.")
		err = true
	}
	if os.Getenv("SMTP_USERNAME") == "" {
		fmt.Println("Missing SMTP_USERNAME config. Edit the environment file under " + setup.GetEnvFile() + " and try again.")
		err = true
	}
	if os.Getenv("SMTP_PASSWORD") == "" {
		fmt.Println("Missing SMTP_PASSWORD config. Edit the environment file under " + setup.GetEnvFile() + " and try again.")
		err = true
	}
	if os.Getenv("SMTP_PORT") == "" {
		fmt.Println("Missing SMTP_PORT config. Edit the environment file under " + setup.GetEnvFile() + " and try again.")
		err = true
	}
	if os.Getenv("SMTP_ADDRESS") == "" {
		fmt.Println("Missing SMTP_ADDRESS config. Edit the environment file under " + setup.GetEnvFile() + " and try again.")
		err = true
	}

	if err {
		if dieOnError {
			l.Emergency("Mail setup incomplete.")
		}

		return false
	}

	return true
}

func SendMail(to []string, subject string, message string) {
	CheckMailEnvironment(true)

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

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

func SendHostDownEmail(to string, host string, status int) {
	subject := fmt.Sprintf("[kulana] Host %s is DOWN!", host)
	message := fmt.Sprintf("The host %s is not reachable. The last request returned with the HTTP status %d.", host, status)
	SendMail([]string{to}, subject, message)
}

func SendHostUpEmail(to string, host string, status int) {
	subject := fmt.Sprintf("[kulana] Host %s is UP!", host)
	message := fmt.Sprintf("The host %s is reachable. The last request returned with the HTTP status %d.", host, status)
	SendMail([]string{to}, subject, message)
}

func SendNotificationMail(to string, host string, status int) {
	if status < 400 {
		SendHostUpEmail(to, host, status)
	} else {
		SendHostDownEmail(to, host, status)
	}
}
