package services

import (
	"os"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	host     string
	port     int
	username string
	password string
}

func NewEmailService() *EmailService {
	return &EmailService{
		host:     os.Getenv("EMAIL_HOST"),
		port:     os.Getenv("EMAIL_PORT"),
		username: os.Getenv("EMAIL_USER"),
		password: os.Getenv("EMAIL_PASSWORD"),
	}
}

func (s *EmailService) SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(s.host, s.port, s.username, s.password)
	return d.DialAndSend(m)
}
