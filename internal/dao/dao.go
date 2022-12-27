package dao

import (
	"context"
	"errors"
	"fmt"
	"log"

	"filippo.io/age"
	"github.com/riadafridishibly/vault-app/ent"
	"github.com/riadafridishibly/vault-app/internal/cryptox"
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

func (d *Dao) Close() error {
	return d.client.Close()
}

func (d *Dao) decryptPasswordWithAnyActiveKeys(ctx context.Context, encryptedPassword string) (passowrd string, err error) {
	identities, err := d.AllIdentities(ctx)
	if err != nil {
		log.Println("error building identity: ", err, encryptedPassword)
		return "", err
	}
	return cryptox.DecryptStringB64(passowrd, identities...)
}

func (d *Dao) encryptPasswordWithAllActiveKeys(ctx context.Context, password string) (encryptedPassword string, err error) {
	keys, err := d.AllActiveKeys(ctx)
	if err != nil {
		return "", err
	}
	if len(keys) == 0 {
		return "", errors.New("no active age key found")
	}

	recipients := make([]age.Recipient, 0)
	for _, key := range keys {
		r, err := cryptox.RecipientFromKey(key.PublicKey)
		if err != nil {
			// TODO: Update logging
			log.Println("error: creating recipient from key", err)
			return "", err
		}
		recipients = append(recipients, r)
	}

	encryptedText, err := cryptox.EncryptStringB64(password, recipients...)
	if err != nil {
		return "", err
	}
	return encryptedText, nil
}

func (d *Dao) CreatePasswordItem(ctx context.Context, p ent.PasswordItem) (*ent.PasswordItem, error) {
	builder := d.client.PasswordItem.Create()
	if p.Description != nil {
		builder = builder.SetDescription(*p.Description)
	}
	if p.Password != nil {
		// TODO: maybe use ent.Hooks for this?
		pass, err := d.encryptPasswordWithAllActiveKeys(ctx, *p.Password)
		if err != nil {
			return nil, err
		}
		builder = builder.SetPassword(pass)
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

func (d *Dao) UpdatePasswordItem(ctx context.Context, p ent.PasswordItem) (*ent.PasswordItem, error) {
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
		pass, err := d.encryptPasswordWithAllActiveKeys(ctx, *p.Password)
		if err != nil {
			return nil, err
		}
		builder = builder.SetPassword(pass)
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

func (d *Dao) DeletePasswordItem(ctx context.Context, id int) (*ent.PasswordItem, error) {
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

func (d *Dao) ReadAllPasswordItems(ctx context.Context) ([]*ent.PasswordItem, error) {
	identities, err := d.AllIdentities(ctx)
	if err != nil {
		return nil, err
	}
	values, err := d.client.PasswordItem.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	for _, item := range values {
		if item.Password == nil {
			continue
		}

		decrypted, err := cryptox.DecryptStringB64(*item.Password, identities...)
		if err != nil {
			return nil, err
		}
		*item.Password = decrypted
	}

	return values, nil
}

func (d *Dao) ReadSinglePasswordItems(ctx context.Context, id int) (*ent.PasswordItem, error) {
	item, err := d.client.PasswordItem.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// Password is not set! nothing to do
	if item.Password == nil {
		return item, err
	}

	decrypted, err := d.decryptPasswordWithAnyActiveKeys(ctx, *item.Password)
	if err != nil {
		return nil, err
	}
	*item.Password = decrypted
	return item, nil
}
