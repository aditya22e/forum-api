package services

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

type EmailService struct {
	host     string
	port     int
	username string
	password string
}

func NewEmailService() *EmailService {
	portStr := os.Getenv("EMAIL_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		logrus.Fatal("Invalid EMAIL_PORT: ", err)
	}

	return &EmailService{
		host:     os.Getenv("EMAIL_HOST"),
		port:     port,
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
	if err := d.DialAndSend(m); err != nil {
		logrus.Error("Failed to send email: ", err)
		return err
	}
	return nil
}
