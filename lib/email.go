package lib

import (
	"bytes"
	"log"
	"text/template"

	"github.com/faizalom/go-web/config"

	mail "github.com/xhit/go-simple-mail/v2"
)

func Mail(email *mail.Email) {
	server := mail.NewSMTPClient()
	server.Host = config.MailStruct.Host
	server.Port = config.MailStruct.Port
	server.Username = config.MailStruct.Username
	server.Password = config.MailStruct.Password
	server.Encryption = config.MailStruct.Encryption
	smtpClient, err := server.Connect()
	if err != nil {
		log.Println(err)
	}

	err = email.Send(smtpClient)
	if err != nil {
		log.Println(err)
	}
}

func SendRegisterMail(data map[string]any, to ...string) *mail.Email {
	email := mail.NewMSG()
	email.SetFrom(config.MailStruct.From)
	email.AddTo(to...)
	email.SetSubject("Please complete you registration process")

	var tpl bytes.Buffer
	t := template.Must(template.New("registration.html").ParseFiles(config.ThemePath + "/email/registration.html"))
	err := t.Execute(&tpl, data)
	if err != nil {
		log.Println(err)
	}
	email.SetBody(mail.TextHTML, tpl.String())
	return email
}
