package main

import (
	"context"
	"os"

	wailsruntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

type DialogService struct {
	ctx context.Context
}

func NewDialog(ctx context.Context) *DialogService {
	return &DialogService{ctx}
}

func (d *DialogService) updateContext(ctx context.Context) {
	d.ctx = ctx
}

func (d *DialogService) OpenDialog() ([]string, error) {
	return wailsruntime.OpenMultipleFilesDialog(d.ctx, wailsruntime.OpenDialogOptions{
		DefaultDirectory:           os.Getenv("HOME"),
		DefaultFilename:            "",
		Title:                      "Select Files",
		ShowHiddenFiles:            false,
		CanCreateDirectories:       false,
		ResolvesAliases:            false,
		TreatPackagesAsDirectories: false,
	})
}
