package mail

import (
	"fmt"
	"net/smtp"
	"strconv"
)

const (
	EMAIL = ""
	PASS  = ""
	HOST  = ""
	PORT  = ""
)

func SendEmail(addr string, duty float64) error {
	address := HOST + ":" + PORT
	subject := "Subject: U have duty\n"
	body := "U have duty:" + strconv.FormatFloat(duty, 'E', -1, 32)
	message := []byte(subject + body)
	to := []string{addr}
	auth := smtp.PlainAuth("", EMAIL, PASS, HOST)
	err := smtp.SendMail(address, auth, EMAIL, to, message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
