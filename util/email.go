package util

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(to string, subject string, html string) bool {

	password := os.Getenv("EMAIL_PASSWORD")
	auth := smtp.PlainAuth(
		"",
		"playatanu@gmail.com",
		password,
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + subject + "\n" + headers + "\n\n" + html
	fmt.Println("test 2")
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"playatanu@gmail.com",
		[]string{to},
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}

}
