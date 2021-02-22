package main

import "github.com/google/uuid"

// Owner model
type Owner struct {
	tableName struct{} `pg:"owner"`
	ID        uuid.UUID
	Name      string
	Token     string
}

// Template model
type Template struct {
	tableName struct{} `pg:"template"`
	ID        int
	OwnerID   uuid.UUID
	Name      string
	HTMLBody  string
	TextBody  string

	Owner Owner `pg:"rel:has-one"`
}

// Mailer model
type Mailer struct {
	tableName struct{} `pg:"mailer"`
	ID        int
	OwnerID   uuid.UUID
	Name      string
	Token     string
	Host      string

	Owner Owner `pg:"rel:has-one"`
}

// Message model
type Message struct {
	tableName       struct{} `pg:"message"`
	ID              int
	OwnerID         uuid.UUID
	TemplateID      int
	DefaultMailerID int
	Name            string
	Subject         string
	MailFrom        string
	SenderName      string

	Owner    Owner    `pg:"rel:has-one"`
	Template Template `pg:"rel:has-one"`
	Mailer   Mailer   `pg:"rel:has-one"`
}
