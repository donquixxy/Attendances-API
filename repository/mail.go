package repository

import (
	"log"

	"gopkg.in/gomail.v2"
)

type MailRepostory interface {
	SendEmailCheckin(receiverEmail string) error
	SendEmailCheckout(receiverEmail string) error
}

type mailRepository struct {
	gm     *gomail.Message
	dialer *gomail.Dialer
}

func NewMailRepository(
	gm *gomail.Message,
	dialer *gomail.Dialer,
) MailRepostory {
	return &mailRepository{
		gm:     gm,
		dialer: dialer,
	}
}

// ("smtp.gmail.com", 587, "agusariputra70@gmail.com", "testpassword!23")

func (m *mailRepository) SendEmailCheckin(receiverEmail string) error {

	m.gm.SetHeader("From", "agusariputra70@gmail.com")
	m.gm.SetHeader("To", receiverEmail)
	m.gm.SetHeader("Subject", "Daily reminder to checkin")
	m.gm.SetBody("text/plain", "Remember to checkin")

	err := m.dialer.DialAndSend(m.gm)

	if err != nil {
		log.Printf("failed send email %v", err)
		return err
	}

	return nil
}

func (m *mailRepository) SendEmailCheckout(receiverEmail string) error {
	m.gm.SetHeader("From", "agusariputra70@gmail.com")
	m.gm.SetHeader("To", receiverEmail)
	m.gm.SetHeader("Subject", "Daily reminder to checkout")
	m.gm.SetBody("text/plain", "Remember to checkout")

	err := m.dialer.DialAndSend(m.gm)

	if err != nil {
		log.Printf("failed send email %v", err)
		return err
	}

	return nil
}
