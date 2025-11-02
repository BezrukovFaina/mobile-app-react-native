package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Sender interface {
	Send(data string) error
}

type EmailSender struct{}

func (s *EmailSender) Send(data string) error {
	// email sending logic here
	return nil
}

type SmsSender struct{}

func (s *SmsSender) Send(data string) error {
	// sms sending logic here
	return nil
}

func GenerateRandomString(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func IsEmailValid(email string) bool {
	return strings.Contains(email, "@")
}

func SHA256(text string) string {
	h := sha256.New()
	io.WriteString(h, text)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func CheckPasswordHash(password, hash string) bool {
	return SHA256(password) == hash
}

func NewSender(senderType string) Sender {
	switch strings.ToLower(senderType) {
	case "email":
		return &EmailSender{}
	case "sms":
		return &SmsSender{}
	default:
		log.Fatal("Invalid sender type")
		return nil
	}
}