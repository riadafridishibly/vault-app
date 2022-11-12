package main

import (
	"github.com/iwpnd/pw"
	"github.com/riadafridishibly/vault-app/internal/breachchecker"
)

type PasswordService struct{}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (p PasswordService) CheckPasswordBreach(password string) (int, error) {
	return breachchecker.GetBreach(password)
}

func (a PasswordService) GeneratePassword(options []string) string {
	var opts []pw.Option
	for _, opt := range options {
		switch opt {
		case "uppercase":
			opts = append(opts, pw.WithUpperCase())
		case "lowercase":
			opts = append(opts, pw.WithLowerCase())
		case "numbers":
			opts = append(opts, pw.WithNumbers())
		case "special":
			opts = append(opts, pw.WithSpecial())
		}
	}

	return pw.NewPassword(20, opts...)
}
