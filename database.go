package main

import (
	"context"

	"github.com/riadafridishibly/vault-app/internal/dao"
)

type DatabaseService struct {
	ctx context.Context
	dao *dao.Dao
}

func NewDatabaseService() (*DatabaseService, error) {
	dao, err := dao.NewDao()
	if err != nil {
		return nil, err
	}

	return &DatabaseService{
		dao: dao,
		ctx: context.TODO(),
	}, nil
}

func (db *DatabaseService) updateContext(ctx context.Context) {
	db.ctx = ctx
}

func (db *DatabaseService) CreatePasswordItem(passwordItem dao.PasswordItem) error {
	_, err := db.dao.CreatePasswordItem(db.ctx, passwordItem)
	return err
}

func (db *DatabaseService) close() error {
	return db.dao.Close()
}
