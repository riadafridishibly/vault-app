package cryptox

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"strings"

	"filippo.io/age"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/chacha20poly1305"
)

func GenerateNewAgeKey() (pubKey, privKey string, err error) {
	identity, err := age.GenerateX25519Identity()
	if err != nil {
		return "", "", err
	}

	pubKey = identity.Recipient().String()
	privKey = identity.String()
	return
}

func Derive32BytesKey(password string) []byte {
	return argon2.IDKey([]byte(password), []byte("vault-app-salt"), 1, 64*1024, 4, 32)
}

func EncryptWithPassword(plaintext, password string) (string, error) {
	aead, err := chacha20poly1305.New(Derive32BytesKey(password))
	if err != nil {
		return "", err
	}

	plainTextBytes := []byte(plaintext)
	nonce := make([]byte,
		aead.NonceSize(),
		aead.NonceSize()+len(plainTextBytes)+aead.Overhead())
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	cipertextBytes := aead.Seal(nonce, nonce, plainTextBytes, nil)
	return base64.StdEncoding.EncodeToString(cipertextBytes), nil
}

func DecryptWithPassword(b64string, password string) (string, error) {
	aead, err := chacha20poly1305.New(Derive32BytesKey(password))
	if err != nil {
		return "", err
	}

	ciphertextString, err := base64.StdEncoding.DecodeString(b64string)
	if err != nil {
		return "", err
	}

	ciphertextBytes := []byte(ciphertextString)
	if len(ciphertextBytes) < aead.NonceSize() {
		return "", errors.New("invalid ciphertext: too short")
	}

	nonce, ciphertext := ciphertextBytes[:aead.NonceSize()],
		ciphertextBytes[aead.NonceSize():]

	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	return string(plaintext), err
}

func IdentityFromKey(privateKey string) (age.Identity, error) {
	ids, err := age.ParseIdentities(strings.NewReader(privateKey))
	if err != nil {
		return nil, err
	}

	return ids[0], nil
}

func RecipientFromKey(publicKey string) (age.Recipient, error) {
	recipients, err := age.ParseRecipients(strings.NewReader(publicKey))
	if err != nil {
		return nil, err
	}

	return recipients[0], nil
}

// EncryptStringB64 encrypts string and return base64 output
func EncryptStringB64(plaintext string, recipients ...age.Recipient) (encryptedText string, err error) {
	buf := new(bytes.Buffer)
	wc, err := age.Encrypt(buf, recipients...)
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(wc, strings.NewReader(plaintext)); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

// DecryptStringB64 accepts an encrypted base64 string, return the plaintext
func DecryptStringB64(base64string string, identities ...age.Identity) (plaintext string, err error) {
	encReader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64string))
	r, err := age.Decrypt(encReader, identities...)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, r); err != nil {
		return "", err
	}

	return buf.String(), nil
}
