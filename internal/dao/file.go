package dao

import (
	"context"

	"github.com/riadafridishibly/vault-app/ent"
)

func (d *Dao) FindFileByID(ctx context.Context, id int) (*ent.File, error) {
	return d.client.File.Get(ctx, id)
}

func (d *Dao) NewFile(ctx context.Context, file ent.File) (*ent.File, error) {
	return d.client.File.
		Create().
		SetSrcPath(file.SrcPath).
		SetCurrentPath(file.CurrentPath).
		SetFiletype(file.Filetype).
		Save(ctx)
}

func (d *Dao) RemoveFileByID(ctx context.Context, id int) error {
	return d.client.File.DeleteOneID(id).Exec(ctx)
}

func (d *Dao) FindAllFiles(ctx context.Context) ([]*ent.File, error) {
	return d.client.File.Query().All(ctx)
}
