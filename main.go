package main

import (
	"bytes"
	"log"
	"net/smtp"
	"text/template"
)

type Email struct {
	From    string
	To      string
	Subject string
	Body    string
}

func main() {
	email := Email{
		From:    "from@example.com",
		To:      "to@example.com",
		Subject: "Test Email",
		Body:    "This is a test email",
	}

	auth := smtp.PlainAuth("", email.From, "$smtp_password", "smtp.gmail.com")
	err := sendEmail(email, auth)
	if err != nil {
		log.Fatal(err)
	}
}

func sendEmail(email Email, auth smtp.Auth) error {
	tmpl, err := template.ParseFiles("example.html")
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, email)
	if err != nil {
		return err
	}

	msg := "From: " + email.From + "\r\n" +
		"To: " + email.To + "\r\n" +
		"Subject: " + email.Subject + "\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" + body.String()

	err = smtp.SendMail("smtp.gmail.com:587", auth, email.From, []string{email.To}, []byte(msg))
	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	log.Printf("Send email successfully to %s", email.To)

	return nil
}