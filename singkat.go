package main

import (
	"crypto/rand"
	"errors"
	"regexp"

	"github.com/google/uuid"
)

type Shorten struct {
	ID         string `json:"id"`
	URL        string `json:"url"`
	URLShorted string `json:"url_shorted"`
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var ErrUrl = errors.New("URL invalid")
var ErrLeng = errors.New("length must be greater than 0")
var ErrCharsetEmpty = errors.New("charset cannot be empty")

func ValidateLength(length uint16) error {
	if length <= 0 {
		return ErrLeng
	}
	return nil
}

func GenerateRandomBytes(length uint16) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func ConvertBytesToChars(bytes []byte, charset string) (string, error) {
	if len(charset) == 0 {
		return "", ErrCharsetEmpty
	}

	randomString := make([]byte, len(bytes))
	for i := range bytes {
		randomString[i] = charset[bytes[i]%byte(len(charset))]
	}
	return string(randomString), nil
}

func GenerateRandomString(length uint16) (string, error) {
	err := ValidateLength(length)
	if err != nil {
		return "", err
	}

	randomBytes, err := GenerateRandomBytes(length)
	if err != nil {
		return "", err
	}

	randomString, err := ConvertBytesToChars(randomBytes, charset)
	if err != nil {
		return "", err
	}

	return randomString, nil
}

func IsValidURL(url string) error {
	regex := regexp.MustCompile(`^(https?://)?(www\.)?([a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-z]{2,6})\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)$`)
	if !regex.MatchString(url) {
		return ErrUrl
	}
	return nil
}

func GenerateID() (string, error) {
	return uuid.NewString(), nil
}

func NewShorten(url string) (*Shorten, error) {
	id, err := GenerateID()
	if err != nil {
		return nil, err
	}

	if err := IsValidURL(url); err != nil {
		return nil, err
	}

	newUrl, err := GenerateRandomString(7)
	if err != nil {
		return nil, err
	}

	return &Shorten{
		ID:         id,
		URL:        url,
		URLShorted: newUrl,
	}, nil
}
