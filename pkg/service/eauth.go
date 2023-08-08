package service

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/smtp"
	"strconv"
	"time"
)

type EAuthService interface {
	Add(data EAuth) error
	Find(email string) (EAuth, error)
	Delete(email string) error
	SendOTP(email string) error
}

type eauth struct {
	EAuthService
	EAuthBasket
}

type EAuthBasket struct {
	Basket []EAuth
}

type EAuth struct {
	OTP   int
	Email string
	time.Time
}

type VerifyedEmails struct {
	Emails []string
	VerifyedEmailsService
}

type VerifyedEmailsService interface {
	Save(email string) error
	Delete(email string) error
	Find(email string) bool
}

func NewEAuthService() eauth {
	return eauth{EAuthBasket: EAuthBasket{Basket: []EAuth{}}}
}

func NewVerifiedEmailsService() *VerifyedEmails {
	return &VerifyedEmails{}
}

func (b *EAuthBasket) Add(data EAuth) error {
	_, err := b.Find(data.Email)
	if err == nil {
		return errors.New("ALREADY EXISTS, WAIT FOR NEXT PERIOD")
	}
	b.Basket = append(b.Basket, data)
	return nil
}

func (b *EAuthBasket) Find(email string) (EAuth, error) {
	for _, v := range b.Basket {
		if v.Email == email {
			return v, nil
		}
	}
	return EAuth{}, errors.New("NOT FOUND")
}

func (b *EAuthBasket) Delete(email string) error {
	for i, v := range b.Basket {
		if v.Email == email {
			b.Basket = append(b.Basket[:i], b.Basket[i+1:]...)
			return nil
		}
	}
	return errors.New("NOT FOUND")
}

func (s *eauth) SendOTP(email string) error {
	OTP := generateOTP()
	otp, err := strconv.Atoi(OTP)
	if err != nil {
		log.Println("Error converting otp to int.", err)
	}

	err = s.EAuthBasket.Add(EAuth{OTP: otp, Email: email, Time: time.Now()})
	if err != nil {
		return err
	}
	go s.EAuthBasket.eAuthBasketCleaner(email)

	if err := SendEmail(email, OTP); err != nil {
		return err
	}
	return nil
}

func SendEmail(email string, OTP string) error {
	from := "mikemcgrat11@gmail.com"
	password := "enkvpkvpyaaogaue"

	to := []string{email}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte(OTP)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}

	log.Println("OTP sent successfully to", email)

	return nil
}

func generateOTP() string {
	var otp string
	max := big.NewInt(999999)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	otp = fmt.Sprintf("%06d", n)
	return otp
}

func (b *EAuthBasket) eAuthBasketCleaner(email string) {
	time.Sleep(2 * time.Minute)
	if err := b.Delete(email); err != nil {
		log.Println("Error deleting saved OTP", err)
	}
}

func (e *VerifyedEmails) Save(email string) error {
	if e.Find(email) {
		return errors.New("ALREADY VERIFIED")
	}
	e.Emails = append(e.Emails, email)
	return nil
}

func (e *VerifyedEmails) Delete(email string) error {
	for i, v := range e.Emails {
		if v == email {
			e.Emails = append(e.Emails[:i], e.Emails[i+1:]...)
			return nil
		}
	}
	return errors.New("NOT FOUND")
}

func (e *VerifyedEmails) Find(email string) (finded bool) {
	finded = false
	for _, v := range e.Emails {
		if v == email {
			finded = true
			return
		}
	}
	return
}
