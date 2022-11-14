package main

import (
	"context"

	"github.com/riadafridishibly/vault-app/ent"
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

func (db *DatabaseService) CreatePasswordItem(p ent.PasswordItem) (*ent.PasswordItem, error) {
	return db.dao.CreatePasswordItem(db.ctx, p)
}

func (db *DatabaseService) UpdatePasswordItem(p ent.PasswordItem) (*ent.PasswordItem, error) {
	return db.dao.UpdatePasswordItem(db.ctx, p)
}

func (db *DatabaseService) DeletePasswordItem(id int) (*ent.PasswordItem, error) {
	return db.dao.DeletePasswordItem(db.ctx, id)
}

func (db *DatabaseService) ReadAllPasswordItems() ([]*ent.PasswordItem, error) {
	return db.dao.ReadAllPasswordItems(db.ctx)
}

func (db *DatabaseService) close() error {
	return db.dao.Close()
}