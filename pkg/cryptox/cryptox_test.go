package cryptox

import (
	"testing"

	"filippo.io/age"
)

func TestEncryptDecryptString(t *testing.T) {
	const s = "Hello, World!"
	id, err := age.GenerateX25519Identity()
	if err != nil {
		t.Error("Failed to create identity", err)
	}

	encryptedStr, err := EncryptStringB64(s, id.Recipient())
	if err != nil {
		t.Error("Failed to encrypt value", err)
	}

	plaintext, err := DecryptStringB64(encryptedStr, id)
	if err != nil {
		t.Error("Failed to decrypt value", err)
	}

	if plaintext != s {
		t.Error("Failed to retrieve original value")
	}
}

func TestEncryptDecryptWithPassword(t *testing.T) {
	const s = "Hello, World!"
	encText, err := EncryptWithPassword(s, "12345")
	if err != nil {
		t.Error("Failed to encrypt")
	}
	decText, err := DecryptWithPassword(encText, "12345")
	if err != nil {
		t.Error("Failed to decrypt")
	}

	if s != decText {
		t.Error("Failed to retrieve original value")
	}
}

func TestConvertKeys(t *testing.T) {
	pubKey, privKey, err := GenerateNewAgeKey()
	if err != nil {
		t.Error("Failed to generate keys", err)
	}

	id, err := IdentityFromKey(privKey)
	if err != nil {
		t.Error("Failed to create identity from private key", err)
	}
	recipient, err := RecipientFromKey(pubKey)
	if err != nil {
		t.Error("Failed to create recipient from public key", err)
	}

	const s = "Hello, World!"
	encryptedStr, err := EncryptStringB64(s, recipient)
	if err != nil {
		t.Error("Failed to encrypt value", err)
	}

	plaintext, err := DecryptStringB64(encryptedStr, id)
	if err != nil {
		t.Error("Failed to decrypt value", err)
	}

	if plaintext != s {
		t.Error("Failed to retrieve original value")
	}
}
