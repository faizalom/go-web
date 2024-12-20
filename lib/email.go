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
	server.Host = config.EMail.Host
	server.Port = config.EMail.Port
	server.Username = config.EMail.Username
	server.Password = config.EMail.Password
	server.Encryption = config.EMail.Encryption
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
	email.SetFrom(config.EMail.From)
	email.AddTo(to...)
	email.SetSubject("Please complete you registration process")

	var tpl bytes.Buffer
	t := template.Must(template.New("email_registration.html").ParseGlob(config.Path.Theme))
	err := t.Execute(&tpl, data)
	if err != nil {
		log.Println(err)
	}
	email.SetBody(mail.TextHTML, tpl.String())
	return email
}
