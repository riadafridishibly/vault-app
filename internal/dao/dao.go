package dao

import (
	"context"
	"fmt"

	"github.com/riadafridishibly/vault-app/ent"
)

type Dao struct {
	client *ent.Client
}

func NewDao() (*Dao, error) {
	client, err := ent.Open("sqlite3", "file:vault.db?cache=shared&_fk=1")
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}
	return &Dao{client: client}, nil
}

func (d Dao) Close() error {
	return d.client.Close()
}

func (d Dao) CreatePasswordItem(ctx context.Context, p ent.PasswordItem) (*ent.PasswordItem, error) {
	builder := d.client.PasswordItem.Create()
	if p.Description != nil {
		builder = builder.SetDescription(*p.Description)
	}
	if p.Password != nil {
		builder = builder.SetPassword(*p.Password)
	}
	if p.Avatar != nil {
		builder = builder.SetAvatar(*p.Avatar)
	}
	if p.SiteURL != nil {
		builder = builder.SetSiteURL(*p.SiteURL)
	}
	if p.SiteName != nil {
		builder = builder.SetSiteName(*p.SiteName)
	}
	if p.Username != nil {
		builder = builder.SetUsername(*p.Username)
	}
	if p.Tags != nil {
		builder = builder.SetTags(p.Tags)
	}
	pp, err := builder.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating password item: %w", err)
	}
	return pp, nil
}

func (d Dao) UpdatePasswordItem(ctx context.Context, p ent.PasswordItem) (*ent.PasswordItem, error) {
	builder := d.client.PasswordItem.UpdateOneID(p.ID)
	if p.Description != nil {
		builder = builder.SetDescription(*p.Description)
	}
	if p.Avatar != nil {
		builder = builder.SetAvatar(*p.Avatar)
	}
	if p.SiteURL != nil {
		builder = builder.SetSiteURL(*p.SiteURL)
	}
	if p.Password != nil {
		builder = builder.SetPassword(*p.Password)
	}
	if p.SiteName != nil {
		builder = builder.SetSiteName(*p.SiteName)
	}
	if p.Username != nil {
		builder = builder.SetUsername(*p.Username)
	}
	if p.Tags != nil {
		builder = builder.SetTags(p.Tags)
	}
	pp, err := builder.Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating password item: %w", err)
	}
	return pp, nil
}

func (d Dao) DeletePasswordItem(ctx context.Context, id int) (*ent.PasswordItem, error) {
	p, err := d.client.PasswordItem.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	err = d.client.PasswordItem.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating password item: %w", err)
	}
	return p, nil
}

func (d Dao) ReadAllPasswordItems(ctx context.Context) ([]*ent.PasswordItem, error) {
	return d.client.PasswordItem.Query().All(ctx)
}
