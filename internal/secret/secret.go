package secret

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"log"

	"golang.org/x/crypto/argon2"
)

var ErrUserNotLoggedIn = errors.New("user not logged in")

var masterPassword []byte

func IsLoggedIn() bool {
	return len(masterPassword) > 0
}

func MasterPassword() []byte {
	return masterPassword
}

func SetMasterPassword(pass []byte) {
	masterPassword = pass
}

func toArgon2(pass []byte) []byte {
	return argon2.IDKey(pass, []byte("vault-password-hash"), 1, 64*1024, 4, 32)
}

func HashPassword(password []byte) string {
	return base64.StdEncoding.EncodeToString(toArgon2(password))
}

func Verify(currentPasswordHash, masterPassword string) bool {
	value, err := base64.StdEncoding.DecodeString(currentPasswordHash)
	if err != nil {
		log.Println("hash is not a valid base64 string")
		return false
	}
	return subtle.ConstantTimeCompare(value, toArgon2([]byte(masterPassword))) == 1
}
