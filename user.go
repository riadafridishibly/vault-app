package main

import (
	"errors"

	"github.com/riadafridishibly/vault-app/internal/secret"
)

type UserService struct {
	db *DatabaseService
}

func (u *UserService) setDB(db *DatabaseService) {
	u.db = db
}

func (u *UserService) IsUserExist() bool {
	pass, err := u.db.ReadMasterPasswordHash()
	if err != nil || pass == "" {
		return false
	}
	return true
}

func (u *UserService) IsUserLoggedIn() bool {
	return u.IsUserExist() && len(secret.MasterPassword()) > 0
}

func (u *UserService) Register(password string) (bool, error) {
	if u.IsUserExist() {
		return false, errors.New("user already exists: please log in")
	}
	passwordHash := secret.HashPassword([]byte(password))
	return true, u.db.CreateMasterPasswordHash(passwordHash)
}

func (u *UserService) Login(password string) error {
	passHash, err := u.db.ReadMasterPasswordHash()
	if err != nil {
		return err
	}
	if !secret.Verify(passHash, password) {
		return errors.New("invalid master password")
	}
	secret.SetMasterPassword([]byte(password))
	return nil
}
