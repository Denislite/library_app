package mail

import (
	"github.com/Denislite/library_app/env"
	"net/smtp"
	"strconv"
)

type Email struct {
	Name string
	Pass string
	Host string
	Port string
}

func SendEmail(addr string, duty float64) error {
	EmailData := &Email{
		Name: env.Env.Email.Email,
		Pass: env.Env.Email.Pass,
		Host: env.Env.Email.Host,
		Port: env.Env.Email.Port,
	}
	address := EmailData.Host + ":" + EmailData.Port
	subject := "Subject: U have duty\n"
	body := "U have duty:" + strconv.FormatFloat(duty, 'E', -1, 32)
	message := []byte(subject + body)
	to := []string{addr}
	auth := smtp.PlainAuth("", EmailData.Name, EmailData.Pass, EmailData.Host)
	err := smtp.SendMail(address, auth, EmailData.Name, to, message)
	if err != nil {
		return err
	}
	return nil
}
