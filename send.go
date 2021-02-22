package main

import (
	"bytes"
	"text/template"

	"github.com/gogearbox/gearbox"
	"github.com/google/uuid"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type request struct {
	Name          string                 `json:"name"`
	RecipientName string                 `json:"recipient_name"`
	MailTo        string                 `json:"mail_to"`
	Data          map[string]interface{} `json:"data"`
}

func parsedTemplate(toParse string, data map[string]interface{}) string {
	var tpl bytes.Buffer
	t, _ := template.New("").Parse(toParse)
	t.Execute(&tpl, data)
	return tpl.String()
}

func sendMessage(ctx gearbox.Context) {
	var req request
	ctx.ParseBody(&req)

	message := &Message{
		OwnerID: ctx.GetLocal("ownerId").(uuid.UUID),
		Name:    req.Name,
	}
	err := DB.Model(message).Relation("Template").Relation("Mailer").
		Where("message.name = ?name AND message.owner_id = ?owner_id").Select()
	if err != nil {
		panic(err)
	}

	m := mail.NewV3Mail()
	m.SetFrom(mail.NewEmail(message.SenderName, message.MailFrom))
	m.AddContent(mail.NewContent("text/plain", parsedTemplate(message.Template.TextBody, req.Data)))
	if message.Template.HTMLBody != "" {
		m.AddContent(mail.NewContent("text/html", parsedTemplate(message.Template.HTMLBody, req.Data)))
	}

	personalization := mail.NewPersonalization()
	personalization.AddTos(mail.NewEmail(req.RecipientName, req.MailTo))
	personalization.Subject = parsedTemplate(message.Subject, req.Data)

	m.AddPersonalizations(personalization)

	mailReq := sendgrid.GetRequest(message.Mailer.Token, "/v3/mail/send", message.Mailer.Host)
	mailReq.Method = "POST"
	mailReq.Body = mail.GetRequestBody(m)
	_, err = sendgrid.API(mailReq)

	if err != nil {
		ctx.Status(gearbox.StatusInternalServerError).SendJSON(&Response{"something went wrong"})
	} else {
		ctx.SendJSON(&Response{"sent"})
	}
}
