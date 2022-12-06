package dao

import (
	"context"

	"github.com/riadafridishibly/vault-app/ent"
	"github.com/riadafridishibly/vault-app/ent/key"
	"github.com/riadafridishibly/vault-app/internal/cryptox"
)

func (d *Dao) FindKeyByID(ctx context.Context, id int) (*ent.Key, error) {
	return d.client.Key.Get(ctx, id)
}

func (d *Dao) NewKey(ctx context.Context) (*ent.Key, error) {
	publicKey, privateKey, err := cryptox.GenerateNewAgeKey()
	if err != nil {
		return nil, err
	}

	return d.client.Key.
		Create().
		SetType("age"). // Age key
		SetPublicKey(publicKey).
		SetIsActive(false).
		SetReferences(0).
		SetPrivateKey(privateKey). // TODO: encrypt the private key with the master key
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
