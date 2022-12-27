package dao

import (
	"context"
	"errors"
	"log"

	"filippo.io/age"
	"github.com/riadafridishibly/vault-app/ent"
	"github.com/riadafridishibly/vault-app/ent/key"
	"github.com/riadafridishibly/vault-app/internal/cryptox"
	"github.com/riadafridishibly/vault-app/internal/secret"
)

func (d *Dao) FindKeyByID(ctx context.Context, id int) (*ent.Key, error) {
	return d.client.Key.Get(ctx, id)
}

func (d *Dao) NewKey(ctx context.Context) (*ent.Key, error) {
	publicKey, privateKey, err := cryptox.GenerateNewAgeKey()
	if err != nil {
		return nil, err
	}

	masterPass := string(secret.MasterPassword())
	if masterPass == "" {
		return nil, errors.New("invalid master password: check if user is logged in")
	}

	encryptedPrivateKey, err := cryptox.EncryptWithPassword(privateKey, masterPass)
	if err != nil {
		return nil, err
	}

	return d.client.Key.
		Create().
		SetType("age"). // Age key
		SetPublicKey(publicKey).
		SetIsActive(false).
		SetReferences(0).
		SetPrivateKey(encryptedPrivateKey).
		Save(ctx)
}

func (d *Dao) RemoveKeyByID(ctx context.Context, id int) error {
	return d.client.Key.DeleteOneID(id).Exec(ctx)
}

func (d *Dao) FindAllKeys(ctx context.Context) ([]*ent.Key, error) {
	return d.client.Key.Query().All(ctx)
}

func (d *Dao) ActivateKey(ctx context.Context, id int) (*ent.Key, error) {
	return d.client.Key.UpdateOneID(id).SetIsActive(true).Save(ctx)
}

func (d *Dao) DeactivateKey(ctx context.Context, id int) (*ent.Key, error) {
	return d.client.Key.UpdateOneID(id).SetIsActive(false).Save(ctx)
}

func (d *Dao) AllActiveKeys(ctx context.Context) ([]*ent.Key, error) {
	return d.client.Key.Query().Where(key.IsActive(true)).All(ctx)
}

func (d *Dao) AllIdentities(ctx context.Context) (identities []age.Identity, err error) {
	// 1. check if we have access to master password
	masterPassword := string(secret.MasterPassword())
	if masterPassword == "" {
		return nil, secret.ErrUserNotLoggedIn
	}

	// 2. Get all keys (these private keys are encrypted with master password)
	allKeys, err := d.AllActiveKeys(ctx)
	if err != nil {
		return nil, err
	}
	for _, key := range allKeys {
		// decrypt the private key
		privKey, err := cryptox.DecryptWithPassword(key.PrivateKey, masterPassword)
		if err != nil {
			// maybe this one is encrypted with totally different master password?
			// we shouldn't allow that! When master password is changed all private keys should be reencrypted with the new key
			log.Println("couldn't decrypt with password", err)
			return nil, err
		}

		identitiy, err := cryptox.IdentityFromKey(privKey)
		if err != nil {
			log.Println("couldn't create identity from key", err)
			return nil, err
		}

		identities = append(identities, identitiy)
	}

	return identities, nil
}
