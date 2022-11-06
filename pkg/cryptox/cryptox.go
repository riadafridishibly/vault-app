package cryptox

import (
	"bytes"
	"encoding/base64"
	"io"
	"strings"

	"filippo.io/age"
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
