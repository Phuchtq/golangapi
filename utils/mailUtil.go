package utils

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"v3/constants/notis"
	"v3/spModels"

	"gopkg.in/mail.v2"
)

var orgEmail string = "your_email"
var securityPass string = "your_app_password"
var host string = "smtp.gmail.com"
var port int = 587

func SendMail(templatePath, subject string, model spModels.MailBody) error {
	var body bytes.Buffer
	if t, err := template.ParseFiles(templatePath); err != nil {
		log.Print(notis.MailMsg + "SendMail - " + err.Error())
		return errors.New(notis.InternalErr)
	} else {
		t.Execute(&body, model)
	}

	m := mail.NewMessage()
	m.SetHeader("From", orgEmail)
	m.SetHeader("To", model.Email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	diabler := mail.NewDialer(host, port, orgEmail, securityPass)

	if err := diabler.DialAndSend(m); err != nil {
		log.Print(notis.MailMsg + "SendMail - " + err.Error())
		return errors.New(notis.GenerateMailWarnMsg)
	}

	return nil
}
