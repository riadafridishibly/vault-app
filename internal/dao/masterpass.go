package dao

import (
	"context"

	"github.com/riadafridishibly/vault-app/ent"
)

func (dao *Dao) MasterPasswordHash(ctx context.Context) (string, error) {
	value, err := dao.client.MasterPassword.Query().First(ctx)
	if err != nil {
		return "", err
	}
	return value.PasswordHash, nil
}

func (dao *Dao) CreateMasterPasswordHash(ctx context.Context, passwordHash string) (*ent.MasterPassword, error) {
	return dao.client.MasterPassword.Create().SetPasswordHash(passwordHash).Save(ctx)
}
