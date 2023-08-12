package service

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/smtp"
	"time"

	online_diler "github.com/cora23tt/online-diler"
)

type Session struct {
	online_diler.User
	OTP string
}

type EAuthService interface {
	Get(email string) (Session, error)
	Delete(email string) error
	SendOTP(user online_diler.User) error
	CheckOTP(email string, OTP string) error
}

type eAuth struct {
	Sessions []Session
	EAuthService
}

func NewEAuthService() *eAuth {
	return &eAuth{}
}

func (s *eAuth) Get(email string) (Session, error) {
	for _, v := range s.Sessions {

		if v.Email == email {
			return v, nil
		}
	}

	return Session{}, errors.New("NOT FOUND")
}

func (s *eAuth) Delete(email string) error {
	_, err := s.Get(email)

	if err != nil { // err=notfound
		return err
	}

	for i, v := range s.Sessions {
		if v.Email == email {
			s.Sessions = append(s.Sessions[:i], s.Sessions[i+1:]...)
			return nil
		}
	}

	return errors.New("CAN`T DELETE")
}

func (s *eAuth) SendOTP(user online_diler.User) error {
	_, err := s.Get(user.Email)

	if err == nil {
		return errors.New("ALREADY EXIST")
	}

	otp := generateOTP()

	s.Sessions = append(s.Sessions, Session{User: user, OTP: otp})

	go s.removeAfter(user.Email)

	if err := SendEmail(user.Email, otp); err != nil {
		return err
	}
	return nil
}

func (s *eAuth) CheckOTP(email string, OTP string) error {
	session, err := s.Get(email)
	if err != nil {
		return err
	}

	if session.OTP == OTP {
		return nil
	}

	return errors.New("DON`T MATCH")
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

// Removes given email from sessions after 2 min
func (s *eAuth) removeAfter(email string) {
	time.Sleep(150 * time.Second)
	if err := s.Delete(email); err != nil {
		log.Println("Error deleting saved OTP", err)
	}
}
