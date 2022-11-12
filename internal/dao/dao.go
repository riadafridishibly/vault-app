package dao

import (
	"context"
	"fmt"

	"github.com/riadafridishibly/vault-app/ent"
)

type PasswordItem struct {
	Description string   `json:"description,omitempty"`
	SiteName    string   `json:"site_name,omitempty"`
	Username    string   `json:"username,omitempty"`
	Password    string   `json:"password,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

type Dao struct {
	client *ent.Client
}

func NewDao() (*Dao, error) {
	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		return nil, err
	}
	return &Dao{client: client}, nil
}

func (d Dao) Close() error {
	return d.client.Close()
}

func (d Dao) CreatePasswordItem(ctx context.Context, passwordItem PasswordItem) (*ent.PasswordItem, error) {
	p, err := d.client.PasswordItem.
		Create().
		SetDescription(passwordItem.Description).
		SetPassword(passwordItem.Password).
		SetSiteName(passwordItem.SiteName).
		SetUsername(passwordItem.Username).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating password item: %w", err)
	}
	return p, nil
}
